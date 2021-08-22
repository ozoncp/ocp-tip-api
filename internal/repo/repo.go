package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-tip-api/internal/models"
	"strings"
)

// Repo - интерфейс хранилища для сущности Tip
type Repo interface {
	AddTip(ctx context.Context, tip models.Tip) (uint64, error)
	AddTips(ctx context.Context, tips []models.Tip) ([]uint64, error)
	UpdateTip(ctx context.Context, tip models.Tip) error
	ListTips(ctx context.Context, limit, offset uint64) ([]models.Tip, error)
	DescribeTip(ctx context.Context, tipId uint64) (*models.Tip, error)
	RemoveTip(ctx context.Context, tipId uint64) (bool, error)
}

type repo struct {
	db sqlx.DB
}

func (r *repo) AddTip(ctx context.Context, tip models.Tip) (uint64, error) {
	query := "INSERT INTO tips(user_id, problem_id, text) VALUES ($1, $2, $3) RETURNING id"
	row := r.db.QueryRowContext(ctx, query, tip.UserId, tip.ProblemId, tip.Text)

	var id uint64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repo) AddTips(ctx context.Context, tips []models.Tip) ([]uint64, error) {
	query := "INSERT INTO tips(user_id, problem_id, text) VALUES "
	placeholders := make([]string, 0, len(tips)*3)
	values := make([]interface{}, 0, len(tips)*3)
	for i, tip := range tips {
		placeholders = append(placeholders, fmt.Sprintf("($%d,$%d,$%d)", i*3+1, i*3+2, i*3+3))
		values = append(values, tip.UserId, tip.ProblemId, tip.Text)
	}
	query += strings.Join(placeholders, ",")
	query += " RETURNING id"
	rows, err := r.db.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	ids := make([]uint64, 0, len(tips))
	for rows.Next() {
		var tipId uint64
		if err := rows.Scan(&tipId); err != nil {
			return nil, err
		}
		ids = append(ids, tipId)
	}

	return ids, nil
}

func (r *repo) UpdateTip(ctx context.Context, tip models.Tip) error {
	query := "UPDATE tips SET user_id = $1, problem_id = $2, text = $3 WHERE id = $4"
	res, err := r.db.ExecContext(ctx, query, tip.UserId, tip.ProblemId, tip.Text, tip.Id)
	if err != nil {
		return err
	}
	rowsCount, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsCount == 0 {
		return errors.New("tip not found")
	}
	return nil
}

func (r *repo) ListTips(ctx context.Context, limit, offset uint64) (tips []models.Tip, err error) {
	query := "SELECT id, user_id, problem_id, text FROM tips LIMIT $1 OFFSET $2"
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			tips = nil
		}
	}()

	for rows.Next() {
		var tip models.Tip
		if err := rows.Scan(&tip.Id, &tip.UserId, &tip.ProblemId, &tip.Text); err != nil {
			return nil, err
		}
		tips = append(tips, tip)
	}
	return tips, nil
}

func (r *repo) DescribeTip(ctx context.Context, tipId uint64) (*models.Tip, error) {
	query := "SELECT id, user_id, problem_id, text FROM tips WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, tipId)
	var tip models.Tip
	err := row.Scan(&tip.Id, &tip.UserId, &tip.ProblemId, &tip.Text)
	if err != nil {
		return nil, err
	}
	return &tip, nil
}

func (r *repo) RemoveTip(ctx context.Context, tipId uint64) (bool, error) {
	query := "DELETE FROM tips WHERE id = $1"
	res, err := r.db.ExecContext(ctx, query, tipId)
	if err != nil {
		return false, err
	}
	rowsCount, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	return rowsCount > 0, nil
}

func NewRepo(db *sqlx.DB) Repo {
	return &repo{db: *db}
}

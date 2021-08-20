package repo

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-tip-api/internal/models"
)

const tableName = "tips"

// Repo - интерфейс хранилища для сущности Tip
type Repo interface {
	AddTip(ctx context.Context, tip models.Tip) (uint64, error)
	AddTips(ctx context.Context, tips []models.Tip) error
	ListTips(ctx context.Context, limit, offset uint64) ([]models.Tip, error)
	DescribeTip(ctx context.Context, tipId uint64) (*models.Tip, error)
	RemoveTip(ctx context.Context, tipId uint64) (bool, error)
}

type repo struct {
	db sqlx.DB
}

func (r *repo) AddTip(ctx context.Context, tip models.Tip) (uint64, error) {
	query := sq.Insert(tableName).
		Columns("user_id", "problem_id", "text").
		Values(tip.UserId, tip.ProblemId, tip.Text).
		Suffix("RETURNING \"id\"").
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	if err := query.QueryRow().Scan(&tip.Id); err != nil {
		return 0, err
	}

	return tip.Id, nil
}

func (r *repo) AddTips(ctx context.Context, tips []models.Tip) error {
	query := sq.Insert(tableName).
		Columns("user_id", "problem_id", "text").
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	for _, tip := range tips {
		query = query.Values(tip.UserId, tip.ProblemId, tip.Text)
	}

	_, err := query.ExecContext(ctx)
	return err
}

func (r *repo) ListTips(ctx context.Context, limit, offset uint64) ([]models.Tip, error) {
	query := sq.Select("id", "user_id", "problem_id", "text").
		From(tableName).
		RunWith(r.db).
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(sq.Dollar)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	var tips []models.Tip

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
	query := sq.Select("id", "user_id", "problem_id", "text").
		From(tableName).
		Where(sq.Eq{"id": tipId}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	var tip models.Tip
	if err := query.QueryRowContext(ctx).Scan(&tip.Id, &tip.UserId, &tip.ProblemId, &tip.Text); err != nil {
		return nil, err
	}
	return &tip, nil
}

func (r *repo) RemoveTip(ctx context.Context, tipId uint64) (bool, error) {
	query := sq.Delete(tableName).
		Where(sq.Eq{"id": tipId}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)
	res, err := query.ExecContext(ctx)
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

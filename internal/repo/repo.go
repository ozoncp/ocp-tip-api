package repo

import "github.com/ozoncp/ocp-tip-api/internal/models"

// Repo - интерфейс хранилища для сущности Tip
type Repo interface {
	AddTips(tips []models.Tip) error
	ListTips(limit, offset uint64) ([]models.Tip, error)
	DescribeTip(tipId uint64) (*models.Tip, error)
}

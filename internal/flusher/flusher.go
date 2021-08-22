package flusher

import (
	"context"
	"github.com/ozoncp/ocp-tip-api/internal/models"
	"github.com/ozoncp/ocp-tip-api/internal/repo"
	"github.com/ozoncp/ocp-tip-api/internal/utils"
)

// Flusher - интерфейс для сброса советов в хранилище
type Flusher interface {
	Flush(tips []models.Tip) []models.Tip
}

type flusher struct {
	batchSize int
	repo      repo.Repo
}

// Flush сбрасывает советы в хранилище по батчам
func (f *flusher) Flush(tips []models.Tip) []models.Tip {
	notFlushedTips := make([]models.Tip, 0, len(tips))
	for _, batch := range utils.SplitTipsByBatches(tips, f.batchSize) {
		if _, err := f.repo.AddTips(context.Background(), batch); err != nil {
			notFlushedTips = append(notFlushedTips, batch...)
		}
	}
	if len(notFlushedTips) == 0 {
		return nil
	}
	return notFlushedTips
}

// NewFlusher возвращает Flusher с поддержкой батчевого сохранения
func NewFlusher(batchSize int, repo repo.Repo) Flusher {
	return &flusher{
		batchSize: batchSize,
		repo:      repo,
	}
}

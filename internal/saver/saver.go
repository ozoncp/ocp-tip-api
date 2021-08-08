package saver

import (
	"errors"
	"github.com/ozoncp/ocp-tip-api/internal/flusher"
	"github.com/ozoncp/ocp-tip-api/internal/models"
	"time"
)

// Saver - интерфейс для сохранения советов в хранилище с периодичностью
type Saver interface {
	Save(tip models.Tip) error
	Init() error
	Close() error
}

type saver struct {
	tips        []models.Tip
	tipsChan    chan models.Tip
	capacity    uint
	interval    time.Duration
	flusher     flusher.Flusher
	sigChan     chan struct{}
	isInitiated bool
}

// Save добавляет совет в очередь для сохранения
func (s *saver) Save(tip models.Tip) error {
	if !s.isInitiated {
		return errors.New("saver is not initiated")
	}
	s.tipsChan <- tip
	return nil
}

// Init запускает периодическое сохранение советов
func (s *saver) Init() error {
	if s.isInitiated {
		return errors.New("saver has been initiated already")
	}
	s.tipsChan = make(chan models.Tip)
	s.sigChan = make(chan struct{})
	s.isInitiated = true

	go func() {
		ticker := time.NewTicker(s.interval * time.Second)
		defer ticker.Stop()

		flush := func() {
			if len(s.tips) == 0 {
				return
			}
			notSavedTips := s.flusher.Flush(s.tips)
			if notSavedTips == nil {
				s.tips = make([]models.Tip, 0, s.capacity)
			} else {
				s.tips = notSavedTips
			}
		}

		for {
			select {
			case <-ticker.C:
				flush()
			case tip := <-s.tipsChan:
				s.tips = append(s.tips, tip)
			case <-s.sigChan:
				flush()
				return
			}
		}
	}()

	return nil
}

// Close останавливает периодическое сохранение советов
func (s *saver) Close() error {
	if !s.isInitiated {
		return errors.New("saver is not initiated")
	}
	defer func() {
		close(s.tipsChan)
		close(s.sigChan)
		s.isInitiated = false
	}()
	s.sigChan <- struct{}{}
	return nil
}

// NewSaver возвращает Saver с поддержкой переодического сохранения
func NewSaver(capacity uint, flusher flusher.Flusher, interval time.Duration) Saver {
	return &saver{
		tips:     make([]models.Tip, 0, capacity),
		capacity: capacity,
		interval: interval,
		flusher:  flusher,
	}
}

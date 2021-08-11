package saver

import (
	"github.com/ozoncp/ocp-tip-api/internal/flusher"
	"github.com/ozoncp/ocp-tip-api/internal/models"
	"time"
)

// Saver - интерфейс для сохранения советов в хранилище с периодичностью
type Saver interface {
	Save(tip models.Tip)
	Init()
	Close()
}

type saver struct {
	tipsChan chan models.Tip
	capacity uint
	interval time.Duration
	flusher  flusher.Flusher
	sigChan  chan struct{}
}

// Save добавляет совет в очередь для сохранения
func (s *saver) Save(tip models.Tip) {
	s.tipsChan <- tip
}

// Init запускает периодическое сохранение советов
func (s *saver) Init() {
	s.tipsChan = make(chan models.Tip)
	s.sigChan = make(chan struct{})

	go func() {
		ticker := time.NewTicker(s.interval * time.Second)
		defer ticker.Stop()

		tips := make([]models.Tip, 0, s.capacity)

		flush := func() {
			if len(tips) == 0 {
				return
			}
			notSavedTips := s.flusher.Flush(tips)
			if notSavedTips == nil {
				tips = tips[:0]
			} else {
				if cap(notSavedTips) < int(s.capacity) {
					tips = append(make([]models.Tip, 0, s.capacity), notSavedTips...)
				} else {
					tips = notSavedTips
				}
			}
		}

		for {
			select {
			case <-ticker.C:
				flush()
			case tip := <-s.tipsChan:
				tips = append(tips, tip)
				if len(tips) == cap(tips) {
					flush()
				}
			case <-s.sigChan:
				flush()
				return
			}
		}
	}()
}

// Close останавливает периодическое сохранение советов
func (s *saver) Close() {
	close(s.tipsChan)
	close(s.sigChan)
}

// NewSaver возвращает Saver с поддержкой переодического сохранения
func NewSaver(capacity uint, flusher flusher.Flusher, interval time.Duration) Saver {
	return &saver{
		capacity: capacity,
		interval: interval,
		flusher:  flusher,
	}
}

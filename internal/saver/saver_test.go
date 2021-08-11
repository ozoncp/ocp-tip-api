package saver_test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/ozoncp/ocp-tip-api/internal/mocks"
	"github.com/ozoncp/ocp-tip-api/internal/models"
	"github.com/ozoncp/ocp-tip-api/internal/saver"
	"math/rand"
	"time"
)

var _ = Describe("Saver", func() {
	const tipsCount int = 50
	const saverCapacity uint = 5
	const saverInterval time.Duration = 3

	var (
		ctrl        *gomock.Controller
		mockFlusher *mocks.MockFlusher
		tips        []models.Tip
		s           saver.Saver
	)

	tips = make([]models.Tip, tipsCount)
	rand.Seed(time.Now().Unix())
	for i := 0; i < tipsCount; i++ {
		tips[i] = models.Tip{
			Id:        uint64(i + 1),
			UserId:    rand.Uint64(),
			ProblemId: rand.Uint64(),
			Text:      fmt.Sprintf("Tip number %d", i+1),
		}
	}

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockFlusher = mocks.NewMockFlusher(ctrl)
		s = saver.NewSaver(saverCapacity, mockFlusher, saverInterval)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("When flushing is successful", func() {
		JustBeforeEach(func() {
			mockFlusher.EXPECT().Flush(gomock.Any()).Return(nil).AnyTimes()
		})
		It("saving is succeeded", func() {
			s.Init()
			for _, tip := range tips {
				s.Save(tip)
			}
			time.Sleep(saverInterval * time.Second)
			s.Close()
		})
	})
})

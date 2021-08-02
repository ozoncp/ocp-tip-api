package flusher_test

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-tip-api/internal/flusher"
	"github.com/ozoncp/ocp-tip-api/internal/mocks"
	"github.com/ozoncp/ocp-tip-api/internal/models"
	"math/rand"
	"time"
)

var _ = Describe("Flusher", func() {
	const tipsCount int = 123
	const batchSize int = 5
	const batchesCount int = 25

	var (
		ctrl     *gomock.Controller
		mockRepo *mocks.MockRepo
		tips     []models.Tip
		result   []models.Tip
		f        flusher.Flusher
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
		mockRepo = mocks.NewMockRepo(ctrl)
		f = flusher.NewFlusher(batchSize, mockRepo)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("When all tips was flushed successfully", func() {
		It("returned value is nil", func() {
			mockRepo.EXPECT().AddTips(gomock.Any()).Return(nil).AnyTimes()
			result = f.Flush(tips)
			Expect(result).Should(BeNil())
		})
	})

	Context("When error was occurred while flushing tips", func() {
		It("returned slice of tips that have not been flushed", func() {
			gomock.InOrder(
				mockRepo.EXPECT().AddTips(gomock.Any()).Return(nil).Times(batchesCount/3),
				mockRepo.EXPECT().AddTips(gomock.Any()).Return(errors.New("unknown error")).Times(batchesCount/3),
				mockRepo.EXPECT().AddTips(gomock.Any()).Return(nil).AnyTimes(),
			)
			result = f.Flush(tips)
			start := batchesCount / 3 * batchSize
			end := 2 * start
			Expect(result).Should(Equal(tips[start:end]))
		})
	})
})

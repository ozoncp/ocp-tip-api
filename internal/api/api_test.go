package api_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-tip-api/internal/api"
	"github.com/ozoncp/ocp-tip-api/internal/models"
	"github.com/ozoncp/ocp-tip-api/internal/repo"
	desc "github.com/ozoncp/ocp-tip-api/pkg/ocp-tip-api"
	"math/rand"
	"time"
)

var _ = Describe("API", func() {
	const tipsCount = 10

	var (
		tips   []models.Tip
		db     *sql.DB
		mock   sqlmock.Sqlmock
		ctx    context.Context
		tipApi desc.OcpTipApiServer
		err    error
	)

	tips = make([]models.Tip, tipsCount)
	rand.Seed(time.Now().Unix())
	for i := 0; i < tipsCount; i++ {
		tips[i] = models.Tip{
			Id:        uint64(i + 1),
			UserId:    uint64(1 + rand.Intn(10)),
			ProblemId: uint64(1 + rand.Intn(10)),
			Text:      fmt.Sprintf("Tip number %d", i+1),
		}
	}

	BeforeEach(func() {
		db, mock, err = sqlmock.New()
		Expect(err).Should(BeNil())
		ctx = context.Background()
		tipApi = api.NewOcpTipApi(repo.NewRepo(sqlx.NewDb(db, "sqlmock")))
	})

	AfterEach(func() {
		mock.ExpectClose()
		err := db.Close()
		Expect(err).Should(BeNil())
	})

	Context("Create tip", func() {

		It("Successful create", func() {
			for _, tip := range tips {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(tip.Id)
				mock.ExpectQuery("INSERT INTO tips").WithArgs(tip.UserId, tip.ProblemId, tip.Text).WillReturnRows(rows)
				req := &desc.CreateTipV1Request{UserId: tip.UserId, ProblemId: tip.ProblemId, Text: tip.Text}
				res, err := tipApi.CreateTipV1(ctx, req)
				Expect(err).Should(BeNil())
				Expect(res.Id)
			}
		})

		It("Failed create", func() {
			tip := tips[0]
			mock.ExpectQuery("INSERT INTO tips").WithArgs(tip.UserId, tip.ProblemId, tip.Text).
				WillReturnError(errors.New("some error"))
			req := &desc.CreateTipV1Request{UserId: tip.UserId, ProblemId: tip.ProblemId, Text: tip.Text}
			res, err := tipApi.CreateTipV1(ctx, req)
			Expect(err.Error()).Should(ContainSubstring("some error"))
			Expect(res).Should(BeNil())
		})
	})

	Context("Describe tip", func() {
		It("Successful describe", func() {
			for _, tip := range tips {
				rows := sqlmock.NewRows([]string{"id", "user_id", "problem_id", "text"}).
					AddRow(tip.Id, tip.UserId, tip.ProblemId, tip.Text)
				mock.ExpectQuery("SELECT id, user_id, problem_id, text FROM tips").WithArgs(tip.Id).WillReturnRows(rows)
				req := &desc.DescribeTipV1Request{Id: tip.Id}
				res, err := tipApi.DescribeTipV1(ctx, req)
				Expect(err).Should(BeNil())
				Expect(res.Tip.Id).Should(Equal(tip.Id))
				Expect(res.Tip.UserId).Should(Equal(tip.UserId))
				Expect(res.Tip.ProblemId).Should(Equal(tip.ProblemId))
				Expect(res.Tip.Text).Should(Equal(tip.Text))
			}
		})

		It("Failed describe", func() {
			tip := tips[0]
			mock.ExpectQuery("SELECT id, user_id, problem_id, text FROM tips").WithArgs(tip.Id).
				WillReturnError(errors.New("some error"))
			req := &desc.DescribeTipV1Request{Id: tip.Id}
			res, err := tipApi.DescribeTipV1(ctx, req)
			Expect(err.Error()).Should(ContainSubstring("some error"))
			Expect(res).Should(BeNil())
		})

		It("Describe nonexistent tip", func() {
			tipId := uint64(tipsCount + 1)
			rows := sqlmock.NewRows([]string{"id", "user_id", "problem_id", "text"})
			mock.ExpectQuery("SELECT id, user_id, problem_id, text FROM tips").WithArgs(tipId).
				WillReturnRows(rows)
			req := &desc.DescribeTipV1Request{Id: tipId}
			res, err := tipApi.DescribeTipV1(ctx, req)
			Expect(err.Error()).Should(ContainSubstring("no rows in result set"))
			Expect(res).Should(BeNil())

		})
	})

	Context("List of tips", func() {
		var limit uint64 = tipsCount / 2
		var offset uint64 = tipsCount / 4
		req := &desc.ListTipsV1Request{Offset: offset, Limit: limit}

		It("Successful retrieve", func() {
			rows := sqlmock.NewRows([]string{"id", "user_id", "problem_id", "text"})
			for _, tip := range tips {
				rows = rows.AddRow(tip.Id, tip.UserId, tip.ProblemId, tip.Text)
			}
			mock.ExpectQuery(
				fmt.Sprintf("SELECT id, user_id, problem_id, text FROM tips LIMIT %d OFFSET %d", limit, offset)).
				WillReturnRows(rows)
			res, err := tipApi.ListTipsV1(ctx, req)
			Expect(err).Should(BeNil())
			for idx, responseTip := range res.Tips {
				Expect(responseTip.Id).Should(Equal(tips[idx].Id))
				Expect(responseTip.UserId).Should(Equal(tips[idx].UserId))
				Expect(responseTip.ProblemId).Should(Equal(tips[idx].ProblemId))
				Expect(responseTip.Text).Should(Equal(tips[idx].Text))
			}
		})

		It("Empty list", func() {
			mock.ExpectQuery(
				fmt.Sprintf("SELECT id, user_id, problem_id, text FROM tips LIMIT %d OFFSET %d", limit, offset)).
				WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "problem_id", "text"}))
			res, err := tipApi.ListTipsV1(ctx, req)
			Expect(err).Should(BeNil())
			Expect(res.Tips).Should(Equal([]*desc.TipV1{}))
		})

		It("Failed retrieve", func() {
			mock.ExpectQuery(
				fmt.Sprintf("SELECT id, user_id, problem_id, text FROM tips LIMIT %d OFFSET %d", limit, offset)).
				WillReturnError(errors.New("some error"))
			res, err := tipApi.ListTipsV1(ctx, req)
			Expect(err.Error()).Should(ContainSubstring("some error"))
			Expect(res).Should(BeNil())
		})
	})

	Context("Remove tip", func() {
		tip := tips[0]
		req := &desc.RemoveTipV1Request{Id: tip.Id}
		It("Successful remove, tip found", func() {
			mock.ExpectExec("DELETE FROM tips").WithArgs(tip.Id).WillReturnResult(sqlmock.NewResult(0, 1))
			res, err := tipApi.RemoveTipV1(ctx, req)
			Expect(err).Should(BeNil())
			Expect(res.Found).Should(Equal(true))
		})

		It("Successful remove, tip not found", func() {
			mock.ExpectExec("DELETE FROM tips").WithArgs(tip.Id).WillReturnResult(sqlmock.NewResult(0, 0))
			res, err := tipApi.RemoveTipV1(ctx, req)
			Expect(err).Should(BeNil())
			Expect(res.Found).Should(Equal(false))
		})

		It("Failed remove", func() {
			mock.ExpectExec("DELETE FROM tips").WithArgs(tip.Id).WillReturnError(errors.New("some error"))
			res, err := tipApi.RemoveTipV1(ctx, req)
			Expect(err.Error()).Should(ContainSubstring("some error"))
			Expect(res).Should(BeNil())
		})
	})
})

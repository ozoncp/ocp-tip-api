package api

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/ozoncp/ocp-tip-api/internal/metrics"
	"github.com/ozoncp/ocp-tip-api/internal/models"
	"github.com/ozoncp/ocp-tip-api/internal/utils"
	"github.com/ozoncp/ocp-tip-api/internal/version"
	desc "github.com/ozoncp/ocp-tip-api/pkg/ocp-tip-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"unsafe"
)

const multiCreateBatchSize int = 50

func (a *api) CreateTipV1(ctx context.Context, req *desc.CreateTipV1Request) (*desc.CreateTipV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	tipId, err := a.r.AddTip(ctx, models.Tip{UserId: req.UserId, ProblemId: req.ProblemId, Text: req.Text})

	if err != nil {
		return nil, err
	}

	_, _, sendErr := a.p.SendMessage(prepareMessage("create_tip", tipId))
	if sendErr != nil {
		log.Error().Err(sendErr)
	}
	metrics.IncCudCounter("create")
	return &desc.CreateTipV1Response{Id: tipId}, nil
}

func (a *api) MultiCreateTipV1(ctx context.Context, req *desc.MultiCreateTipV1Request) (*desc.MultiCreateTipV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("MultiCreateTipV1")
	defer span.Finish()

	tips := make([]models.Tip, 0, len(req.Tips))
	for _, reqTip := range req.Tips {
		tips = append(tips, models.Tip{UserId: reqTip.UserId, ProblemId: reqTip.ProblemId, Text: reqTip.Text})
	}
	errChan := make(chan []models.Tip)
	createChan := make(chan []uint64)
	defer func() {
		close(errChan)
		close(createChan)
	}()

	batches := utils.SplitTipsByBatches(tips, multiCreateBatchSize)
	for idx, batch := range batches {
		go func(i int, b []models.Tip) {
			childSpan := tracer.StartSpan(fmt.Sprintf("batch %d", i), opentracing.ChildOf(span.Context()))
			ids, err := a.r.AddTips(ctx, b)
			if err != nil {
				log.Error().Err(err)
				childSpan.SetTag("size", 0)
				errChan <- b
			} else {
				childSpan.SetTag("size", unsafe.Sizeof(b))
				createChan <- ids
			}
			childSpan.Finish()
		}(idx, batch)
	}

	createdIds := make([]uint64, 0, len(req.Tips))
	var notCreatedTips []*desc.MultiCreateFailedTipV1

	for range batches {
		select {
		case failedBatch := <-errChan:
			for _, tip := range failedBatch {
				notCreatedTips = append(notCreatedTips, &desc.MultiCreateFailedTipV1{
					UserId:    tip.UserId,
					ProblemId: tip.ProblemId,
					Text:      tip.Text,
				})
			}
		case ids := <-createChan:
			createdIds = append(createdIds, ids...)
		}
	}

	for _, tipId := range createdIds {
		_, _, sendErr := a.p.SendMessage(prepareMessage("create_tip", tipId))
		if sendErr != nil {
			log.Error().Err(sendErr)
		}
	}
	if len(notCreatedTips) == 0 {
		metrics.IncCudCounter("multi_create")
	}
	return &desc.MultiCreateTipV1Response{Ids: createdIds, NotCreatedTips: notCreatedTips}, nil
}

func (a *api) UpdateTipV1(ctx context.Context, req *desc.UpdateTipV1Request) (*desc.UpdateTipV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := a.r.UpdateTip(ctx, models.Tip{Id: req.Id, UserId: req.UserId, ProblemId: req.ProblemId, Text: req.Text})
	if err != nil {
		return nil, err
	}

	_, _, sendErr := a.p.SendMessage(prepareMessage("update_tip", req.Id))
	if sendErr != nil {
		log.Error().Err(sendErr)
	}
	metrics.IncCudCounter("update")
	return &desc.UpdateTipV1Response{}, nil
}

func (a *api) DescribeTipV1(ctx context.Context, req *desc.DescribeTipV1Request) (*desc.DescribeTipV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	tip, err := a.r.DescribeTip(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	log.Info().
		Uint64("Id", req.Id).
		Msg("describe tip")

	return &desc.DescribeTipV1Response{
		Tip: &desc.TipV1{
			Id: tip.Id, UserId: tip.UserId, ProblemId: tip.ProblemId, Text: tip.Text,
		},
	}, nil
}

func (a *api) ListTipsV1(ctx context.Context, req *desc.ListTipsV1Request) (*desc.ListTipsV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	tips, err := a.r.ListTips(ctx, req.Limit, req.Offset, req.SearchQuery)
	if err != nil {
		return nil, err
	}

	log.Info().
		Uint64("Limit", req.Limit).
		Uint64("Offset", req.Offset).
		Msg("list tips")

	responseTips := make([]*desc.TipV1, 0, len(tips))
	for _, tip := range tips {
		responseTips = append(responseTips, &desc.TipV1{
			Id: tip.Id, UserId: tip.UserId, ProblemId: tip.ProblemId, Text: tip.Text,
		})
	}
	return &desc.ListTipsV1Response{Tips: responseTips}, nil
}

func (a *api) RemoveTipV1(ctx context.Context, req *desc.RemoveTipV1Request) (*desc.RemoveTipV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	found, err := a.r.RemoveTip(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	_, _, sendErr := a.p.SendMessage(prepareMessage("delete_tip", req.Id))
	if sendErr != nil {
		log.Error().Err(sendErr)
	}
	metrics.IncCudCounter("delete")
	return &desc.RemoveTipV1Response{Found: found}, nil
}

func (a *api) ServiceInfoV1(ctx context.Context, req *desc.ServiceInfoV1Request) (*desc.ServiceInfoV1Response, error) {
	return &desc.ServiceInfoV1Response{
		Release:   version.Release,
		Commit:    version.Commit,
		BuildTime: version.BuildTime,
	}, nil
}

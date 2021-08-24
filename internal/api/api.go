package api

import (
	"context"
	"github.com/ozoncp/ocp-tip-api/internal/models"
	desc "github.com/ozoncp/ocp-tip-api/pkg/ocp-tip-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *api) CreateTipV1(ctx context.Context, req *desc.CreateTipV1Request) (*desc.CreateTipV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	tipId, err := a.r.AddTip(ctx, models.Tip{UserId: req.UserId, ProblemId: req.ProblemId, Text: req.Text})

	if err != nil {
		return nil, err
	}

	log.Info().
		Uint64("UserId", req.UserId).
		Uint64("ProblemId", req.ProblemId).
		Str("Text", req.Text).
		Msg("create tip")

	return &desc.CreateTipV1Response{Id: tipId}, nil
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

	tips, err := a.r.ListTips(ctx, req.Limit, req.Offset)
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

	log.Info().
		Uint64("Id", req.Id).
		Msg("remove tip")

	return &desc.RemoveTipV1Response{Found: found}, nil
}

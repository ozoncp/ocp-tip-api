package api

import (
	"context"
	desc "github.com/ozoncp/ocp-tip-api/pkg/ocp-tip-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *api) CreateTipV1(ctx context.Context, req *desc.CreateTipV1Request) (*desc.CreateTipV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info().
		Uint64("UserId", req.UserId).
		Uint64("ProblemId", req.ProblemId).
		Str("Text", req.Text).
		Msg("create tip")

	return &desc.CreateTipV1Response{Id: 1}, nil
}

func (a *api) DescribeTipV1(ctx context.Context, req *desc.DescribeTipV1Request) (*desc.DescribeTipV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info().
		Uint64("Id", req.Id).
		Msg("describe tip")

	return &desc.DescribeTipV1Response{
		Tip: &desc.TipV1{
			Id: req.Id, UserId: 1, ProblemId: 1, Text: "example",
		},
	}, nil
}

func (a *api) ListTipsV1(ctx context.Context, req *desc.ListTipsV1Request) (*desc.ListTipsV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info().
		Uint64("Limit", req.Limit).
		Uint64("Offset", req.Offset).
		Msg("list tips")

	return &desc.ListTipsV1Response{
		Tips: []*desc.TipV1{
			{Id: 1, UserId: 1, ProblemId: 1, Text: "example"},
		},
	}, nil
}

func (a *api) RemoveTipV1(ctx context.Context, req *desc.RemoveTipV1Request) (*desc.RemoveTipV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info().
		Uint64("Id", req.Id).
		Msg("remove tip")

	return &desc.RemoveTipV1Response{Found: true}, nil
}

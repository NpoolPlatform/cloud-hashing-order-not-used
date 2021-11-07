// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-order/message/npool"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/good-paying" //nolint

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateGoodPaying(ctx context.Context, in *npool.CreateGoodPayingRequest) (*npool.CreateGoodPayingResponse, error) {
	resp, err := goodpaying.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create good paying error: %w", err)
		return &npool.CreateGoodPayingResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateGoodPaying(ctx context.Context, in *npool.UpdateGoodPayingRequest) (*npool.UpdateGoodPayingResponse, error) {
	resp, err := goodpaying.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create good paying error: %w", err)
		return &npool.UpdateGoodPayingResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetGoodPaying(ctx context.Context, in *npool.GetGoodPayingRequest) (*npool.GetGoodPayingResponse, error) {
	resp, err := goodpaying.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create good paying error: %w", err)
		return &npool.GetGoodPayingResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

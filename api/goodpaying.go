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

func (s *Server) GetGoodPayingByOrder(ctx context.Context, in *npool.GetGoodPayingByOrderRequest) (*npool.GetGoodPayingByOrderResponse, error) {
	resp, err := goodpaying.GetByOrder(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create good paying by order error: %w", err)
		return &npool.GetGoodPayingByOrderResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

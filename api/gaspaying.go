// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-order/message/npool"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/gas-paying" //nolint

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateGasPaying(ctx context.Context, in *npool.CreateGasPayingRequest) (*npool.CreateGasPayingResponse, error) {
	resp, err := gaspaying.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create gas paying error: %w", err)
		return &npool.CreateGasPayingResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetGasPaying(ctx context.Context, in *npool.GetGasPayingRequest) (*npool.GetGasPayingResponse, error) {
	resp, err := gaspaying.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create gas paying error: %w", err)
		return &npool.GetGasPayingResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

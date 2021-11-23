// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-order/message/npool"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/out-of-gas" //nolint

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateOutOfGas(ctx context.Context, in *npool.CreateOutOfGasRequest) (*npool.CreateOutOfGasResponse, error) {
	resp, err := outofgas.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create out of gas error: %w", err)
		return &npool.CreateOutOfGasResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetOutOfGasesByOrder(ctx context.Context, in *npool.GetOutOfGasesByOrderRequest) (*npool.GetOutOfGasesByOrderResponse, error) {
	resp, err := outofgas.GetByOrder(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create out of gass by order error: %w", err)
		return &npool.GetOutOfGasesByOrderResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

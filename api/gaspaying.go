// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/gas-paying" //nolint
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-order"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateGasPaying(ctx context.Context, in *npool.CreateGasPayingRequest) (*npool.CreateGasPayingResponse, error) {
	resp, err := gaspaying.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create gas paying error: %w", err)
		return &npool.CreateGasPayingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetGasPayingsByOrder(ctx context.Context, in *npool.GetGasPayingsByOrderRequest) (*npool.GetGasPayingsByOrderResponse, error) {
	resp, err := gaspaying.GetByOrder(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create gas payings by order error: %w", err)
		return &npool.GetGasPayingsByOrderResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

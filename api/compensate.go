// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/compensate" //nolint
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-order"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCompensate(ctx context.Context, in *npool.CreateCompensateRequest) (*npool.CreateCompensateResponse, error) {
	resp, err := compensate.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create compensate error: %w", err)
		return &npool.CreateCompensateResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetCompensatesByOrder(ctx context.Context, in *npool.GetCompensatesByOrderRequest) (*npool.GetCompensatesByOrderResponse, error) {
	resp, err := compensate.GetByOrder(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create compensates by order error: %w", err)
		return &npool.GetCompensatesByOrderResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

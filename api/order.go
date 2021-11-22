// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-order/message/npool"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/order"          //nolint
	mw "github.com/NpoolPlatform/cloud-hashing-order/pkg/middleware/order" //nolint

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateOrder(ctx context.Context, in *npool.CreateOrderRequest) (*npool.CreateOrderResponse, error) {
	resp, err := order.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create order error: %w", err)
		return &npool.CreateOrderResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetOrder(ctx context.Context, in *npool.GetOrderRequest) (*npool.GetOrderResponse, error) {
	resp, err := order.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get order error: %w", err)
		return &npool.GetOrderResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetOrdersByAppUser(ctx context.Context, in *npool.GetOrdersByAppUserRequest) (*npool.GetOrdersByAppUserResponse, error) {
	resp, err := order.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get orders by app user error: %w", err)
		return &npool.GetOrdersByAppUserResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetOrdersByApp(ctx context.Context, in *npool.GetOrdersByAppRequest) (*npool.GetOrdersByAppResponse, error) {
	resp, err := order.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get orders by app error: %w", err)
		return &npool.GetOrdersByAppResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetOrdersByGood(ctx context.Context, in *npool.GetOrdersByGoodRequest) (*npool.GetOrdersByGoodResponse, error) {
	resp, err := order.GetByGood(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get orders by good error: %w", err)
		return &npool.GetOrdersByGoodResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetOrderDetail(ctx context.Context, in *npool.GetOrderDetailRequest) (*npool.GetOrderDetailResponse, error) {
	resp, err := mw.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get order error: %w", err)
		return &npool.GetOrderDetailResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetOrdersDetailByAppUser(ctx context.Context, in *npool.GetOrdersDetailByAppUserRequest) (*npool.GetOrdersDetailByAppUserResponse, error) {
	resp, err := mw.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get orders detail by app user error: %w", err)
		return &npool.GetOrdersDetailByAppUserResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetOrdersDetailByApp(ctx context.Context, in *npool.GetOrdersDetailByAppRequest) (*npool.GetOrdersDetailByAppResponse, error) {
	resp, err := mw.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get orders detail by app error: %w", err)
		return &npool.GetOrdersDetailByAppResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetOrdersDetailByGood(ctx context.Context, in *npool.GetOrdersDetailByGoodRequest) (*npool.GetOrdersDetailByGoodResponse, error) {
	resp, err := mw.GetByGood(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get orders detail by good error: %w", err)
		return &npool.GetOrdersDetailByGoodResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

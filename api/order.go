// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/order"          //nolint
	mw "github.com/NpoolPlatform/cloud-hashing-order/pkg/middleware/order" //nolint
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-order"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateOrder(ctx context.Context, in *npool.CreateOrderRequest) (*npool.CreateOrderResponse, error) {
	resp, err := order.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create order error: %v", err)
		return &npool.CreateOrderResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetOrder(ctx context.Context, in *npool.GetOrderRequest) (*npool.GetOrderResponse, error) {
	resp, err := order.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get order error: %v", err)
		return &npool.GetOrderResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetOrders(ctx context.Context, in *npool.GetOrdersRequest) (*npool.GetOrdersResponse, error) {
	resp, err := order.GetAll(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get orders error: %v", err)
		return &npool.GetOrdersResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetOrdersByAppUser(ctx context.Context, in *npool.GetOrdersByAppUserRequest) (*npool.GetOrdersByAppUserResponse, error) {
	resp, err := order.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get orders by app user error: %v", err)
		return &npool.GetOrdersByAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetOrderByAppUserCouponTypeID(ctx context.Context, in *npool.GetOrderByAppUserCouponTypeIDRequest) (*npool.GetOrderByAppUserCouponTypeIDResponse, error) {
	resp, err := order.GetByAppUserCouponTypeID(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get orders by app user coupon type id error: %v", err)
		return &npool.GetOrderByAppUserCouponTypeIDResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetOrdersByApp(ctx context.Context, in *npool.GetOrdersByAppRequest) (*npool.GetOrdersByAppResponse, error) {
	resp, err := order.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get orders by app error: %v", err)
		return &npool.GetOrdersByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetOrdersByOtherApp(ctx context.Context, in *npool.GetOrdersByOtherAppRequest) (*npool.GetOrdersByOtherAppResponse, error) {
	resp, err := order.GetByApp(ctx, &npool.GetOrdersByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorw("get orders by app error: %v", err)
		return &npool.GetOrdersByOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetOrdersByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}

func (s *Server) GetOrdersByGood(ctx context.Context, in *npool.GetOrdersByGoodRequest) (*npool.GetOrdersByGoodResponse, error) {
	resp, err := order.GetByGood(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get orders by good error: %v", err)
		return &npool.GetOrdersByGoodResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetOrderDetail(ctx context.Context, in *npool.GetOrderDetailRequest) (*npool.GetOrderDetailResponse, error) {
	resp, err := mw.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get order error: %v", err)
		return &npool.GetOrderDetailResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetOrdersDetailByAppUser(ctx context.Context, in *npool.GetOrdersDetailByAppUserRequest) (*npool.GetOrdersDetailByAppUserResponse, error) {
	resp, err := mw.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get orders detail by app user error: %v", err)
		return &npool.GetOrdersDetailByAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetOrdersShortDetailByAppUser(ctx context.Context, in *npool.GetOrdersShortDetailByAppUserRequest) (*npool.GetOrdersShortDetailByAppUserResponse, error) {
	resp, err := mw.GetShortByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get orders short detail by app user error: %v", err)
		return &npool.GetOrdersShortDetailByAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetOrdersDetailByApp(ctx context.Context, in *npool.GetOrdersDetailByAppRequest) (*npool.GetOrdersDetailByAppResponse, error) {
	resp, err := mw.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get orders detail by app error: %v", err)
		return &npool.GetOrdersDetailByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetOrdersDetailByGood(ctx context.Context, in *npool.GetOrdersDetailByGoodRequest) (*npool.GetOrdersDetailByGoodResponse, error) {
	resp, err := mw.GetByGood(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get orders detail by good error: %v", err)
		return &npool.GetOrdersDetailByGoodResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/payment" //nolint
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-order"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreatePayment(ctx context.Context, in *npool.CreatePaymentRequest) (*npool.CreatePaymentResponse, error) {
	resp, err := payment.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create payment error: %v", err)
		return &npool.CreatePaymentResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetPayment(ctx context.Context, in *npool.GetPaymentRequest) (*npool.GetPaymentResponse, error) {
	resp, err := payment.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get payment error: %v", err)
		return &npool.GetPaymentResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdatePayment(ctx context.Context, in *npool.UpdatePaymentRequest) (*npool.UpdatePaymentResponse, error) {
	resp, err := payment.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("update payment error: %v", err)
		return &npool.UpdatePaymentResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdatePaymentByUser(ctx context.Context, in *npool.UpdatePaymentByUserRequest) (*npool.UpdatePaymentByUserResponse, error) {
	resp, err := payment.UpdateByUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("update payment error: %v", err)
		return &npool.UpdatePaymentByUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetPaymentByOrder(ctx context.Context, in *npool.GetPaymentByOrderRequest) (*npool.GetPaymentByOrderResponse, error) {
	resp, err := payment.GetByOrder(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get payment by order error: %v", err)
		return &npool.GetPaymentByOrderResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetPaymentsByState(ctx context.Context, in *npool.GetPaymentsByStateRequest) (*npool.GetPaymentsByStateResponse, error) {
	resp, err := payment.GetByState(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get payments by state error: %v", err)
		return &npool.GetPaymentsByStateResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetPaymentsByApp(ctx context.Context, in *npool.GetPaymentsByAppRequest) (*npool.GetPaymentsByAppResponse, error) {
	resp, err := payment.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get payments by app error: %v", err)
		return &npool.GetPaymentsByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetPaymentsByOtherApp(ctx context.Context, in *npool.GetPaymentsByOtherAppRequest) (*npool.GetPaymentsByOtherAppResponse, error) {
	resp, err := payment.GetByApp(ctx, &npool.GetPaymentsByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorw("get payments by app error: %v", err)
		return &npool.GetPaymentsByOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetPaymentsByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}

func (s *Server) GetPaymentsByAppUser(ctx context.Context, in *npool.GetPaymentsByAppUserRequest) (*npool.GetPaymentsByAppUserResponse, error) {
	resp, err := payment.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get payments by app user error: %v", err)
		return &npool.GetPaymentsByAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetPayments(ctx context.Context, in *npool.GetPaymentsRequest) (*npool.GetPaymentsResponse, error) {
	resp, err := payment.GetAll(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get payments error: %v", err)
		return &npool.GetPaymentsResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

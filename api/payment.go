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

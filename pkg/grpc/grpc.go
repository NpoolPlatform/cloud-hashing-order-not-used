package grpc

import (
	"context"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	coininfopb "github.com/NpoolPlatform/message/npool/coininfo"
	coininfoconst "github.com/NpoolPlatform/sphinx-coininfo/pkg/message/const" //nolint

	tradingpb "github.com/NpoolPlatform/message/npool/trading"
	tradingconst "github.com/NpoolPlatform/sphinx-service/pkg/message/const" //nolint

	billingpb "github.com/NpoolPlatform/cloud-hashing-billing/message/npool"
	billingconst "github.com/NpoolPlatform/cloud-hashing-billing/pkg/message/const" //nolint

	"golang.org/x/xerrors"
)

func GetCoinInfo(ctx context.Context, in *coininfopb.GetCoinInfoRequest) (*coininfopb.GetCoinInfoResponse, error) {
	conn, err := grpc2.GetGRPCConn(coininfoconst.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, xerrors.Errorf("fail get coininfo connection: %v", err)
	}

	cli := coininfopb.NewSphinxCoinInfoClient(conn)
	return cli.GetCoinInfo(ctx, in)
}

//--------------------------------------------------------------------------------------------------------------------------

func GetWalletBalance(ctx context.Context, in *tradingpb.GetWalletBalanceRequest) (*tradingpb.GetWalletBalanceResponse, error) {
	conn, err := grpc2.GetGRPCConn(tradingconst.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, xerrors.Errorf("fail get trading connection: %v", err)
	}

	cli := tradingpb.NewTradingClient(conn)
	return cli.GetWalletBalance(ctx, in)
}

//--------------------------------------------------------------------------------------------------------------------------

func GetBillingAccount(ctx context.Context, in *billingpb.GetCoinAccountRequest) (*billingpb.GetCoinAccountResponse, error) {
	conn, err := grpc2.GetGRPCConn(billingconst.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, xerrors.Errorf("fail get billing connection: %v", err)
	}

	cli := billingpb.NewCloudHashingBillingClient(conn)
	return cli.GetCoinAccount(ctx, in)
}

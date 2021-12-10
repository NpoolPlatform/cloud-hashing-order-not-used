package grpc

import (
	"context"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	coininfopb "github.com/NpoolPlatform/message/npool/coininfo"
	coininfoconst "github.com/NpoolPlatform/sphinx-coininfo/pkg/message/const" //nolint

	sphinxproxypb "github.com/NpoolPlatform/message/npool/sphinxproxy"
	sphinxproxyconst "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const" //nolint

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

func GetBalance(ctx context.Context, in *sphinxproxypb.GetBalanceRequest) (*sphinxproxypb.GetBalanceResponse, error) {
	conn, err := grpc2.GetGRPCConn(sphinxproxyconst.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, xerrors.Errorf("fail get sphinxproxy connection: %v", err)
	}

	cli := sphinxproxypb.NewSphinxProxyClient(conn)
	return cli.GetBalance(ctx, in)
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

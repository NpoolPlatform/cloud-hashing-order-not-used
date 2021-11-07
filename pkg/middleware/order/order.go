package order

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-order/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/gas-paying"  //nolint
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/good-paying" //nolint
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/order"

	"golang.org/x/xerrors"
)

func constructOrderDetail(info *npool.Order, goodPaying *npool.GoodPaying, gasPayings []*npool.GasPaying) *npool.OrderDetail {
	return &npool.OrderDetail{
		ID:                       info.ID,
		GoodID:                   info.GoodID,
		Units:                    info.Units,
		Discount:                 info.Discount,
		SpecialReductionAmount:   info.SpecialReductionAmount,
		UserID:                   info.UserID,
		AppID:                    info.AppID,
		State:                    info.State,
		GoodPaying:               goodPaying,
		Start:                    info.Start,
		End:                      info.End,
		CompensateMinutes:        info.CompensateMinutes,
		CompensateElapsedMinutes: info.CompensateElapsedMinutes,
		GasStart:                 info.GasStart,
		GasEnd:                   info.GasEnd,
		GasPayings:               gasPayings,
		CouponID:                 info.CouponID,
	}
}

func Get(ctx context.Context, in *npool.GetOrderDetailRequest) (*npool.GetOrderDetailResponse, error) {
	info, err := order.Get(ctx, &npool.GetOrderRequest{
		ID: in.GetID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get order: %v", err)
	}

	goodPaying, err := goodpaying.Get(ctx, &npool.GetGoodPayingRequest{
		ID: info.Info.GoodPayID,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get good pay: %v", err)
	}

	gasPayings := []*npool.GasPaying{}
	for _, gasPayID := range info.Info.GasPayIDs {
		gasPaying, err := gaspaying.Get(ctx, &npool.GetGasPayingRequest{
			ID: gasPayID,
		})
		if err != nil {
			return nil, xerrors.Errorf("fail get gas pay: %v", err)
		}
		gasPayings = append(gasPayings, gasPaying.Info)
	}

	return &npool.GetOrderDetailResponse{
		Detail: constructOrderDetail(info.Info, goodPaying.Info, gasPayings),
	}, nil
}

func GetOrdersDetailByAppUser(ctx context.Context, in *npool.GetOrdersDetailByAppUserRequest) (*npool.GetOrdersDetailByAppUserResponse, error) {
	return nil, nil
}

func GetOrdersDetailByApp(ctx context.Context, in *npool.GetOrdersDetailByAppRequest) (*npool.GetOrdersDetailByAppResponse, error) {
	return nil, nil
}

func GetOrdersDetailByGood(ctx context.Context, in *npool.GetOrdersDetailByGoodRequest) (*npool.GetOrdersDetailByGoodResponse, error) {
	return nil, nil
}

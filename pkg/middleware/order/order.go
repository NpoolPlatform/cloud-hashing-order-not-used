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

func getOrderDetail(ctx context.Context, info *npool.Order) (*npool.OrderDetail, error) {
	goodPaying, err := goodpaying.Get(ctx, &npool.GetGoodPayingRequest{
		ID: info.GoodPayID,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get good pay: %v", err)
	}

	gasPayings := []*npool.GasPaying{}
	for _, gasPayID := range info.GasPayIDs {
		gasPaying, err := gaspaying.Get(ctx, &npool.GetGasPayingRequest{
			ID: gasPayID,
		})
		if err != nil {
			return nil, xerrors.Errorf("fail get gas pay: %v", err)
		}
		gasPayings = append(gasPayings, gasPaying.Info)
	}

	return constructOrderDetail(info, goodPaying.Info, gasPayings), nil
}

func Get(ctx context.Context, in *npool.GetOrderDetailRequest) (*npool.GetOrderDetailResponse, error) {
	info, err := order.Get(ctx, &npool.GetOrderRequest{
		ID: in.GetID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get order: %v", err)
	}

	detail, err := getOrderDetail(ctx, info.Info)
	if err != nil {
		return nil, xerrors.Errorf("fail get order detail: %v", err)
	}

	return &npool.GetOrderDetailResponse{
		Detail: detail,
	}, nil
}

func getOrdersDetail(ctx context.Context, infos []*npool.Order) ([]*npool.OrderDetail, error) {
	details := []*npool.OrderDetail{}
	for _, info := range infos {
		detail, err := getOrderDetail(ctx, info)
		if err != nil {
			return nil, xerrors.Errorf("fail get order detail: %v", err)
		}
		details = append(details, detail)
	}
	return details, nil
}

func GetByAppUser(ctx context.Context, in *npool.GetOrdersDetailByAppUserRequest) (*npool.GetOrdersDetailByAppUserResponse, error) {
	resp, err := order.GetByAppUser(ctx, &npool.GetOrdersByAppUserRequest{
		AppID:  in.GetAppID(),
		UserID: in.GetUserID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get orders by app user: %v", err)
	}

	details, err := getOrdersDetail(ctx, resp.Infos)
	if err != nil {
		return nil, xerrors.Errorf("fail get orders detail: %v", err)
	}

	return &npool.GetOrdersDetailByAppUserResponse{
		Details: details,
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetOrdersDetailByAppRequest) (*npool.GetOrdersDetailByAppResponse, error) {
	resp, err := order.GetByApp(ctx, &npool.GetOrdersByAppRequest{
		AppID: in.GetAppID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get orders by app: %v", err)
	}

	details, err := getOrdersDetail(ctx, resp.Infos)
	if err != nil {
		return nil, xerrors.Errorf("fail get orders detail: %v", err)
	}

	return &npool.GetOrdersDetailByAppResponse{
		Details: details,
	}, nil
}

func GetByGood(ctx context.Context, in *npool.GetOrdersDetailByGoodRequest) (*npool.GetOrdersDetailByGoodResponse, error) {
	resp, err := order.GetByGood(ctx, &npool.GetOrdersByGoodRequest{
		GoodID: in.GetGoodID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get orders by app: %v", err)
	}

	details, err := getOrdersDetail(ctx, resp.Infos)
	if err != nil {
		return nil, xerrors.Errorf("fail get orders detail: %v", err)
	}

	return &npool.GetOrdersDetailByGoodResponse{
		Details: details,
	}, nil
}

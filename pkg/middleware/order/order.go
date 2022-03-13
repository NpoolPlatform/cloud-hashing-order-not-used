package order

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	constant "github.com/NpoolPlatform/cloud-hashing-order/pkg/const"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-order"

	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/compensate"  //nolint
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/gas-paying"  //nolint
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/good-paying" //nolint
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/order"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/out-of-gas" //nolint
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/payment"    //nolint

	"golang.org/x/xerrors"
)

func constructOrderDetail(
	info *npool.Order,
	goodPaying *npool.GoodPaying,
	gasPayings []*npool.GasPaying,
	compensates []*npool.Compensate,
	outOfGases []*npool.OutOfGas,
	goodPayment *npool.Payment) *npool.OrderDetail {
	return &npool.OrderDetail{
		Order:       info,
		GoodPaying:  goodPaying,
		GasPayings:  gasPayings,
		Compensates: compensates,
		OutOfGases:  outOfGases,
		Payment:     goodPayment,
	}
}

func getOrderDetail(ctx context.Context, info *npool.Order, short bool) (*npool.OrderDetail, error) {
	var paymentInfo *npool.Payment

	goodPayment, err := payment.GetByOrder(ctx, &npool.GetPaymentByOrderRequest{
		OrderID: info.ID,
	})
	if err != nil {
		logger.Sugar().Warnf("cannot find payment for order: %v", err)
	} else {
		paymentInfo = goodPayment.Info
	}

	var goodPayingInfo *npool.GoodPaying

	if !short {
		goodPaying, err := goodpaying.GetByOrder(ctx, &npool.GetGoodPayingByOrderRequest{
			OrderID: info.ID,
		})
		if err != nil {
			logger.Sugar().Warnf("cannot find good paying for order: %v", err)
		} else {
			goodPayingInfo = goodPaying.Info
		}
	}

	gasPayingInfos := []*npool.GasPaying{}

	if !short {
		gasPayings, err := gaspaying.GetByOrder(ctx, &npool.GetGasPayingsByOrderRequest{
			OrderID: info.ID,
		})
		if err != nil {
			logger.Sugar().Warnf("cannot find gas paying for order: %v", err)
		} else {
			gasPayingInfos = gasPayings.Infos
		}
	}

	var compensateInfos []*npool.Compensate

	if !short {
		compensates, err := compensate.GetByOrder(ctx, &npool.GetCompensatesByOrderRequest{
			OrderID: info.ID,
		})
		if err != nil {
			return nil, xerrors.Errorf("fail to get compensates for order: %v", err)
		}
		compensateInfos = compensates.Infos
	}

	var outOfGasInfos []*npool.OutOfGas

	if !short {
		outOfGases, err := outofgas.GetByOrder(ctx, &npool.GetOutOfGasesByOrderRequest{
			OrderID: info.ID,
		})
		if err != nil {
			return nil, xerrors.Errorf("fail to get out of gas for order: %v", err)
		}
		outOfGasInfos = outOfGases.Infos
	}

	return constructOrderDetail(
		info,
		goodPayingInfo,
		gasPayingInfos,
		compensateInfos,
		outOfGasInfos,
		paymentInfo), nil
}

func Get(ctx context.Context, in *npool.GetOrderDetailRequest) (*npool.GetOrderDetailResponse, error) {
	info, err := order.Get(ctx, &npool.GetOrderRequest{
		ID: in.GetID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get order: %v", err)
	}

	detail, err := getOrderDetail(ctx, info.Info, false)
	if err != nil {
		return nil, xerrors.Errorf("fail get order detail: %v", err)
	}

	return &npool.GetOrderDetailResponse{
		Info: detail,
	}, nil
}

func getOrdersDetail(ctx context.Context, infos []*npool.Order, short bool) ([]*npool.OrderDetail, error) {
	details := []*npool.OrderDetail{}
	for _, info := range infos {
		detail, err := getOrderDetail(ctx, info, short)
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

	details, err := getOrdersDetail(ctx, resp.Infos, false)
	if err != nil {
		return nil, xerrors.Errorf("fail get orders detail: %v", err)
	}

	return &npool.GetOrdersDetailByAppUserResponse{
		Infos: details,
	}, nil
}

func GetShortByAppUser(ctx context.Context, in *npool.GetOrdersShortDetailByAppUserRequest) (*npool.GetOrdersShortDetailByAppUserResponse, error) {
	resp, err := order.GetByAppUser(ctx, &npool.GetOrdersByAppUserRequest{
		AppID:  in.GetAppID(),
		UserID: in.GetUserID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get orders short by app user: %v", err)
	}

	details, err := getOrdersDetail(ctx, resp.Infos, true)
	if err != nil {
		return nil, xerrors.Errorf("fail get orders detail: %v", err)
	}

	return &npool.GetOrdersShortDetailByAppUserResponse{
		Infos: details,
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetOrdersDetailByAppRequest) (*npool.GetOrdersDetailByAppResponse, error) {
	resp, err := order.GetByApp(ctx, &npool.GetOrdersByAppRequest{
		AppID: in.GetAppID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get orders by app: %v", err)
	}

	details, err := getOrdersDetail(ctx, resp.Infos, false)
	if err != nil {
		return nil, xerrors.Errorf("fail get orders detail: %v", err)
	}

	return &npool.GetOrdersDetailByAppResponse{
		Infos: details,
	}, nil
}

func GetByGood(ctx context.Context, in *npool.GetOrdersDetailByGoodRequest) (*npool.GetOrdersDetailByGoodResponse, error) {
	resp, err := order.GetByGood(ctx, &npool.GetOrdersByGoodRequest{
		GoodID: in.GetGoodID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get orders by good: %v", err)
	}

	details, err := getOrdersDetail(ctx, resp.Infos, false)
	if err != nil {
		return nil, xerrors.Errorf("fail get orders detail: %v", err)
	}

	return &npool.GetOrdersDetailByGoodResponse{
		Infos: details,
	}, nil
}

func SoldByGood(ctx context.Context, in *npool.GetSoldByGoodRequest) (*npool.GetSoldByGoodResponse, error) {
	resp, err := order.GetByGood(ctx, &npool.GetOrdersByGoodRequest{
		GoodID: in.GetGoodID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get orders by good: %v", err)
	}

	units := uint32(0)

	for _, info := range resp.Infos {
		pay, err := payment.GetByOrder(ctx, &npool.GetPaymentByOrderRequest{
			OrderID: info.ID,
		})
		if err != nil || pay.Info == nil {
			continue
		}

		if pay.Info.State == constant.PaymentStateCanceled || pay.Info.State == constant.PaymentStateTimeout {
			continue
		}
		if info.End < uint32(time.Now().Unix()) {
			continue
		}

		units += info.Units
	}

	return &npool.GetSoldByGoodResponse{
		Sold: units,
	}, nil
}

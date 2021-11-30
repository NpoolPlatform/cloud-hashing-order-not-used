package order

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-order/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/order"

	"github.com/NpoolPlatform/go-service-framework/pkg/price"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateOrder(info *npool.Order) error {
	if _, err := uuid.Parse(info.GetGoodID()); err != nil {
		return xerrors.Errorf("invalid good id: %v", err)
	}
	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		return xerrors.Errorf("invalid user id: %v", err)
	}
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	return nil
}

func dbRowToOrder(row *ent.Order) *npool.Order {
	return &npool.Order{
		ID:                     row.ID.String(),
		GoodID:                 row.GoodID.String(),
		Units:                  row.Units,
		Discount:               row.Discount,
		SpecialReductionAmount: price.DBPriceToVisualPrice(row.SpecialReductionAmount),
		UserID:                 row.UserID.String(),
		AppID:                  row.AppID.String(),
		Start:                  row.Start,
		End:                    row.End,
		CouponID:               row.CouponID.String(),
	}
}

func Create(ctx context.Context, in *npool.CreateOrderRequest) (*npool.CreateOrderResponse, error) {
	if err := validateOrder(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	couponID, err := uuid.Parse(in.GetInfo().GetCouponID())
	if err != nil {
		couponID = uuid.UUID{}
	}

	info, err := db.Client().
		Order.
		Create().
		SetGoodID(uuid.MustParse(in.GetInfo().GetGoodID())).
		SetUnits(in.GetInfo().GetUnits()).
		SetDiscount(in.GetInfo().GetDiscount()).
		SetSpecialReductionAmount(price.VisualPriceToDBPrice(in.GetInfo().GetSpecialReductionAmount())).
		SetUserID(uuid.MustParse(in.GetInfo().GetUserID())).
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetStart(in.GetInfo().GetStart()).
		SetEnd(in.GetInfo().GetEnd()).
		SetCouponID(couponID).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create order: %v", err)
	}

	return &npool.CreateOrderResponse{
		Info: dbRowToOrder(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetOrderRequest) (*npool.GetOrderResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	infos, err := db.Client().
		Order.
		Query().
		Where(
			order.And(
				order.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query order: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty order")
	}

	return &npool.GetOrderResponse{
		Info: dbRowToOrder(infos[0]),
	}, nil
}

func GetByAppUser(ctx context.Context, in *npool.GetOrdersByAppUserRequest) (*npool.GetOrdersByAppUserResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	userID, err := uuid.Parse(in.GetUserID())
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	infos, err := db.Client().
		Order.
		Query().
		Where(
			order.And(
				order.AppID(appID),
				order.UserID(userID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query order: %v", err)
	}

	orders := []*npool.Order{}
	for _, info := range infos {
		orders = append(orders, dbRowToOrder(info))
	}

	return &npool.GetOrdersByAppUserResponse{
		Infos: orders,
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetOrdersByAppRequest) (*npool.GetOrdersByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	infos, err := db.Client().
		Order.
		Query().
		Where(
			order.And(
				order.AppID(appID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query order: %v", err)
	}

	orders := []*npool.Order{}
	for _, info := range infos {
		orders = append(orders, dbRowToOrder(info))
	}

	return &npool.GetOrdersByAppResponse{
		Infos: orders,
	}, nil
}

func GetByGood(ctx context.Context, in *npool.GetOrdersByGoodRequest) (*npool.GetOrdersByGoodResponse, error) {
	goodID, err := uuid.Parse(in.GetGoodID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good id: %v", err)
	}

	infos, err := db.Client().
		Order.
		Query().
		Where(
			order.And(
				order.GoodID(goodID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query order: %v", err)
	}

	orders := []*npool.Order{}
	for _, info := range infos {
		orders = append(orders, dbRowToOrder(info))
	}

	return &npool.GetOrdersByGoodResponse{
		Infos: orders,
	}, nil
}

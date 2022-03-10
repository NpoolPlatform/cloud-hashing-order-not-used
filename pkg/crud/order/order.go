package order

import (
	"context"

	constant "github.com/NpoolPlatform/cloud-hashing-order/pkg/const"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-order"

	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/order"

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
		DiscountCouponID:       row.DiscountCouponID.String(),
		UserSpecialReductionID: row.UserSpecialReductionID.String(),
		UserID:                 row.UserID.String(),
		AppID:                  row.AppID.String(),
		Start:                  row.Start,
		End:                    row.End,
		CouponID:               row.CouponID.String(),
		PromotionID:            row.PromotionID.String(),
		CreateAt:               row.CreateAt,
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

	discountCouponID, err := uuid.Parse(in.GetInfo().GetDiscountCouponID())
	if err != nil {
		discountCouponID = uuid.UUID{}
	}

	promotionID, err := uuid.Parse(in.GetInfo().GetPromotionID())
	if err != nil {
		promotionID = uuid.UUID{}
	}

	userSpecialReductionID, err := uuid.Parse(in.GetInfo().GetUserSpecialReductionID())
	if err != nil {
		userSpecialReductionID = uuid.UUID{}
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		Order.
		Create().
		SetGoodID(uuid.MustParse(in.GetInfo().GetGoodID())).
		SetUnits(in.GetInfo().GetUnits()).
		SetDiscountCouponID(discountCouponID).
		SetUserSpecialReductionID(userSpecialReductionID).
		SetUserID(uuid.MustParse(in.GetInfo().GetUserID())).
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetStart(in.GetInfo().GetStart()).
		SetEnd(in.GetInfo().GetEnd()).
		SetCouponID(couponID).
		SetPromotionID(promotionID).
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

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		Order.
		Query().
		Where(
			order.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query order: %v", err)
	}

	var _order *npool.Order
	for _, info := range infos {
		_order = dbRowToOrder(info)
	}

	return &npool.GetOrderResponse{
		Info: _order,
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

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
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

func GetByAppUserCouponTypeID(ctx context.Context, in *npool.GetOrderByAppUserCouponTypeIDRequest) (*npool.GetOrderByAppUserCouponTypeIDResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	userID, err := uuid.Parse(in.GetUserID())
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	couponID, err := uuid.Parse(in.GetCouponID())
	if err != nil {
		return nil, xerrors.Errorf("invalid coupon id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	query := cli.Order.Query()

	switch in.GetCouponType() {
	case constant.FixAmountCoupon:
		query = query.Where(
			order.And(
				order.AppID(appID),
				order.UserID(userID),
				order.CouponID(couponID),
			),
		)
	case constant.DiscountCoupon:
		query = query.Where(
			order.And(
				order.AppID(appID),
				order.UserID(userID),
				order.DiscountCouponID(couponID),
			),
		)
	case constant.UserSpecialReductionCoupon:
		query = query.Where(
			order.And(
				order.AppID(appID),
				order.UserID(userID),
				order.UserSpecialReductionID(couponID),
			),
		)
	}
	infos, err := query.All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query order: %v", err)
	}

	var _order *npool.Order
	for _, info := range infos {
		_order = dbRowToOrder(info)
		break
	}

	return &npool.GetOrderByAppUserCouponTypeIDResponse{
		Info: _order,
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetOrdersByAppRequest) (*npool.GetOrdersByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		Order.
		Query().
		Where(
			order.AppID(appID),
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

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		Order.
		Query().
		Where(
			order.GoodID(goodID),
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

func SoldByGood(ctx context.Context, in *npool.GetSoldByGoodRequest) (*npool.GetSoldByGoodResponse, error) {
	goodID, err := uuid.Parse(in.GetGoodID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	var units []struct {
		GoodID string `json:"good_id"`
		Units  uint32 `json:"sum"`
	}

	err = cli.
		Order.
		Query().
		Where(
			order.GoodID(goodID),
		).
		GroupBy(order.FieldGoodID).
		Aggregate(ent.Sum(order.FieldUnits)).
		Scan(ctx, &units)
	if err != nil {
		return nil, xerrors.Errorf("fail count good order: %v", err)
	}

	return &npool.GetSoldByGoodResponse{
		Sold: units[0].Units,
	}, nil
}

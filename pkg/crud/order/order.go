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
	if _, err := uuid.Parse(info.GetGoodPayID()); err != nil {
		return xerrors.Errorf("invalid good pay id: %v", err)
	}
	if _, err := uuid.Parse(info.GetCouponID()); err != nil {
		return xerrors.Errorf("invalid coupon id: %v", err)
	}
	for _, payID := range info.GetGasPayIDs() {
		if _, err := uuid.Parse(payID); err != nil {
			return xerrors.Errorf("invalid gas pay id: %v", err)
		}
	}
	return nil
}

func dbRowToOrder(row *ent.Order) *npool.Order {
	gasPayIDs := []string{}
	for _, payID := range row.GasPayIds {
		gasPayIDs = append(gasPayIDs, payID.String())
	}

	return &npool.Order{
		ID:                       row.ID.String(),
		GoodID:                   row.GoodID.String(),
		Units:                    row.Units,
		Discount:                 row.Discount,
		SpecialReductionAmount:   price.DBPriceToVisualPrice(row.SpecialReductionAmount),
		UserID:                   row.UserID.String(),
		AppID:                    row.AppID.String(),
		State:                    string(row.State),
		GoodPayID:                row.GoodPayID.String(),
		Start:                    row.Start,
		End:                      row.End,
		CompensateMinutes:        row.CompensateMinutes,
		CompensateElapsedMinutes: row.CompensateElapsedMinutes,
		GasStart:                 row.GasStart,
		GasEnd:                   row.GasEnd,
		GasPayIDs:                gasPayIDs,
		CouponID:                 row.CouponID.String(),
	}
}

func Create(ctx context.Context, in *npool.CreateOrderRequest) (*npool.CreateOrderResponse, error) {
	if err := validateOrder(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	gasPayIDs := []uuid.UUID{}
	for _, gasPayID := range in.GetInfo().GetGasPayIDs() {
		gasPayIDs = append(gasPayIDs, uuid.MustParse(gasPayID))
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
		SetState(order.State(in.GetInfo().GetState())).
		SetGoodPayID(uuid.MustParse(in.GetInfo().GetGoodPayID())).
		SetStart(in.GetInfo().GetStart()).
		SetEnd(in.GetInfo().GetEnd()).
		SetCompensateMinutes(in.GetInfo().GetCompensateMinutes()).
		SetCompensateElapsedMinutes(in.GetInfo().GetCompensateElapsedMinutes()).
		SetGasStart(in.GetInfo().GetGasStart()).
		SetGasEnd(in.GetInfo().GetGasEnd()).
		SetGasPayIds(gasPayIDs).
		SetCouponID(uuid.MustParse(in.GetInfo().GetCouponID())).
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

func Update(ctx context.Context, in *npool.UpdateOrderRequest) (*npool.UpdateOrderResponse, error) {
	if err := validateOrder(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	gasPayIDs := []uuid.UUID{}
	for _, gasPayID := range in.GetInfo().GetGasPayIDs() {
		gasPayIDs = append(gasPayIDs, uuid.MustParse(gasPayID))
	}

	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	oldGasPayIDs, err := db.Client().
		Order.
		Query().
		Where(
			order.And(
				order.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query old gas pay id: %v", err)
	}
	if len(oldGasPayIDs) == 0 {
		return nil, xerrors.Errorf("empty gas pay ids")
	}

	for _, oldGasPayID := range oldGasPayIDs[0].GasPayIds {
		found := false
		for _, gasPayID := range gasPayIDs {
			if gasPayID == oldGasPayID {
				found = true
				break
			}
		}
		if found {
			continue
		}
		gasPayIDs = append(gasPayIDs, oldGasPayID)
	}

	info, err := db.Client().
		Order.
		UpdateOneID(id).
		SetState(order.State(in.GetInfo().GetState())).
		SetGasPayIds(gasPayIDs).
		SetGasEnd(in.GetInfo().GetGasEnd()).
		SetCompensateMinutes(in.GetInfo().GetCompensateMinutes()).
		SetCompensateElapsedMinutes(in.GetInfo().GetCompensateElapsedMinutes()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update order: %v", err)
	}

	return &npool.UpdateOrderResponse{
		Info: dbRowToOrder(info),
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
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty order")
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
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty order")
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
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty order")
	}

	orders := []*npool.Order{}
	for _, info := range infos {
		orders = append(orders, dbRowToOrder(info))
	}

	return &npool.GetOrdersByGoodResponse{
		Infos: orders,
	}, nil
}

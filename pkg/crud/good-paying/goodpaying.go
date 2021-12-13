package goodpaying

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-order/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/goodpaying"

	"github.com/google/uuid" //nolint

	"golang.org/x/xerrors" //nolint
)

func validateGoodPaying(info *npool.GoodPaying) error {
	if _, err := uuid.Parse(info.GetOrderID()); err != nil {
		return xerrors.Errorf("invalid order id: %v", err)
	}
	if _, err := uuid.Parse(info.GetPaymentID()); err != nil {
		return xerrors.Errorf("invalid payment id: %v", err)
	}
	return nil
}

func dbRowToGoodPaying(row *ent.GoodPaying) *npool.GoodPaying {
	return &npool.GoodPaying{
		ID:        row.ID.String(),
		OrderID:   row.OrderID.String(),
		PaymentID: row.PaymentID.String(),
	}
}

func Create(ctx context.Context, in *npool.CreateGoodPayingRequest) (*npool.CreateGoodPayingResponse, error) {
	if err := validateGoodPaying(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	info, err := db.Client().
		GoodPaying.
		Create().
		SetOrderID(uuid.MustParse(in.GetInfo().GetOrderID())).
		SetPaymentID(uuid.MustParse(in.GetInfo().GetPaymentID())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create good paying: %v", err)
	}

	return &npool.CreateGoodPayingResponse{
		Info: dbRowToGoodPaying(info),
	}, nil
}

func GetByOrder(ctx context.Context, in *npool.GetGoodPayingByOrderRequest) (*npool.GetGoodPayingByOrderResponse, error) {
	orderID, err := uuid.Parse(in.GetOrderID())
	if err != nil {
		return nil, xerrors.Errorf("invalid order id: %v", err)
	}

	infos, err := db.Client().
		GoodPaying.
		Query().
		Where(
			goodpaying.And(
				goodpaying.OrderID(orderID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query good paying: %v", err)
	}

	var paying *npool.GoodPaying
	for _, info := range infos {
		paying = dbRowToGoodPaying(info)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty good paying")
	}

	return &npool.GetGoodPayingByOrderResponse{
		Info: paying,
	}, nil
}

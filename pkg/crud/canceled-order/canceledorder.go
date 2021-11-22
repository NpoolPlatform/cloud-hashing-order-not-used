package canceledorder

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-order/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/canceledorder"

	"github.com/google/uuid" //nolint

	"golang.org/x/xerrors" //nolint
)

func validateCanceledOrder(info *npool.CanceledOrder) error {
	if _, err := uuid.Parse(info.GetOrderID()); err != nil {
		return xerrors.Errorf("invalid order id: %v", err)
	}
	if _, err := uuid.Parse(info.GetPaymentID()); err != nil {
		return xerrors.Errorf("invalid payment id: %v", err)
	}
	return nil
}

func dbRowToCanceledOrder(row *ent.CanceledOrder) *npool.CanceledOrder {
	return &npool.CanceledOrder{
		ID:        row.ID.String(),
		OrderID:   row.OrderID.String(),
		PaymentID: row.PaymentID.String(),
	}
}

func Create(ctx context.Context, in *npool.CreateCanceledOrderRequest) (*npool.CreateCanceledOrderResponse, error) {
	if err := validateCanceledOrder(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	info, err := db.Client().
		CanceledOrder.
		Create().
		SetOrderID(uuid.MustParse(in.GetInfo().GetOrderID())).
		SetPaymentID(uuid.MustParse(in.GetInfo().GetPaymentID())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create canceled order: %v", err)
	}

	return &npool.CreateCanceledOrderResponse{
		Info: dbRowToCanceledOrder(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetCanceledOrderRequest) (*npool.GetCanceledOrderResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	infos, err := db.Client().
		CanceledOrder.
		Query().
		Where(
			canceledorder.And(
				canceledorder.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query canceled order: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty canceled order")
	}

	return &npool.GetCanceledOrderResponse{
		Info: dbRowToCanceledOrder(infos[0]),
	}, nil
}

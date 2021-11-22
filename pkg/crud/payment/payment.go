package payment

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-order/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/payment"

	"github.com/google/uuid" //nolint

	"golang.org/x/xerrors" //nolint
)

func validatePayment(info *npool.Payment) error {
	if _, err := uuid.Parse(info.GetOrderID()); err != nil {
		return xerrors.Errorf("invalid order id: %v", err)
	}
	if _, err := uuid.Parse(info.GetPaymentID()); err != nil {
		return xerrors.Errorf("invalid payment id: %v", err)
	}
	return nil
}

func dbRowToPayment(row *ent.Payment) *npool.Payment {
	return &npool.Payment{
		ID:        row.ID.String(),
		OrderID:   row.OrderID.String(),
		PaymentID: row.PaymentID.String(),
	}
}

func Create(ctx context.Context, in *npool.CreatePaymentRequest) (*npool.CreatePaymentResponse, error) {
	if err := validatePayment(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	info, err := db.Client().
		Payment.
		Create().
		SetOrderID(uuid.MustParse(in.GetInfo().GetOrderID())).
		SetPaymentID(uuid.MustParse(in.GetInfo().GetPaymentID())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create good paying: %v", err)
	}

	return &npool.CreatePaymentResponse{
		Info: dbRowToPayment(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetPaymentRequest) (*npool.GetPaymentResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	infos, err := db.Client().
		Payment.
		Query().
		Where(
			payment.And(
				payment.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query good paying: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty good paying")
	}

	return &npool.GetPaymentResponse{
		Info: dbRowToPayment(infos[0]),
	}, nil
}

package compensate

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-order/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/compensate"

	"github.com/google/uuid" //nolint

	"golang.org/x/xerrors" //nolint
)

func validateCompensate(info *npool.Compensate) error {
	if _, err := uuid.Parse(info.GetOrderID()); err != nil {
		return xerrors.Errorf("invalid order id: %v", err)
	}
	if _, err := uuid.Parse(info.GetPaymentID()); err != nil {
		return xerrors.Errorf("invalid payment id: %v", err)
	}
	return nil
}

func dbRowToCompensate(row *ent.Compensate) *npool.Compensate {
	return &npool.Compensate{
		ID:        row.ID.String(),
		OrderID:   row.OrderID.String(),
		PaymentID: row.PaymentID.String(),
	}
}

func Create(ctx context.Context, in *npool.CreateCompensateRequest) (*npool.CreateCompensateResponse, error) {
	if err := validateCompensate(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	info, err := db.Client().
		Compensate.
		Create().
		SetOrderID(uuid.MustParse(in.GetInfo().GetOrderID())).
		SetPaymentID(uuid.MustParse(in.GetInfo().GetPaymentID())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create compensate: %v", err)
	}

	return &npool.CreateCompensateResponse{
		Info: dbRowToCompensate(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetCompensateRequest) (*npool.GetCompensateResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	infos, err := db.Client().
		Compensate.
		Query().
		Where(
			compensate.And(
				compensate.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query compensate: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty compensate")
	}

	return &npool.GetCompensateResponse{
		Info: dbRowToCompensate(infos[0]),
	}, nil
}

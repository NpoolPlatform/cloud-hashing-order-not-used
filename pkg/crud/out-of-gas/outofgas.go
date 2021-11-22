package outofgas

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-order/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/outofgas"

	"github.com/google/uuid" //nolint

	"golang.org/x/xerrors" //nolint
)

func validateOutOfGas(info *npool.OutOfGas) error {
	if _, err := uuid.Parse(info.GetOrderID()); err != nil {
		return xerrors.Errorf("invalid order id: %v", err)
	}
	if _, err := uuid.Parse(info.GetPaymentID()); err != nil {
		return xerrors.Errorf("invalid payment id: %v", err)
	}
	return nil
}

func dbRowToOutOfGas(row *ent.OutOfGas) *npool.OutOfGas {
	return &npool.OutOfGas{
		ID:        row.ID.String(),
		OrderID:   row.OrderID.String(),
		PaymentID: row.PaymentID.String(),
	}
}

func Create(ctx context.Context, in *npool.CreateOutOfGasRequest) (*npool.CreateOutOfGasResponse, error) {
	if err := validateOutOfGas(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	info, err := db.Client().
		OutOfGas.
		Create().
		SetOrderID(uuid.MustParse(in.GetInfo().GetOrderID())).
		SetPaymentID(uuid.MustParse(in.GetInfo().GetPaymentID())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create out of gas: %v", err)
	}

	return &npool.CreateOutOfGasResponse{
		Info: dbRowToOutOfGas(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetOutOfGasRequest) (*npool.GetOutOfGasResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	infos, err := db.Client().
		OutOfGas.
		Query().
		Where(
			outofgas.And(
				outofgas.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query out of gas: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty out of gas")
	}

	return &npool.GetOutOfGasResponse{
		Info: dbRowToOutOfGas(infos[0]),
	}, nil
}

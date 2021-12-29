package gaspaying

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-order/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/gaspaying"

	"github.com/google/uuid" //nolint

	"golang.org/x/xerrors" //nolint
)

func validateGasPaying(info *npool.GasPaying) error {
	if _, err := uuid.Parse(info.GetOrderID()); err != nil {
		return xerrors.Errorf("invalid order id: %v", err)
	}
	if _, err := uuid.Parse(info.GetPaymentID()); err != nil {
		return xerrors.Errorf("invalid payment id: %v", err)
	}
	if _, err := uuid.Parse(info.GetFeeTypeID()); err != nil {
		return xerrors.Errorf("invalid fee type id: %v", err)
	}
	return nil
}

func dbRowToGasPaying(row *ent.GasPaying) *npool.GasPaying {
	return &npool.GasPaying{
		ID:              row.ID.String(),
		OrderID:         row.OrderID.String(),
		DurationMinutes: row.DurationMinutes,
		PaymentID:       row.PaymentID.String(),
		FeeTypeID:       row.FeeTypeID.String(),
	}
}

func Create(ctx context.Context, in *npool.CreateGasPayingRequest) (*npool.CreateGasPayingResponse, error) {
	if err := validateGasPaying(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		GasPaying.
		Create().
		SetOrderID(uuid.MustParse(in.GetInfo().GetOrderID())).
		SetPaymentID(uuid.MustParse(in.GetInfo().GetPaymentID())).
		SetFeeTypeID(uuid.MustParse(in.GetInfo().GetFeeTypeID())).
		SetDurationMinutes(in.GetInfo().GetDurationMinutes()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create gas paying: %v", err)
	}

	return &npool.CreateGasPayingResponse{
		Info: dbRowToGasPaying(info),
	}, nil
}

func GetByOrder(ctx context.Context, in *npool.GetGasPayingsByOrderRequest) (*npool.GetGasPayingsByOrderResponse, error) {
	orderID, err := uuid.Parse(in.GetOrderID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		GasPaying.
		Query().
		Where(
			gaspaying.And(
				gaspaying.OrderID(orderID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query gas paying: %v", err)
	}

	payings := []*npool.GasPaying{}
	for _, info := range infos {
		payings = append(payings, dbRowToGasPaying(info))
	}

	return &npool.GetGasPayingsByOrderResponse{
		Infos: payings,
	}, nil
}

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
	if _, err := uuid.Parse(info.GetAccountID()); err != nil {
		return xerrors.Errorf("invalid account id: %v", err)
	}
	return nil
}

func dbRowToGasPaying(row *ent.GasPaying) *npool.GasPaying {
	return &npool.GasPaying{
		ID:                    row.ID.String(),
		OrderID:               row.OrderID.String(),
		AccountID:             row.AccountID.String(),
		State:                 string(row.State),
		ChainTransactionID:    row.ChainTransactionID,
		PlatformTransactionID: row.PlatformTransactionID.String(),
		DurationMinutes:       row.DurationMinutes,
		Used:                  row.Used,
	}
}

func Create(ctx context.Context, in *npool.CreateGasPayingRequest) (*npool.CreateGasPayingResponse, error) {
	if err := validateGasPaying(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	info, err := db.Client().
		GasPaying.
		Create().
		SetOrderID(uuid.MustParse(in.GetInfo().GetOrderID())).
		SetAccountID(uuid.MustParse(in.GetInfo().GetAccountID())).
		SetChainTransactionID("").
		SetPlatformTransactionID(uuid.UUID{}).
		SetState("wait").
		SetDurationMinutes(in.GetInfo().GetDurationMinutes()).
		SetUsed(false).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create gas paying: %v", err)
	}

	return &npool.CreateGasPayingResponse{
		Info: dbRowToGasPaying(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetGasPayingRequest) (*npool.GetGasPayingResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	infos, err := db.Client().
		GasPaying.
		Query().
		Where(
			gaspaying.And(
				gaspaying.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query gas paying: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty gas paying")
	}

	return &npool.GetGasPayingResponse{
		Info: dbRowToGasPaying(infos[0]),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateGasPayingRequest) (*npool.UpdateGasPayingResponse, error) {
	if err := validateGasPaying(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	platformTID, err := uuid.Parse(in.GetInfo().GetPlatformTransactionID())
	if err != nil {
		return nil, xerrors.Errorf("invalid platform transaction id: %v", err)
	}

	info, err := db.Client().
		GasPaying.
		UpdateOneID(id).
		SetState(gaspaying.State(in.GetInfo().GetState())).
		SetChainTransactionID(in.GetInfo().GetChainTransactionID()).
		SetPlatformTransactionID(platformTID).
		SetUsed(in.GetInfo().GetUsed()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update gas paying: %v", err)
	}

	return &npool.UpdateGasPayingResponse{
		Info: dbRowToGasPaying(info),
	}, nil
}

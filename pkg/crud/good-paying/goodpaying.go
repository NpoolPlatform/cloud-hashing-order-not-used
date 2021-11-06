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
	if _, err := uuid.Parse(info.GetAccountID()); err != nil {
		return xerrors.Errorf("invalid account id: %v", err)
	}
	return nil
}

func dbRowToGoodPaying(row *ent.GoodPaying) *npool.GoodPaying {
	return &npool.GoodPaying{
		ID:                    row.ID.String(),
		OrderID:               row.OrderID.String(),
		AccountID:             row.AccountID.String(),
		State:                 string(row.State),
		ChainTransactionID:    row.ChainTransactionID,
		PlatformTransactionID: row.PlatformTransactionID.String(),
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
		SetAccountID(uuid.MustParse(in.GetInfo().GetAccountID())).
		SetChainTransactionID("").
		SetPlatformTransactionID(uuid.UUID{}).
		SetState("wait").
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create good paying: %v", err)
	}

	return &npool.CreateGoodPayingResponse{
		Info: dbRowToGoodPaying(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetGoodPayingRequest) (*npool.GetGoodPayingResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	infos, err := db.Client().
		GoodPaying.
		Query().
		Where(
			goodpaying.And(
				goodpaying.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query good paying: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty good paying")
	}

	return &npool.GetGoodPayingResponse{
		Info: dbRowToGoodPaying(infos[0]),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateGoodPayingRequest) (*npool.UpdateGoodPayingResponse, error) {
	if err := validateGoodPaying(in.GetInfo()); err != nil {
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
		GoodPaying.
		UpdateOneID(id).
		SetState(goodpaying.State(in.GetInfo().GetState())).
		SetChainTransactionID(in.GetInfo().GetChainTransactionID()).
		SetPlatformTransactionID(platformTID).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update good paying: %v", err)
	}

	return &npool.UpdateGoodPayingResponse{
		Info: dbRowToGoodPaying(info),
	}, nil
}

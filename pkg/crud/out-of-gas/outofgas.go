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
	return nil
}

func dbRowToOutOfGas(row *ent.OutOfGas) *npool.OutOfGas {
	return &npool.OutOfGas{
		ID:      row.ID.String(),
		OrderID: row.OrderID.String(),
		Start:   row.Start,
		End:     row.End,
	}
}

func Create(ctx context.Context, in *npool.CreateOutOfGasRequest) (*npool.CreateOutOfGasResponse, error) {
	if err := validateOutOfGas(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		OutOfGas.
		Create().
		SetOrderID(uuid.MustParse(in.GetInfo().GetOrderID())).
		SetStart(in.GetInfo().GetStart()).
		SetEnd(in.GetInfo().GetEnd()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create out of gas: %v", err)
	}

	return &npool.CreateOutOfGasResponse{
		Info: dbRowToOutOfGas(info),
	}, nil
}

func GetByOrder(ctx context.Context, in *npool.GetOutOfGasesByOrderRequest) (*npool.GetOutOfGasesByOrderResponse, error) {
	orderID, err := uuid.Parse(in.GetOrderID())
	if err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		OutOfGas.
		Query().
		Where(
			outofgas.And(
				outofgas.OrderID(orderID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query out of gas: %v", err)
	}

	outOfGases := []*npool.OutOfGas{}
	for _, info := range infos {
		outOfGases = append(outOfGases, dbRowToOutOfGas(info))
	}

	return &npool.GetOutOfGasesByOrderResponse{
		Infos: outOfGases,
	}, nil
}

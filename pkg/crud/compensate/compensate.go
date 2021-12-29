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
	return nil
}

func dbRowToCompensate(row *ent.Compensate) *npool.Compensate {
	return &npool.Compensate{
		ID:      row.ID.String(),
		OrderID: row.OrderID.String(),
		Start:   row.Start,
		End:     row.End,
		Message: row.Message,
	}
}

func Create(ctx context.Context, in *npool.CreateCompensateRequest) (*npool.CreateCompensateResponse, error) {
	if err := validateCompensate(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		Compensate.
		Create().
		SetOrderID(uuid.MustParse(in.GetInfo().GetOrderID())).
		SetStart(in.GetInfo().GetStart()).
		SetEnd(in.GetInfo().GetEnd()).
		SetMessage(in.GetInfo().GetMessage()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create compensate: %v", err)
	}

	return &npool.CreateCompensateResponse{
		Info: dbRowToCompensate(info),
	}, nil
}

func GetByOrder(ctx context.Context, in *npool.GetCompensatesByOrderRequest) (*npool.GetCompensatesByOrderResponse, error) {
	orderID, err := uuid.Parse(in.GetOrderID())
	if err != nil {
		return nil, xerrors.Errorf("invalid order id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		Compensate.
		Query().
		Where(
			compensate.And(
				compensate.OrderID(orderID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query compensate: %v", err)
	}

	compensates := []*npool.Compensate{}
	for _, info := range infos {
		compensates = append(compensates, dbRowToCompensate(info))
	}

	return &npool.GetCompensatesByOrderResponse{
		Infos: compensates,
	}, nil
}

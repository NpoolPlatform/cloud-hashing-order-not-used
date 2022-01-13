package payment

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-order"

	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/payment"

	"github.com/NpoolPlatform/go-service-framework/pkg/price"

	"github.com/google/uuid" //nolint

	"golang.org/x/xerrors" //nolint
)

func validatePayment(info *npool.Payment) error {
	if _, err := uuid.Parse(info.GetOrderID()); err != nil {
		return xerrors.Errorf("invalid order id: %v", err)
	}
	if _, err := uuid.Parse(info.GetAccountID()); err != nil {
		return xerrors.Errorf("invalid order id: %v", err)
	}
	if _, err := uuid.Parse(info.GetCoinInfoID()); err != nil {
		return xerrors.Errorf("invalid order id: %v", err)
	}
	return nil
}

func dbRowToPayment(row *ent.Payment) *npool.Payment {
	return &npool.Payment{
		ID:                    row.ID.String(),
		OrderID:               row.OrderID.String(),
		AccountID:             row.AccountID.String(),
		StartAmount:           price.DBPriceToVisualPrice(row.StartAmount),
		Amount:                price.DBPriceToVisualPrice(row.Amount),
		CoinInfoID:            row.CoinInfoID.String(),
		State:                 string(row.State),
		ChainTransactionID:    row.ChainTransactionID,
		PlatformTransactionID: row.PlatformTransactionID.String(),
	}
}

func Create(ctx context.Context, in *npool.CreatePaymentRequest) (*npool.CreatePaymentResponse, error) {
	if err := validatePayment(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	myPtid := uuid.UUID{}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		Payment.
		Create().
		SetOrderID(uuid.MustParse(in.GetInfo().GetOrderID())).
		SetAccountID(uuid.MustParse(in.GetInfo().GetAccountID())).
		SetStartAmount(price.VisualPriceToDBPrice(in.GetInfo().GetStartAmount())).
		SetAmount(price.VisualPriceToDBPrice(in.GetInfo().GetAmount())).
		SetCoinInfoID(uuid.MustParse(in.GetInfo().GetCoinInfoID())).
		SetState("wait").
		SetChainTransactionID("").
		SetPlatformTransactionID(myPtid).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create good paying: %v", err)
	}

	return &npool.CreatePaymentResponse{
		Info: dbRowToPayment(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdatePaymentRequest) (*npool.UpdatePaymentResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid payment id: %v", err)
	}

	ptid, err := uuid.Parse(in.GetInfo().GetPlatformTransactionID())
	if err != nil {
		return nil, xerrors.Errorf("invalid payment platform transaction id: %v", err)
	}

	if err := validatePayment(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		Payment.
		UpdateOneID(id).
		SetState(payment.State(in.GetInfo().GetState())).
		SetChainTransactionID(in.GetInfo().GetChainTransactionID()).
		SetPlatformTransactionID(ptid).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update payment: %v", err)
	}

	return &npool.UpdatePaymentResponse{
		Info: dbRowToPayment(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetPaymentRequest) (*npool.GetPaymentResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		Payment.
		Query().
		Where(
			payment.And(
				payment.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query payment: %v", err)
	}

	var pay *npool.Payment
	for _, info := range infos {
		pay = dbRowToPayment(info)
		break
	}

	return &npool.GetPaymentResponse{
		Info: pay,
	}, nil
}

func GetByOrder(ctx context.Context, in *npool.GetPaymentByOrderRequest) (*npool.GetPaymentByOrderResponse, error) {
	orderID, err := uuid.Parse(in.GetOrderID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		Payment.
		Query().
		Where(
			payment.And(
				payment.OrderID(orderID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query payment by order: %v", err)
	}

	var pay *npool.Payment
	for _, info := range infos {
		pay = dbRowToPayment(info)
		break
	}

	return &npool.GetPaymentByOrderResponse{
		Info: pay,
	}, nil
}

//---------------------------------------------------------------------------------------------------------------------------

func GetByState(ctx context.Context, state string) ([]*npool.Payment, error) {
	payState := payment.State(state)

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		Payment.
		Query().
		Where(
			payment.And(
				payment.StateEQ(payState),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query payment by state: %v", err)
	}

	payments := []*npool.Payment{}
	for _, info := range infos {
		payments = append(payments, dbRowToPayment(info))
	}

	return payments, nil
}

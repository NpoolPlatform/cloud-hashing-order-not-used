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
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		return xerrors.Errorf("invalid user id: %v", err)
	}
	if _, err := uuid.Parse(info.GetGoodID()); err != nil {
		return xerrors.Errorf("invalid good id: %v", err)
	}
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
		AppID:                 row.AppID.String(),
		UserID:                row.UserID.String(),
		GoodID:                row.GoodID.String(),
		OrderID:               row.OrderID.String(),
		AccountID:             row.AccountID.String(),
		StartAmount:           price.DBPriceToVisualPrice(row.StartAmount),
		FinishAmount:          price.DBPriceToVisualPrice(row.FinishAmount),
		Amount:                price.DBPriceToVisualPrice(row.Amount),
		CoinUSDCurrency:       price.DBPriceToVisualPrice(row.CoinUsdCurrency),
		CoinInfoID:            row.CoinInfoID.String(),
		State:                 string(row.State),
		ChainTransactionID:    row.ChainTransactionID,
		PlatformTransactionID: row.PlatformTransactionID.String(),
		UserSetPaid:           row.UserSetPaid,
		UserSetCanceled:       row.UserSetCanceled,
		UserPaymentTXID:       row.UserPaymentTxid,
		CreateAt:              row.CreateAt,
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
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetUserID(uuid.MustParse(in.GetInfo().GetUserID())).
		SetGoodID(uuid.MustParse(in.GetInfo().GetGoodID())).
		SetOrderID(uuid.MustParse(in.GetInfo().GetOrderID())).
		SetAccountID(uuid.MustParse(in.GetInfo().GetAccountID())).
		SetStartAmount(price.VisualPriceToDBPrice(in.GetInfo().GetStartAmount())).
		SetFinishAmount(price.VisualPriceToDBPrice(in.GetInfo().GetFinishAmount())).
		SetAmount(price.VisualPriceToDBPrice(in.GetInfo().GetAmount())).
		SetCoinUsdCurrency(price.VisualPriceToDBPrice(in.GetInfo().GetCoinUSDCurrency())).
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
		SetFinishAmount(price.VisualPriceToDBPrice(in.GetInfo().GetFinishAmount())).
		SetPlatformTransactionID(ptid).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update payment: %v", err)
	}

	return &npool.UpdatePaymentResponse{
		Info: dbRowToPayment(info),
	}, nil
}

func UpdateByUser(ctx context.Context, in *npool.UpdatePaymentByUserRequest) (*npool.UpdatePaymentByUserResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid payment id: %v", err)
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
		SetUserPaymentTxid(in.GetInfo().GetUserPaymentTXID()).
		SetUserSetPaid(in.GetInfo().GetUserSetPaid()).
		SetUserSetCanceled(in.GetInfo().GetUserSetCanceled()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update payment: %v", err)
	}

	return &npool.UpdatePaymentByUserResponse{
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
			payment.ID(id),
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
			payment.OrderID(orderID),
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

func GetByState(ctx context.Context, in *npool.GetPaymentsByStateRequest) (*npool.GetPaymentsByStateResponse, error) {
	payState := payment.State(in.GetState())

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		Payment.
		Query().
		Where(
			payment.StateEQ(payState),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query payment by state: %v", err)
	}

	payments := []*npool.Payment{}
	for _, info := range infos {
		payments = append(payments, dbRowToPayment(info))
	}

	return &npool.GetPaymentsByStateResponse{
		Infos: payments,
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetPaymentsByAppRequest) (*npool.GetPaymentsByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		Payment.
		Query().
		Where(
			payment.AppID(appID),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query payment by state: %v", err)
	}

	payments := []*npool.Payment{}
	for _, info := range infos {
		payments = append(payments, dbRowToPayment(info))
	}

	return &npool.GetPaymentsByAppResponse{
		Infos: payments,
	}, nil
}

func GetByAppUser(ctx context.Context, in *npool.GetPaymentsByAppUserRequest) (*npool.GetPaymentsByAppUserResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	userID, err := uuid.Parse(in.GetUserID())
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
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
				payment.AppID(appID),
				payment.UserID(userID),
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

	return &npool.GetPaymentsByAppUserResponse{
		Infos: payments,
	}, nil
}

func GetByAppUserState(ctx context.Context, in *npool.GetPaymentsByAppUserStateRequest) (*npool.GetPaymentsByAppUserStateResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	userID, err := uuid.Parse(in.GetUserID())
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
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
				payment.AppID(appID),
				payment.UserID(userID),
				payment.StateEQ(payment.State(in.GetState())),
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

	return &npool.GetPaymentsByAppUserStateResponse{
		Infos: payments,
	}, nil
}

func GetAll(ctx context.Context, in *npool.GetPaymentsRequest) (*npool.GetPaymentsResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		Payment.
		Query().
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query payment by state: %v", err)
	}

	payments := []*npool.Payment{}
	for _, info := range infos {
		payments = append(payments, dbRowToPayment(info))
	}

	return &npool.GetPaymentsResponse{
		Infos: payments,
	}, nil
}

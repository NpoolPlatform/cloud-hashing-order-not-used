package paymentwatcher

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/cloud-hashing-order/message/npool"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/payment"
	grpc2 "github.com/NpoolPlatform/cloud-hashing-order/pkg/grpc"

	billingpb "github.com/NpoolPlatform/cloud-hashing-billing/message/npool"
	coininfopb "github.com/NpoolPlatform/message/npool/coininfo"
	sphinxproxypb "github.com/NpoolPlatform/message/npool/sphinxproxy"
)

func watchPaymentState(ctx context.Context) {
	payments, err := payment.GetByState(ctx, "wait")
	if err != nil {
		logger.Sugar().Errorf("fail to get wait payments: %v", err)
		return
	}

	for _, pay := range payments {
		coinInfo, err := grpc2.GetCoinInfo(ctx, &coininfopb.GetCoinInfoRequest{
			ID: pay.CoinInfoID,
		})
		if err != nil {
			logger.Sugar().Errorf("fail to get coin %v info: %v", pay.CoinInfoID, err)
			continue
		}

		account, err := grpc2.GetBillingAccount(ctx, &billingpb.GetCoinAccountRequest{
			ID: pay.AccountID,
		})
		if err != nil {
			logger.Sugar().Errorf("fail to get payment account: %v", err)
			continue
		}

		balance, err := grpc2.GetBalance(ctx, &sphinxproxypb.GetBalanceRequest{
			Name:    coinInfo.Info.Name,
			Address: account.Info.Address,
		})
		if err != nil {
			logger.Sugar().Errorf("fail to get wallet balance: %v", err)
			continue
		}

		logger.Sugar().Infof("payment %v checking coin %v balance %v start amount %v pay amount %v",
			pay.ID, coinInfo.Info.Name, balance.Info.Balance, pay.StartAmount, pay.Amount)

		if balance.Info.Balance-pay.StartAmount > pay.Amount {
			pay.State = "done"
			_, err := payment.Update(ctx, &npool.UpdatePaymentRequest{
				Info: pay,
			})
			if err != nil {
				logger.Sugar().Errorf("fail to update payment state: %v", err)
				continue
			}
			logger.Sugar().Infof("payment %v done", pay.ID)
		}

		// TODO: payment timeout
	}
}

func Watch(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)

	for { //nolint
		select {
		case <-ticker.C:
			watchPaymentState(ctx)
		}
	}
}

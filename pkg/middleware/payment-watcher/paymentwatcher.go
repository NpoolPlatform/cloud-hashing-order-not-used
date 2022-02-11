package paymentwatcher

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"

	constant "github.com/NpoolPlatform/cloud-hashing-order/pkg/const"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/payment"
	grpc2 "github.com/NpoolPlatform/cloud-hashing-order/pkg/grpc"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-order"

	billingpb "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"
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

		newState := pay.State
		if balance.Info.Balance-pay.StartAmount > pay.Amount {
			newState = constant.PaymentStateDone
		}
		if pay.CreateAt+constant.TimeoutSeconds < uint32(time.Now().Unix()) {
			newState = constant.PaymentStateTimeout
		}

		if newState != pay.State {
			logger.Sugar().Infof("payment %v try %v -> %v", pay.ID, pay.State, newState)
			pay.State = newState
			_, err := payment.Update(ctx, &npool.UpdatePaymentRequest{
				Info: pay,
			})
			if err != nil {
				logger.Sugar().Errorf("fail to update payment state: %v", err)
				continue
			}

			lockKey := AccountLockKey(account.Info.ID)
			err = redis2.Unlock(lockKey)
			if err != nil {
				logger.Sugar().Errorf("fail unlock %v: %v", lockKey, err)
			}
		}
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

func AccountLockKey(accountID string) string {
	return fmt.Sprintf("%v:%v", constant.OrderPaymentLockKeyPrefix, accountID)
}

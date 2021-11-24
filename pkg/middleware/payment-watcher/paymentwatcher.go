package paymentwatcher

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/payment"
)

func watchPaymentState(ctx context.Context) {
	payments, err := payment.GetByState(ctx, "wait")
	if err != nil {
		logger.Sugar().Errorf("fail to get wait payments: %v", err)
		return
	}

	for _, pay := range payments {
		logger.Sugar().Infof("%v", pay)
	}
}

func Watch(ctx context.Context) {
	ticker := time.NewTicker(time.Second)

	for { //nolint
		select {
		case <-ticker.C:
			watchPaymentState(ctx)
		}
	}
}

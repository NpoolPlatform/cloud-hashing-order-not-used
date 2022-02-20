package payment

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/cloud-hashing-order/pkg/test-init" //nolint
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-order"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func assertPayment(t *testing.T, actual, expected *npool.Payment) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.UserID, expected.UserID)
	assert.Equal(t, actual.OrderID, expected.OrderID)
	assert.Equal(t, actual.AccountID, expected.AccountID)
	assert.Equal(t, actual.State, expected.State)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	payment := npool.Payment{
		AppID:      uuid.New().String(),
		UserID:     uuid.New().String(),
		OrderID:    uuid.New().String(),
		AccountID:  uuid.New().String(),
		CoinInfoID: uuid.New().String(),
	}
	resp, err := Create(context.Background(), &npool.CreatePaymentRequest{
		Info: &payment,
	})
	payment.State = "wait"
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assert.Equal(t, resp.Info.ChainTransactionID, "")
		assert.Equal(t, resp.Info.PlatformTransactionID, uuid.UUID{}.String())
		assertPayment(t, resp.Info, &payment)
	}

	payment.State = "done"
	payment.ChainTransactionID = "MOCKTRANSACTIONID"
	payment.PlatformTransactionID = uuid.New().String()
	payment.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdatePaymentRequest{
		Info: &payment,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assert.Equal(t, resp1.Info.ChainTransactionID, payment.ChainTransactionID)
		assert.Equal(t, resp1.Info.PlatformTransactionID, payment.PlatformTransactionID)
		assertPayment(t, resp1.Info, &payment)
	}

	resp2, err := GetByOrder(context.Background(), &npool.GetPaymentByOrderRequest{
		OrderID: payment.OrderID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assert.Equal(t, resp2.Info.ChainTransactionID, payment.ChainTransactionID)
		assert.Equal(t, resp2.Info.PlatformTransactionID, payment.PlatformTransactionID)
		assertPayment(t, resp2.Info, &payment)
	}

	resp3, err := Get(context.Background(), &npool.GetPaymentRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assert.Equal(t, resp3.Info.ChainTransactionID, payment.ChainTransactionID)
		assert.Equal(t, resp3.Info.PlatformTransactionID, payment.PlatformTransactionID)
		assertPayment(t, resp3.Info, &payment)
	}
}

package gaspaying

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/cloud-hashing-order/message/npool"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/test-init" //nolint

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

func assertGasPaying(t *testing.T, actual, expected *npool.GasPaying) {
	assert.Equal(t, actual.OrderID, expected.OrderID)
	assert.Equal(t, actual.PaymentID, expected.PaymentID)
	assert.Equal(t, actual.FeeTypeID, expected.FeeTypeID)
	assert.Equal(t, actual.DurationMinutes, expected.DurationMinutes)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gasPaying := npool.GasPaying{
		OrderID:         uuid.New().String(),
		PaymentID:       uuid.New().String(),
		FeeTypeID:       uuid.New().String(),
		DurationMinutes: 24 * 10 * 60,
	}
	resp, err := Create(context.Background(), &npool.CreateGasPayingRequest{
		Info: &gasPaying,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertGasPaying(t, resp.Info, &gasPaying)
	}

	gasPaying.ID = resp.Info.ID

	resp1, err := GetByOrder(context.Background(), &npool.GetGasPayingsByOrderRequest{
		OrderID: gasPaying.OrderID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp1.Infos), 1)
	}
}

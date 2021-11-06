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
	assert.Equal(t, actual.AccountID, expected.AccountID)
	assert.Equal(t, actual.State, expected.State)
	assert.Equal(t, actual.DurationMinutes, expected.DurationMinutes)
	assert.Equal(t, actual.Used, expected.Used)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gasPaying := npool.GasPaying{
		OrderID:         uuid.New().String(),
		AccountID:       uuid.New().String(),
		DurationMinutes: 24 * 10 * 60,
	}
	resp, err := Create(context.Background(), &npool.CreateGasPayingRequest{
		Info: &gasPaying,
	})
	gasPaying.State = "wait"
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assert.Equal(t, resp.Info.ChainTransactionID, "")
		assert.Equal(t, resp.Info.PlatformTransactionID, uuid.UUID{}.String())
		assertGasPaying(t, resp.Info, &gasPaying)
	}

	gasPaying.State = "done"
	gasPaying.Used = true
	gasPaying.ChainTransactionID = "MOCKTRANSACTIONID"
	gasPaying.PlatformTransactionID = uuid.New().String()
	gasPaying.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateGasPayingRequest{
		Info: &gasPaying,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assert.Equal(t, resp1.Info.ChainTransactionID, gasPaying.ChainTransactionID)
		assert.Equal(t, resp1.Info.PlatformTransactionID, gasPaying.PlatformTransactionID)
		assertGasPaying(t, resp1.Info, &gasPaying)
	}

	resp2, err := Get(context.Background(), &npool.GetGasPayingRequest{
		ID: resp1.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assert.Equal(t, resp2.Info.ChainTransactionID, gasPaying.ChainTransactionID)
		assert.Equal(t, resp2.Info.PlatformTransactionID, gasPaying.PlatformTransactionID)
		assertGasPaying(t, resp2.Info, &gasPaying)
	}
}

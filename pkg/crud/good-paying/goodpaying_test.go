package goodpaying

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

func assertGoodPaying(t *testing.T, actual, expected *npool.GoodPaying) {
	assert.Equal(t, actual.OrderID, expected.OrderID)
	assert.Equal(t, actual.AccountID, expected.AccountID)
	assert.Equal(t, actual.State, expected.State)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	goodPaying := npool.GoodPaying{
		OrderID:   uuid.New().String(),
		AccountID: uuid.New().String(),
	}
	resp, err := Create(context.Background(), &npool.CreateGoodPayingRequest{
		Info: &goodPaying,
	})
	goodPaying.State = "wait"
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assert.Equal(t, resp.Info.ChainTransactionID, "")
		assert.Equal(t, resp.Info.PlatformTransactionID, uuid.UUID{}.String())
		assertGoodPaying(t, resp.Info, &goodPaying)
	}

	goodPaying.State = "done"
	goodPaying.ChainTransactionID = "MOCKTRANSACTIONID"
	goodPaying.PlatformTransactionID = uuid.New().String()
	goodPaying.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateGoodPayingRequest{
		Info: &goodPaying,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assert.Equal(t, resp1.Info.ChainTransactionID, goodPaying.ChainTransactionID)
		assert.Equal(t, resp1.Info.PlatformTransactionID, goodPaying.PlatformTransactionID)
		assertGoodPaying(t, resp1.Info, &goodPaying)
	}

	resp2, err := Get(context.Background(), &npool.GetGoodPayingRequest{
		ID: resp1.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assert.Equal(t, resp2.Info.ChainTransactionID, goodPaying.ChainTransactionID)
		assert.Equal(t, resp2.Info.PlatformTransactionID, goodPaying.PlatformTransactionID)
		assertGoodPaying(t, resp2.Info, &goodPaying)
	}
}

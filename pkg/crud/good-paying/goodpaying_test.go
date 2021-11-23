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
	assert.Equal(t, actual.PaymentID, expected.PaymentID)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	goodPaying := npool.GoodPaying{
		OrderID:   uuid.New().String(),
		PaymentID: uuid.New().String(),
	}
	resp, err := Create(context.Background(), &npool.CreateGoodPayingRequest{
		Info: &goodPaying,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertGoodPaying(t, resp.Info, &goodPaying)
	}

	goodPaying.ID = resp.Info.ID

	resp1, err := GetByOrder(context.Background(), &npool.GetGoodPayingByOrderRequest{
		OrderID: goodPaying.OrderID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertGoodPaying(t, resp1.Info, &goodPaying)
	}
}

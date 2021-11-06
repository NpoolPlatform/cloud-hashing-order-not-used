package order

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

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

func assertOrder(t *testing.T, actual, expected *npool.Order) {
	assert.Equal(t, actual.GoodID, expected.GoodID)
	assert.Equal(t, actual.Units, expected.Units)
	assert.Equal(t, actual.Discount, expected.Discount)
	assert.Equal(t, actual.SpecialReductionAmount, expected.SpecialReductionAmount)
	assert.Equal(t, actual.UserID, expected.UserID)
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.State, expected.State)
	assert.Equal(t, actual.GoodPayID, expected.GoodPayID)
	assert.Equal(t, actual.Start, expected.Start)
	assert.Equal(t, actual.End, expected.End)
	assert.Equal(t, actual.CompensateMinutes, expected.CompensateMinutes)
	assert.Equal(t, actual.CompensateElapsedMinutes, expected.CompensateElapsedMinutes)
	assert.Equal(t, actual.GasStart, expected.GasStart)
	assert.Equal(t, actual.GasEnd, expected.GasEnd)
	assert.Equal(t, actual.GasPayIDs, expected.GasPayIDs)
	assert.Equal(t, actual.CouponID, expected.CouponID)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	second := uint32(time.Now().Unix())

	order := npool.Order{
		GoodID:                   uuid.New().String(),
		Units:                    10,
		Discount:                 10,
		SpecialReductionAmount:   20,
		UserID:                   uuid.New().String(),
		AppID:                    uuid.New().String(),
		GoodPayID:                uuid.New().String(),
		Start:                    second,
		End:                      second + 20,
		CompensateMinutes:        1000,
		CompensateElapsedMinutes: 100,
		GasStart:                 second,
		GasEnd:                   second + 10,
		GasPayIDs:                []string{uuid.New().String(), uuid.New().String()},
		CouponID:                 uuid.New().String(),
	}

	order.State = "created"

	resp, err := Create(context.Background(), &npool.CreateOrderRequest{
		Info: &order,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assertOrder(t, resp.Info, &order)
	}
}

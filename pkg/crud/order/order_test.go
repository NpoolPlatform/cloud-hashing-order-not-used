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
		GoodPayID:                uuid.UUID{}.String(),
		Start:                    second,
		End:                      second + 20,
		CompensateMinutes:        1000,
		CompensateElapsedMinutes: 100,
		GasStart:                 second,
		GasEnd:                   second + 10,
		GasPayIDs:                []string{},
		CouponID:                 uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateOrderRequest{
		Info: &order,
	})

	order.State = "created"
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assertOrder(t, resp.Info, &order)
	}

	order.State = "paying"
	order.ID = resp.Info.ID
	order.GoodPayID = uuid.New().String()
	order.GasPayIDs = []string{uuid.New().String(), uuid.New().String()}

	resp1, err := Update(context.Background(), &npool.UpdateOrderRequest{
		Info: &order,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assert.Equal(t, resp1.Info.GoodPayID, order.GoodPayID)
		assert.Equal(t, resp1.Info.GasPayIDs, order.GasPayIDs)
		assertOrder(t, resp1.Info, &order)
	}

	resp2, err := Get(context.Background(), &npool.GetOrderRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assert.Equal(t, resp2.Info.GoodPayID, order.GoodPayID)
		assert.Equal(t, resp2.Info.GasPayIDs, order.GasPayIDs)
		assertOrder(t, resp2.Info, &order)
	}

	resp3, err := GetByAppUser(context.Background(), &npool.GetOrdersByAppUserRequest{
		AppID:  resp.Info.AppID,
		UserID: resp.Info.UserID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp3.Infos), 1)
	}

	resp4, err := GetByApp(context.Background(), &npool.GetOrdersByAppRequest{
		AppID: resp.Info.AppID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp4.Infos), 1)
	}

	resp5, err := GetByGood(context.Background(), &npool.GetOrdersByGoodRequest{
		GoodID: resp.Info.GoodID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp5.Infos), 1)
	}
}

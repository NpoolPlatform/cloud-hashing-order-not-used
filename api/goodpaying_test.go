package api

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/cloud-hashing-order/message/npool"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func assertGoodPaying(t *testing.T, actual, expected *npool.GoodPaying) {
	assert.Equal(t, actual.OrderID, expected.OrderID)
	assert.Equal(t, actual.PaymentID, expected.PaymentID)
}

func TestGoodPayingCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	goodPaying := npool.GoodPaying{
		OrderID:   uuid.New().String(),
		PaymentID: uuid.New().String(),
	}
	firstCreateInfo := npool.CreateGoodPayingResponse{}

	cli := resty.New()

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateGoodPayingRequest{
			Info: &goodPaying,
		}).
		Post("http://localhost:50040/v1/create/good/paying")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		err := json.Unmarshal(resp.Body(), &firstCreateInfo)
		if assert.Nil(t, err) {
			assert.NotEqual(t, firstCreateInfo.Info.ID, uuid.UUID{}.String())
			assertGoodPaying(t, firstCreateInfo.Info, &goodPaying)
		}
	}

	goodPaying.ID = firstCreateInfo.Info.ID

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetGoodPayingByOrderRequest{
			OrderID: goodPaying.OrderID,
		}).
		Post("http://localhost:50040/v1/get/good/paying/by/order")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.GetGoodPayingByOrderResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.ID, firstCreateInfo.Info.ID)
			assertGoodPaying(t, info.Info, &goodPaying)
		}
	}
}

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
	firstCreateInfo := npool.CreateGoodPayingResponse{}

	cli := resty.New()

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateGoodPayingRequest{
			Info: &goodPaying,
		}).
		Post("http://localhost:36759/v1/create/good/paying")
	goodPaying.State = "wait"
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		err := json.Unmarshal(resp.Body(), &firstCreateInfo)
		if assert.Nil(t, err) {
			assert.NotEqual(t, firstCreateInfo.Info.ID, uuid.UUID{}.String())
			assert.Equal(t, firstCreateInfo.Info.ChainTransactionID, "")
			assert.Equal(t, firstCreateInfo.Info.PlatformTransactionID, uuid.UUID{}.String())
			assertGoodPaying(t, firstCreateInfo.Info, &goodPaying)
		}
	}

	goodPaying.State = "done"
	goodPaying.ChainTransactionID = "MOCKTRANSACTIONID"
	goodPaying.PlatformTransactionID = uuid.New().String()
	goodPaying.ID = firstCreateInfo.Info.ID

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UpdateGoodPayingRequest{
			Info: &goodPaying,
		}).
		Post("http://localhost:36759/v1/update/good/paying")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.UpdateGoodPayingResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.ID, firstCreateInfo.Info.ID)
			assert.Equal(t, info.Info.ChainTransactionID, goodPaying.ChainTransactionID)
			assert.Equal(t, info.Info.PlatformTransactionID, goodPaying.PlatformTransactionID)
			assertGoodPaying(t, info.Info, &goodPaying)
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetGoodPayingRequest{
			ID: goodPaying.ID,
		}).
		Post("http://localhost:36759/v1/get/good/paying")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.GetGoodPayingResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.ID, firstCreateInfo.Info.ID)
			assert.Equal(t, info.Info.ChainTransactionID, goodPaying.ChainTransactionID)
			assert.Equal(t, info.Info.PlatformTransactionID, goodPaying.PlatformTransactionID)
			assertGoodPaying(t, info.Info, &goodPaying)
		}
	}
}

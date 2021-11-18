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

func assertGasPaying(t *testing.T, actual, expected *npool.GasPaying) {
	assert.Equal(t, actual.OrderID, expected.OrderID)
	assert.Equal(t, actual.AccountID, expected.AccountID)
	assert.Equal(t, actual.State, expected.State)
	assert.Equal(t, actual.DurationMinutes, expected.DurationMinutes)
	assert.Equal(t, actual.Used, expected.Used)
}

func TestGasPayingCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gasPaying := npool.GasPaying{
		OrderID:         uuid.New().String(),
		AccountID:       uuid.New().String(),
		DurationMinutes: 24 * 10 * 60,
	}
	firstCreateInfo := npool.CreateGasPayingResponse{}

	cli := resty.New()

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateGasPayingRequest{
			Info: &gasPaying,
		}).
		Post("http://localhost:50040/v1/create/gas/paying")
	gasPaying.State = "wait"
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		err := json.Unmarshal(resp.Body(), &firstCreateInfo)
		if assert.Nil(t, err) {
			assert.NotEqual(t, firstCreateInfo.Info.ID, uuid.UUID{}.String())
			assert.Equal(t, firstCreateInfo.Info.ChainTransactionID, "")
			assert.Equal(t, firstCreateInfo.Info.PlatformTransactionID, uuid.UUID{}.String())
			assertGasPaying(t, firstCreateInfo.Info, &gasPaying)
		}
	}

	gasPaying.State = "done"
	gasPaying.Used = true
	gasPaying.ChainTransactionID = "MOCKTRANSACTIONID"
	gasPaying.PlatformTransactionID = uuid.New().String()
	gasPaying.ID = firstCreateInfo.Info.ID

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UpdateGasPayingRequest{
			Info: &gasPaying,
		}).
		Post("http://localhost:50040/v1/update/gas/paying")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.UpdateGasPayingResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.ID, firstCreateInfo.Info.ID)
			assert.Equal(t, info.Info.ChainTransactionID, gasPaying.ChainTransactionID)
			assert.Equal(t, info.Info.PlatformTransactionID, gasPaying.PlatformTransactionID)
			assertGasPaying(t, info.Info, &gasPaying)
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetGasPayingRequest{
			ID: gasPaying.ID,
		}).
		Post("http://localhost:50040/v1/get/gas/paying")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.GetGasPayingResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.ID, firstCreateInfo.Info.ID)
			assert.Equal(t, info.Info.ChainTransactionID, gasPaying.ChainTransactionID)
			assert.Equal(t, info.Info.PlatformTransactionID, gasPaying.PlatformTransactionID)
			assertGasPaying(t, info.Info, &gasPaying)
		}
	}
}

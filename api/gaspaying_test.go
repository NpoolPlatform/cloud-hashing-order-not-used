package api

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-order"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func assertGasPaying(t *testing.T, actual, expected *npool.GasPaying) {
	assert.Equal(t, actual.OrderID, expected.OrderID)
	assert.Equal(t, actual.PaymentID, expected.PaymentID)
	assert.Equal(t, actual.FeeTypeID, expected.FeeTypeID)
	assert.Equal(t, actual.DurationMinutes, expected.DurationMinutes)
}

func TestGasPayingCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gasPaying := npool.GasPaying{
		OrderID:         uuid.New().String(),
		PaymentID:       uuid.New().String(),
		FeeTypeID:       uuid.New().String(),
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
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		err := json.Unmarshal(resp.Body(), &firstCreateInfo)
		if assert.Nil(t, err) {
			assert.NotEqual(t, firstCreateInfo.Info.ID, uuid.UUID{}.String())
			assertGasPaying(t, firstCreateInfo.Info, &gasPaying)
		}
	}

	gasPaying.ID = firstCreateInfo.Info.ID

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetGasPayingsByOrderRequest{
			OrderID: gasPaying.OrderID,
		}).
		Post("http://localhost:50040/v1/get/gas/payings/by/order")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.GetGasPayingsByOrderResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, len(info.Infos), 1)
		}
	}
}

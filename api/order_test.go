package api

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/cloud-hashing-order/message/npool"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestOrderCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	second := uint32(time.Now().Unix())
	appID := uuid.New().String()
	userID := uuid.New().String()
	goodID := uuid.New().String()

	myOrder := npool.Order{
		GoodID:                   goodID,
		Units:                    10,
		Discount:                 10,
		SpecialReductionAmount:   20,
		UserID:                   userID,
		AppID:                    appID,
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
	orderResp := npool.CreateOrderResponse{}

	cli := resty.New()

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateOrderRequest{
			Info: &myOrder,
		}).
		Post("http://localhost:36759/v1/create/order")
	assert.Nil(t, err)

	err = json.Unmarshal(resp.Body(), &orderResp)
	assert.Nil(t, err)

	goodPaying := npool.GoodPaying{
		OrderID:   orderResp.Info.ID,
		AccountID: uuid.New().String(),
	}
	_, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateGoodPayingRequest{
			Info: &goodPaying,
		}).
		Post("http://localhost:36759/v1/create/good/paying")
	assert.Nil(t, err)

	gasPaying1 := npool.GasPaying{
		OrderID:         orderResp.Info.ID,
		AccountID:       uuid.New().String(),
		DurationMinutes: 24 * 10 * 60,
	}
	_, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateGasPayingRequest{
			Info: &gasPaying1,
		}).
		Post("http://localhost:36759/v1/create/gas/paying")
	assert.Nil(t, err)

	gasPaying2 := npool.GasPaying{
		OrderID:         orderResp.Info.ID,
		AccountID:       uuid.New().String(),
		DurationMinutes: 24 * 10 * 60,
	}
	_, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateGasPayingRequest{
			Info: &gasPaying2,
		}).
		Post("http://localhost:36759/v1/create/gas/paying")
	assert.Nil(t, err)
}

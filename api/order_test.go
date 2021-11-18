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
		Post("http://localhost:50040/v1/create/order")
	assert.Nil(t, err)

	err = json.Unmarshal(resp.Body(), &orderResp)
	assert.Nil(t, err)

	goodPaying := npool.GoodPaying{
		OrderID:   orderResp.Info.ID,
		AccountID: uuid.New().String(),
	}
	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateGoodPayingRequest{
			Info: &goodPaying,
		}).
		Post("http://localhost:50040/v1/create/good/paying")
	assert.Nil(t, err)

	goodPayingResp := npool.CreateGoodPayingResponse{}
	err = json.Unmarshal(resp.Body(), &goodPayingResp)
	assert.Nil(t, err)
	myOrder.GoodPayID = goodPayingResp.Info.ID

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
		Post("http://localhost:50040/v1/create/gas/paying")
	assert.Nil(t, err)

	gasPayingResp := npool.CreateGasPayingResponse{}
	err = json.Unmarshal(resp.Body(), &gasPayingResp)
	assert.Nil(t, err)
	myOrder.GasPayIDs = append(myOrder.GasPayIDs, gasPayingResp.Info.ID)

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
		Post("http://localhost:50040/v1/create/gas/paying")
	assert.Nil(t, err)

	err = json.Unmarshal(resp.Body(), &gasPayingResp)
	assert.Nil(t, err)
	myOrder.GasPayIDs = append(myOrder.GasPayIDs, gasPayingResp.Info.ID)

	_, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetOrderDetailRequest{
			ID: orderResp.Info.ID,
		}).
		Post("http://localhost:50040/v1/get/order/detail")
	assert.Nil(t, err)

	myOrder.ID = orderResp.Info.ID
	myOrder.State = "paying"

	_, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UpdateOrderRequest{
			Info: &myOrder,
		}).
		Post("http://localhost:50040/v1/update/order")
	assert.Nil(t, err)

	_, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetOrderRequest{
			ID: myOrder.ID,
		}).
		Post("http://localhost:50040/v1/get/order")
	assert.Nil(t, err)

	_, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetOrderDetailRequest{
			ID: myOrder.ID,
		}).
		Post("http://localhost:50040/v1/get/order/detail")
	assert.Nil(t, err)

	_, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetOrdersDetailByAppUserRequest{
			AppID:  myOrder.AppID,
			UserID: myOrder.UserID,
		}).
		Post("http://localhost:50040/v1/get/orders/detail/by/app/user")
	assert.Nil(t, err)

	_, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetOrdersDetailByGoodRequest{
			GoodID: myOrder.GoodID,
		}).
		Post("http://localhost:50040/v1/get/orders/detail/by/good")
	assert.Nil(t, err)
}

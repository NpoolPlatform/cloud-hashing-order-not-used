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
		AppID:                  appID,
		GoodID:                 goodID,
		UserID:                 userID,
		Units:                  10,
		DiscountCouponID:       uuid.New().String(),
		UserSpecialReductionID: uuid.New().String(),
		Start:                  second,
		End:                    second + 20,
		CouponID:               uuid.New().String(),
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

	_, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetOrderDetailRequest{
			ID: orderResp.Info.ID,
		}).
		Post("http://localhost:50040/v1/get/order/detail")
	assert.Nil(t, err)

	myOrder.ID = orderResp.Info.ID

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

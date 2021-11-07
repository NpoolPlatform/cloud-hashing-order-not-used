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

	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/gas-paying"  //nolint
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/good-paying" //nolint
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/order"       //nolint

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

func assertOrderDetail(t *testing.T, actual *npool.OrderDetail, orderInfo *npool.Order, goodPaying *npool.GoodPaying, _gasPayings []*npool.GasPaying) { //nolint
	assert.Equal(t, actual.GoodID, orderInfo.GoodID)
	assert.Equal(t, actual.Units, orderInfo.Units)
	assert.Equal(t, actual.Discount, orderInfo.Discount)
	assert.Equal(t, actual.SpecialReductionAmount, orderInfo.SpecialReductionAmount)
	assert.Equal(t, actual.UserID, orderInfo.UserID)
	assert.Equal(t, actual.AppID, orderInfo.AppID)

	assert.Equal(t, actual.GoodPaying.ID, goodPaying.ID)
	assert.Equal(t, actual.GoodPaying.OrderID, goodPaying.OrderID)
	assert.Equal(t, actual.GoodPaying.AccountID, goodPaying.AccountID)
	assert.Equal(t, actual.GoodPaying.State, goodPaying.State)
	assert.Equal(t, actual.GoodPaying.ChainTransactionID, goodPaying.ChainTransactionID)
	assert.Equal(t, actual.GoodPaying.PlatformTransactionID, goodPaying.PlatformTransactionID)

	assert.Equal(t, actual.Start, orderInfo.Start)
	assert.Equal(t, actual.End, orderInfo.End)
	assert.Equal(t, actual.CompensateMinutes, orderInfo.CompensateMinutes)
	assert.Equal(t, actual.CompensateElapsedMinutes, orderInfo.CompensateElapsedMinutes)
	assert.Equal(t, actual.GasStart, orderInfo.GasStart)
	assert.Equal(t, actual.GasEnd, orderInfo.GasEnd)
	assert.Equal(t, actual.CouponID, orderInfo.CouponID)
}

func TestGetDetail(t *testing.T) {
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
	orderResp, err := order.Create(context.Background(), &npool.CreateOrderRequest{
		Info: &myOrder,
	})
	assert.Nil(t, err)

	goodPaying := npool.GoodPaying{
		OrderID:   orderResp.Info.ID,
		AccountID: uuid.New().String(),
	}
	goodPayingResp, err := goodpaying.Create(context.Background(), &npool.CreateGoodPayingRequest{
		Info: &goodPaying,
	})
	assert.Nil(t, err)

	gasPaying1 := npool.GasPaying{
		OrderID:         orderResp.Info.ID,
		AccountID:       uuid.New().String(),
		DurationMinutes: 24 * 10 * 60,
	}
	gasPayingResp1, err := gaspaying.Create(context.Background(), &npool.CreateGasPayingRequest{
		Info: &gasPaying1,
	})
	assert.Nil(t, err)
	gasPaying1.ID = gasPayingResp1.Info.ID

	gasPaying2 := npool.GasPaying{
		OrderID:         orderResp.Info.ID,
		AccountID:       uuid.New().String(),
		DurationMinutes: 24 * 10 * 60,
	}
	gasPayingResp2, err := gaspaying.Create(context.Background(), &npool.CreateGasPayingRequest{
		Info: &gasPaying2,
	})
	assert.Nil(t, err)
	gasPaying2.ID = gasPayingResp2.Info.ID

	myOrder.GasPayIDs = []string{gasPayingResp1.Info.ID, gasPayingResp2.Info.ID}
	myOrder.GoodPayID = goodPayingResp.Info.ID
	myOrder.ID = orderResp.Info.ID
	myOrder.State = "paying"

	orderResp1, err := order.Update(context.Background(), &npool.UpdateOrderRequest{
		Info: &myOrder,
	})
	assert.Nil(t, err)

	orderDetail, err := Get(context.Background(), &npool.GetOrderDetailRequest{
		ID: orderResp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, orderDetail.Detail.ID, orderResp1.Info.ID)
		assertOrderDetail(t, orderDetail.Detail, &myOrder, goodPayingResp.Info,
			[]*npool.GasPaying{gasPayingResp1.Info, gasPayingResp2.Info})
	}

	orderDetails, err := GetByAppUser(context.Background(), &npool.GetOrdersDetailByAppUserRequest{
		AppID:  appID,
		UserID: userID,
	})
	if assert.Nil(t, err) {
		if assert.Equal(t, len(orderDetails.Details), 1) {
			assert.Equal(t, orderDetails.Details[0].ID, orderResp1.Info.ID)
			assertOrderDetail(t, orderDetails.Details[0], &myOrder, goodPayingResp.Info,
				[]*npool.GasPaying{gasPayingResp1.Info, gasPayingResp2.Info})
		}
	}

	orderDetails1, err := GetByApp(context.Background(), &npool.GetOrdersDetailByAppRequest{
		AppID: appID,
	})
	if assert.Nil(t, err) {
		if assert.Equal(t, len(orderDetails1.Details), 1) {
			assert.Equal(t, orderDetails1.Details[0].ID, orderResp1.Info.ID)
			assertOrderDetail(t, orderDetails1.Details[0], &myOrder, goodPayingResp.Info,
				[]*npool.GasPaying{gasPayingResp1.Info, gasPayingResp2.Info})
		}
	}

	orderDetails2, err := GetByGood(context.Background(), &npool.GetOrdersDetailByGoodRequest{
		GoodID: goodID,
	})
	if assert.Nil(t, err) {
		if assert.Equal(t, len(orderDetails2.Details), 1) {
			assert.Equal(t, orderDetails2.Details[0].ID, orderResp1.Info.ID)
			assertOrderDetail(t, orderDetails2.Details[0], &myOrder, goodPayingResp.Info,
				[]*npool.GasPaying{gasPayingResp1.Info, gasPayingResp2.Info})
		}
	}
}

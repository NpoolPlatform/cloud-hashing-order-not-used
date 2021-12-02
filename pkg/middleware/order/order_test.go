// +build !codeanalysis

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

	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/compensate"  //nolint
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/gas-paying"  //nolint
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/good-paying" //nolint
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/order"       //nolint
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/out-of-gas"  //nolint
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/payment"     //nolint

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

func assertOrderDetail( //nolint
	t *testing.T,
	actual *npool.OrderDetail,
	orderInfo *npool.Order,
	goodPaying *npool.GoodPaying,
	gasPayings []*npool.GasPaying, //nolint
	compensates []*npool.Compensate, //nolint
	outOfGases []*npool.OutOfGas, //nolint
	myPayment *npool.Payment) { //nolint
	assert.Equal(t, actual.GoodID, orderInfo.GoodID)
	assert.Equal(t, actual.AppID, orderInfo.AppID)
	assert.Equal(t, actual.UserID, orderInfo.UserID)
	assert.Equal(t, actual.Units, orderInfo.Units)
	assert.Equal(t, actual.DiscountCouponID, orderInfo.DiscountCouponID)
	assert.Equal(t, actual.UserSpecialReductionID, orderInfo.UserSpecialReductionID)

	assert.Equal(t, actual.GoodPaying, goodPaying)

	if goodPaying != nil && actual.GoodPaying != nil {
		assert.Equal(t, actual.GoodPaying.ID, goodPaying.ID)
		assert.Equal(t, actual.GoodPaying.OrderID, goodPaying.OrderID)
		assert.Equal(t, actual.GoodPaying.PaymentID, goodPaying.PaymentID)
	}

	// assert.EqualValues(t, actual.GasPayings, gasPayings) //nolint
	// assert.EqualValues(t, actual.Compensates, compensates) //nolint
	// assert.EqualValues(t, actual.OutOfGases, outOfGases) //nolint
	assert.Equal(t, actual.Payment, myPayment)

	assert.Equal(t, actual.Start, orderInfo.Start)
	assert.Equal(t, actual.End, orderInfo.End)
	assert.Equal(t, actual.CouponID, orderInfo.CouponID)
}

func TestGetDetail(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	second := uint32(time.Now().Unix())

	appID := uuid.New().String()
	userID := uuid.New().String()
	goodID := uuid.New().String()

	myOrder := npool.Order{
		GoodID:                 goodID,
		AppID:                  appID,
		UserID:                 userID,
		Units:                  10,
		DiscountCouponID:       uuid.New().String(),
		UserSpecialReductionID: uuid.New().String(),
		Start:                  second,
		End:                    second + 20,
		CouponID:               uuid.New().String(),
	}
	orderResp, err := order.Create(context.Background(), &npool.CreateOrderRequest{
		Info: &myOrder,
	})
	assert.Nil(t, err)

	myPayment := npool.Payment{
		OrderID:    orderResp.Info.ID,
		AccountID:  uuid.New().String(),
		Amount:     1490.6,
		CoinInfoID: uuid.New().String(),
	}
	paymentResp, err := payment.Create(context.Background(), &npool.CreatePaymentRequest{
		Info: &myPayment,
	})
	assert.Nil(t, err)

	goodPaying := npool.GoodPaying{
		OrderID:   orderResp.Info.ID,
		PaymentID: paymentResp.Info.ID,
	}
	goodPayingResp, err := goodpaying.Create(context.Background(), &npool.CreateGoodPayingRequest{
		Info: &goodPaying,
	})
	assert.Nil(t, err)

	gasPaying1 := npool.GasPaying{
		OrderID:         orderResp.Info.ID,
		PaymentID:       paymentResp.Info.ID,
		FeeTypeID:       uuid.New().String(),
		DurationMinutes: 24 * 10 * 60,
	}
	gasPayingResp1, err := gaspaying.Create(context.Background(), &npool.CreateGasPayingRequest{
		Info: &gasPaying1,
	})
	assert.Nil(t, err)
	gasPaying1.ID = gasPayingResp1.Info.ID

	gasPaying2 := npool.GasPaying{
		OrderID:         orderResp.Info.ID,
		PaymentID:       paymentResp.Info.ID,
		FeeTypeID:       uuid.New().String(),
		DurationMinutes: 24 * 10 * 60,
	}
	gasPayingResp2, err := gaspaying.Create(context.Background(), &npool.CreateGasPayingRequest{
		Info: &gasPaying2,
	})
	assert.Nil(t, err)
	gasPaying2.ID = gasPayingResp2.Info.ID

	compensate1 := npool.Compensate{
		OrderID: orderResp.Info.ID,
		Start:   uint32(time.Now().Unix()),
		End:     uint32(time.Now().Unix()),
	}
	compensateResp1, err := compensate.Create(context.Background(), &npool.CreateCompensateRequest{
		Info: &compensate1,
	})
	assert.Nil(t, err)

	compensate2 := npool.Compensate{
		OrderID: orderResp.Info.ID,
		Start:   uint32(time.Now().Unix()) + 20,
		End:     uint32(time.Now().Unix()) + 29,
	}
	compensateResp2, err := compensate.Create(context.Background(), &npool.CreateCompensateRequest{
		Info: &compensate2,
	})
	assert.Nil(t, err)

	outOfGas1 := npool.OutOfGas{
		OrderID: orderResp.Info.ID,
		Start:   uint32(time.Now().Unix()) + 20,
		End:     uint32(time.Now().Unix()) + 29,
	}
	outOfGasResp1, err := outofgas.Create(context.Background(), &npool.CreateOutOfGasRequest{
		Info: &outOfGas1,
	})
	assert.Nil(t, err)

	outOfGas2 := npool.OutOfGas{
		OrderID: orderResp.Info.ID,
		Start:   uint32(time.Now().Unix()) + 20,
		End:     uint32(time.Now().Unix()) + 29,
	}
	outOfGasResp2, err := outofgas.Create(context.Background(), &npool.CreateOutOfGasRequest{
		Info: &outOfGas2,
	})
	assert.Nil(t, err)

	orderDetail, err := Get(context.Background(), &npool.GetOrderDetailRequest{
		ID: orderResp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, orderDetail.Detail.ID, orderResp.Info.ID)
		assertOrderDetail(
			t,
			orderDetail.Detail,
			&myOrder,
			goodPayingResp.Info,
			[]*npool.GasPaying{
				gasPayingResp1.Info,
				gasPayingResp2.Info,
			},
			[]*npool.Compensate{
				compensateResp1.Info,
				compensateResp2.Info,
			},
			[]*npool.OutOfGas{
				outOfGasResp1.Info,
				outOfGasResp2.Info,
			},
			paymentResp.Info)
	}

	orderDetails, err := GetByAppUser(context.Background(), &npool.GetOrdersDetailByAppUserRequest{
		AppID:  appID,
		UserID: userID,
	})
	if assert.Nil(t, err) {
		if assert.Equal(t, len(orderDetails.Details), 1) {
			assert.Equal(t, orderDetails.Details[0].ID, orderResp.Info.ID)
			assertOrderDetail(
				t,
				orderDetail.Detail,
				&myOrder,
				goodPayingResp.Info,
				[]*npool.GasPaying{
					gasPayingResp1.Info,
					gasPayingResp2.Info,
				},
				[]*npool.Compensate{
					compensateResp1.Info,
					compensateResp2.Info,
				},
				[]*npool.OutOfGas{
					outOfGasResp1.Info,
					outOfGasResp2.Info,
				},
				paymentResp.Info)
		}
	}

	orderDetails1, err := GetByApp(context.Background(), &npool.GetOrdersDetailByAppRequest{
		AppID: appID,
	})
	if assert.Nil(t, err) {
		if assert.Equal(t, len(orderDetails1.Details), 1) {
			assert.Equal(t, orderDetails1.Details[0].ID, orderResp.Info.ID)
			assertOrderDetail(
				t,
				orderDetail.Detail,
				&myOrder,
				goodPayingResp.Info,
				[]*npool.GasPaying{
					gasPayingResp1.Info,
					gasPayingResp2.Info,
				},
				[]*npool.Compensate{
					compensateResp1.Info,
					compensateResp2.Info,
				},
				[]*npool.OutOfGas{
					outOfGasResp1.Info,
					outOfGasResp2.Info,
				},
				paymentResp.Info)
		}
	}

	orderDetails2, err := GetByGood(context.Background(), &npool.GetOrdersDetailByGoodRequest{
		GoodID: goodID,
	})
	if assert.Nil(t, err) {
		if assert.Equal(t, len(orderDetails2.Details), 1) {
			assert.Equal(t, orderDetails2.Details[0].ID, orderResp.Info.ID)
			assertOrderDetail(
				t,
				orderDetail.Detail,
				&myOrder,
				goodPayingResp.Info,
				[]*npool.GasPaying{
					gasPayingResp1.Info,
					gasPayingResp2.Info,
				},
				[]*npool.Compensate{
					compensateResp1.Info,
					compensateResp2.Info,
				},
				[]*npool.OutOfGas{
					outOfGasResp1.Info,
					outOfGasResp2.Info,
				},
				paymentResp.Info)
		}
	}
}

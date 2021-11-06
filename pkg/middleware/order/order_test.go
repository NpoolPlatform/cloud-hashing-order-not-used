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

func TestGetDetail(t *testing.T) {
	second := uint32(time.Now().Unix())

	myOrder := npool.Order{
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
		GasPayIDs:                []string{uuid.UUID{}.String(), uuid.UUID{}.String()},
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
	_, err = goodpaying.Create(context.Background(), &npool.CreateGoodPayingRequest{
		Info: &goodPaying,
	})
	assert.Nil(t, err)

	gasPaying := npool.GasPaying{
		OrderID:         orderResp.Info.ID,
		AccountID:       uuid.New().String(),
		DurationMinutes: 24 * 10 * 60,
	}
	_, err = gaspaying.Create(context.Background(), &npool.CreateGasPayingRequest{
		Info: &gasPaying,
	})
	assert.Nil(t, err)
}

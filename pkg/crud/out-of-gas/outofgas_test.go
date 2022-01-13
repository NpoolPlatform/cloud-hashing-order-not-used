package outofgas

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-order"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/test-init" //nolint

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

func assertOutOfGas(t *testing.T, actual, expected *npool.OutOfGas) {
	assert.Equal(t, actual.OrderID, expected.OrderID)
	assert.Equal(t, actual.Start, expected.Start)
	assert.Equal(t, actual.End, expected.End)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	outOfGas := npool.OutOfGas{
		OrderID: uuid.New().String(),
		Start:   uint32(time.Now().Unix()),
		End:     uint32(time.Now().Unix()),
	}
	resp, err := Create(context.Background(), &npool.CreateOutOfGasRequest{
		Info: &outOfGas,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertOutOfGas(t, resp.Info, &outOfGas)
	}

	resp1, err := GetByOrder(context.Background(), &npool.GetOutOfGasesByOrderRequest{
		OrderID: outOfGas.OrderID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp1.Infos), 1)
	}
}

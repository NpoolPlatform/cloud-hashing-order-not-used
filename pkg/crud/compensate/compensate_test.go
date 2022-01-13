package compensate

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/cloud-hashing-order/pkg/test-init" //nolint
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-order"

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

func assertCompensate(t *testing.T, actual, expected *npool.Compensate) {
	assert.Equal(t, actual.OrderID, expected.OrderID)
	assert.Equal(t, actual.Start, expected.Start)
	assert.Equal(t, actual.End, expected.End)
	assert.Equal(t, actual.Message, expected.Message)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	compensate := npool.Compensate{
		OrderID: uuid.New().String(),
		Start:   uint32(time.Now().Unix()),
		End:     uint32(time.Now().Unix()),
		Message: "Test compensate message",
	}
	resp, err := Create(context.Background(), &npool.CreateCompensateRequest{
		Info: &compensate,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertCompensate(t, resp.Info, &compensate)
	}

	resp1, err := GetByOrder(context.Background(), &npool.GetCompensatesByOrderRequest{
		OrderID: compensate.OrderID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp1.Infos), 1)
	}
}

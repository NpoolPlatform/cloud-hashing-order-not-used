package order

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-order/message/npool"

	_ "github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/order" //nolint

	_ "github.com/google/uuid" //nolint

	_ "golang.org/x/xerrors" //nolint
)

func Get(ctx context.Context, in *npool.GetOrderDetailRequest) (*npool.GetOrderDetailResponse, error) {
	return nil, nil
}

func GetOrdersDetailByAppUser(ctx context.Context, in *npool.GetOrdersDetailByAppUserRequest) (*npool.GetOrdersDetailByAppUserResponse, error) {
	return nil, nil
}

func GetOrdersDetailByApp(ctx context.Context, in *npool.GetOrdersDetailByAppRequest) (*npool.GetOrdersDetailByAppResponse, error) {
	return nil, nil
}

func GetOrdersDetailByGood(ctx context.Context, in *npool.GetOrdersDetailByGoodRequest) (*npool.GetOrdersDetailByGoodResponse, error) {
	return nil, nil
}

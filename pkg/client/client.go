package client

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/cloud-hashing-order/pkg/message/const"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-order"
)

func do(ctx context.Context, fn func(_ctx context.Context, cli npool.CloudHashingOrderClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get good payment connection: %v", err)
	}
	defer conn.Close()

	cli := npool.NewCloudHashingOrderClient(conn)

	return fn(_ctx, cli)
}

func GetOrders(ctx context.Context, offset, limit int32) ([]*npool.Order, error) {
	// conds: NOT USED NOW, will be used after refactor code
	infos, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingOrderClient) (cruder.Any, error) {
		resp, err := cli.GetOrders(ctx, &npool.GetOrdersRequest{
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get orders: %v", err)
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get orders: %v", err)
	}
	return infos.([]*npool.Order), nil
}

func GetOrderPayment(ctx context.Context, orderID string) (*npool.Payment, error) {
	// conds: NOT USED NOW, will be used after refactor code
	info, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingOrderClient) (cruder.Any, error) {
		resp, err := cli.GetPaymentByOrder(ctx, &npool.GetPaymentByOrderRequest{
			OrderID: orderID,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get payment: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get payment: %v", err)
	}
	return info.(*npool.Payment), nil
}

func GetStatePayments(ctx context.Context, state string) ([]*npool.Payment, error) {
	// conds: NOT USED NOW, will be used after refactor code
	infos, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingOrderClient) (cruder.Any, error) {
		resp, err := cli.GetPaymentsByState(ctx, &npool.GetPaymentsByStateRequest{
			State: state,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get payments: %v", err)
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get payments: %v", err)
	}
	return infos.([]*npool.Payment), nil
}

func UpdatePayment(ctx context.Context, in *npool.Payment) (*npool.Payment, error) {
	// conds: NOT USED NOW, will be used after refactor code
	info, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingOrderClient) (cruder.Any, error) {
		resp, err := cli.UpdatePayment(ctx, &npool.UpdatePaymentRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail update payment: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update payment: %v", err)
	}
	return info.(*npool.Payment), nil
}

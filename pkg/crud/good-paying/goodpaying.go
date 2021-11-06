package goodpaying

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-order/message/npool"

	_ "github.com/google/uuid"

	_ "golang.org/x/xerrors"
)

func Create(ctx context.Context, in *npool.CreateGoodPayingRequest) (*npool.CreateGoodPayingResponse, error) {
	return nil, nil
}

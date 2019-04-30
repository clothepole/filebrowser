package das

import (
	"context"

	"google.golang.org/grpc"

	core "go.hxyy.com/core/v1"
)

// Cabinet ...
type Cabinet struct {
	StoreClient core.StoreClient
}

// NewCabinet ...
func NewCabinet(ctx context.Context, addr string) (*Cabinet, error) {
	cabinet := &Cabinet{}

	if cc, err := grpc.DialContext(ctx, addr, grpc.WithInsecure()); err == nil {
		cabinet.StoreClient = core.NewStoreClient(cc)
	} else {
		return nil, err
	}

	return cabinet, nil
}

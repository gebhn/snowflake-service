package snowflake

import "context"

type Creator interface {
	Create(ctx context.Context) (string, error)
}

var _ Creator = (*storeCreator)(nil)

package snowflake

import (
	"context"
	"strconv"
	"time"

	"github.com/uplite/snowflake-service/internal/store"
)

type storeCreator struct {
	store store.Store
}

func NewStoreCreator(store store.Store) *storeCreator {
	return &storeCreator{
		store: store,
	}
}

func (c *storeCreator) Create(ctx context.Context) (string, error) {
	// TODO @gebhartn @aapclark: generate a snowflake here
	fakeSnowflake := strconv.FormatInt(time.Now().Unix(), 10)

	if err := c.store.Create(ctx, fakeSnowflake); err != nil {
		return "", err
	}

	return fakeSnowflake, nil
}

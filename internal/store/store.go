package store

import "context"

type Store interface {
	Create(context.Context, string) error
	Get(context.Context, string) (string, error)
}

var _ Store = (*snowflakeStore)(nil)

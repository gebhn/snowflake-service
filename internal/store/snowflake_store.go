package store

import (
	"context"

	"github.com/uplite/snowflake-service/internal/db"
)

type snowflakeStore struct {
	queries *db.Queries
}

func NewSnowflakeStore(conn db.DBTX) *snowflakeStore {
	return &snowflakeStore{
		queries: db.NewQueries(conn),
	}
}

func (s *snowflakeStore) Create(ctx context.Context, id string) error {
	return s.queries.Create(ctx, id)
}

func (s *snowflakeStore) Get(ctx context.Context, id string) (string, error) {
	return s.queries.Get(ctx, id)
}

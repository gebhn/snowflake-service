package snowflake

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockStore struct{}

func (s *mockStore) Create(context.Context, string) error {
	return nil
}

func (s *mockStore) Get(ctx context.Context, snowflake string) (string, error) {
	return snowflake, nil
}

func TestStoreCreator(t *testing.T) {
	c := NewStoreCreator(new(mockStore))

	t.Run("should create", func(t *testing.T) {
		snowflake, err := c.Create(context.Background())
		assert.NoError(t, err)
		assert.NotEmpty(t, snowflake)
	})
}

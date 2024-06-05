package store

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

const mockSnowflake = "mockSnowflake"

func TestSnowflakeStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	s := NewSnowflakeStore(db)

	t.Run("should create snowflake", func(t *testing.T) {
		var err error

		mock.ExpectExec("INSERT INTO snowflakes").WithArgs(mockSnowflake).WillReturnResult(sqlmock.NewResult(1, 1))

		err = s.Create(context.Background(), mockSnowflake)
		assert.NoError(t, err)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("should get snowflake", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id"}).AddRow(mockSnowflake)

		mock.ExpectQuery("SELECT id FROM snowflakes WHERE id = ?").WithArgs(mockSnowflake).WillReturnRows(rows)

		res, err := s.Get(context.Background(), mockSnowflake)
		assert.NoError(t, err)
		assert.Equal(t, mockSnowflake, res)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

package db

import "github.com/uplite/snowflake-service/internal/db/sqlc"

// DBTX Interface
type DBTX = sqlc.DBTX

// Query implementation
type Queries = sqlc.Queries

var NewQueries = sqlc.New

var _ DBTX = (*libsqlDb)(nil)

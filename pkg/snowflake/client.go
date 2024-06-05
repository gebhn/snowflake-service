package snowflake

import (
	"io"

	"github.com/uplite/snowflake-service/api/pb"
)

type Client interface {
	pb.SnowflakeServiceClient
	io.Closer
}

var _ Client = (*snowflakeClient)(nil)

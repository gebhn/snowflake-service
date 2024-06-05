package snowflake

import (
	"context"

	"google.golang.org/grpc"

	"github.com/uplite/snowflake-service/api/pb"
)

type snowflakeClient struct {
	conn   *grpc.ClientConn
	client pb.SnowflakeServiceClient
}

func New(conn *grpc.ClientConn) *snowflakeClient {
	return &snowflakeClient{
		conn:   conn,
		client: pb.NewSnowflakeServiceClient(conn),
	}
}

func (c *snowflakeClient) Create(ctx context.Context, req *pb.CreateRequest, opts ...grpc.CallOption) (*pb.CreateResponse, error) {
	return c.client.Create(ctx, req, opts...)
}

func (c *snowflakeClient) Close() error {
	return c.conn.Close()
}

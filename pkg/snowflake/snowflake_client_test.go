package snowflake

import (
	"context"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uplite/snowflake-service/api/pb"
	"github.com/uplite/snowflake-service/internal/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const mockSnowflake = "mock_snowflake"

type mockCreator struct{}

func (m *mockCreator) Create(ctx context.Context) (string, error) { return mockSnowflake, nil }

func TestSnowflakeClient(t *testing.T) {
	srv := server.NewSnowflakeServer(new(mockCreator))

	grpcServer := grpc.NewServer()

	pb.RegisterSnowflakeServiceServer(grpcServer, srv)

	lis, err := net.Listen("tcp", ":50053")
	assert.NoError(t, err)

	go grpcServer.Serve(lis)
	defer grpcServer.Stop()

	conn, err := grpc.NewClient(":50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoError(t, err)

	c := New(conn)

	t.Run("should create a snowflake", func(t *testing.T) {
		res, err := c.Create(context.Background(), &pb.CreateRequest{})
		assert.NoError(t, err)
		assert.Equal(t, mockSnowflake, res.GetSnowflake())
	})
}

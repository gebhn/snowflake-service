package server

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/uplite/snowflake-service/api/pb"
	"github.com/uplite/snowflake-service/internal/config"
)

const mockSnowflake = "mock_snowflake"

type mockCreator struct{}

func (m *mockCreator) Create(ctx context.Context) (string, error) { return mockSnowflake, nil }

func TestSnowflakeServer(t *testing.T) {
	grpcServer := grpc.NewServer()

	s := NewSnowflakeServer(&mockCreator{})
	pb.RegisterSnowflakeServiceServer(grpcServer, s)

	lis, err := net.Listen("tcp", ":"+config.GetGrpcServerPort())
	if err != nil {
		t.Fatal(err)
	}

	go grpcServer.Serve(lis)
	defer grpcServer.Stop()

	conn, err := grpc.NewClient(":"+config.GetGrpcServerPort(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewSnowflakeServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	t.Run("should create one", func(t *testing.T) {
		res, err := client.Create(ctx, nil)
		assert.NoError(t, err)
		assert.Equal(t, res.GetSnowflake(), mockSnowflake)
	})
}

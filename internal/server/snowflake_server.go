package server

import (
	"context"

	"github.com/uplite/snowflake-service/api/pb"
	"github.com/uplite/snowflake-service/internal/snowflake"
)

type snowflakeServer struct {
	pb.UnimplementedSnowflakeServiceServer
	creator snowflake.Creator
}

func NewSnowflakeServer(creator snowflake.Creator) *snowflakeServer {
	return &snowflakeServer{
		creator: creator,
	}
}

func (s *snowflakeServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	snowflake, err := s.creator.Create(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.CreateResponse{Snowflake: snowflake}, nil
}

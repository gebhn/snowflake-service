package service

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/uplite/snowflake-service/api/pb"
	"github.com/uplite/snowflake-service/internal/config"
	"github.com/uplite/snowflake-service/internal/db"
	"github.com/uplite/snowflake-service/internal/server"
	"github.com/uplite/snowflake-service/internal/snowflake"
	"github.com/uplite/snowflake-service/internal/store"
)

type snowflakeService struct {
	grpcServer      *grpc.Server
	snowflakeServer pb.SnowflakeServiceServer
}

func NewSnowflakeService() *snowflakeService {
	g := grpc.NewServer()
	d := db.NewLibsqlConn(config.GetTursoDbToken(), config.GetTursoDbToken())
	s := store.NewSnowflakeStore(d)
	c := snowflake.NewStoreCreator(s)

	snowflakeServer := server.NewSnowflakeServer(c)

	pb.RegisterSnowflakeServiceServer(g, snowflakeServer)

	return &snowflakeService{
		grpcServer:      g,
		snowflakeServer: snowflakeServer,
	}
}

func (s *snowflakeService) Serve() error {
	lis, err := net.Listen("tcp", ":"+config.GetGrpcServerPort())
	if err != nil {
		log.Fatal(err)
	}

	return s.grpcServer.Serve(lis)
}

func (s *snowflakeService) Close() {
	s.grpcServer.GracefulStop()
}

package grpc

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/turnixxd/grpc-test/server/database"
	"github.com/turnixxd/grpc-test/server/env"
	pb "github.com/turnixxd/grpc-test/server/proto"
	"google.golang.org/grpc"
)

// Constructor
func NewBasicServiceServer() *BasicServiceServer {
	return &BasicServiceServer{}
}

type BasicServiceServer struct {
	pb.BasicServiceServer
}

var db database.Database

func (server *BasicServiceServer) Run() error {
	lis, err := net.Listen("tcp", ":"+env.Process("SERVER_PORT"))
	if err != nil {
		log.Fatal("Failed to listen on port %v, %v", env.Process("SERVER_PORT"), err)
	}

	s := grpc.NewServer()
	databaseImplementation := os.Args[1]
	db, err = database.Factory(databaseImplementation)
	if err != nil {
		panic(err)
	}

	// cannot be type of &UserManagementServer{} because we are using the server that Run() receives
	pb.RegisterBasicServiceServer(s, server)
	log.Printf("gRPC server listening at %v", lis.Addr())
	return s.Serve(lis)
}

func (s *BasicServiceServer) Set(ctx context.Context, in *pb.SetRequest) (*pb.ServerResponse, error) {
	value, err := db.Set(in.GetKey(), in.GetValue())
	return generateResponse(value, err)
}

func (s *BasicServiceServer) Get(ctx context.Context, in *pb.GetRequest) (*pb.ServerResponse, error) {
	value, err := db.Get(in.GetKey())
	return generateResponse(value, err)
}

func (s *BasicServiceServer) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.ServerResponse, error) {
	value, err := db.Delete(in.GetKey())
	return generateResponse(value, err)
}

func generateResponse(value string, err error) (*pb.ServerResponse, error) {
	if err != nil {
		return &pb.ServerResponse{Success: false, Value: value, Error: err.Error()}, nil
	}
	return &pb.ServerResponse{Success: true, Value: value, Error: ""}, nil
}

func Serve() {
	var basic_service_server *BasicServiceServer = NewBasicServiceServer()
	if err := basic_service_server.Run(); err != nil {
		log.Fatal("Failed to serve: %v", err)
	}
}

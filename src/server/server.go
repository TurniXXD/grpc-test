package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	// ProtoBuff
	pb "github.com/turnixxd/grpc-test/proto"
	"github.com/turnixxd/grpc-test/server/env"
	"google.golang.org/grpc"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())
	var user_id int32 = int32(rand.Intn(1000))
	return &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: user_id}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+env.Process("PORT"))
	if err != nil {
		log.Fatal("Failed to listen on port %v, %v", env.Process("PORT"), err)
	}

	s := grpc.NewServer()

	pb.RegisterUserManagementServer(s, &UserManagementServer{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve gRPC server over port %v", env.Process("PORT"))
	}
}

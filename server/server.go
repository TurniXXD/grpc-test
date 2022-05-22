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

// Constructor
func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{
		user_list: &pb.UserList{},
	}
}

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
	user_list *pb.UserList
}

// Server instance
func (server *UserManagementServer) Run() error {
	lis, err := net.Listen("tcp", ":"+env.Process("PORT"))
	if err != nil {
		log.Fatal("Failed to listen on port %v, %v", env.Process("PORT"), err)
	}

	s := grpc.NewServer()

	// cannot be type of &UserManagementServer{} because we are using the server that Run() receives
	pb.RegisterUserManagementServer(s, server)
	log.Printf("gRPC server listening at %v", lis.Addr())
	return s.Serve(lis)
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())
	var user_id int32 = int32(rand.Intn(1000))
	created_user := &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: user_id}
	// Append created user to user list
	s.user_list.Users = append(s.user_list.Users, created_user)
	return created_user, nil
}

func (s *UserManagementServer) GetUsers(ctx context.Context, in *pb.GetUsersParams) (*pb.UserList, error) {
	return s.user_list, nil
}

func main() {
	var user_mgmt_server *UserManagementServer = NewUserManagementServer()
	if err := user_mgmt_server.Run(); err != nil {
		log.Fatal("Failed to serve: %v", err)
	}
}

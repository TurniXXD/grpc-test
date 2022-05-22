package main

import (
	"context"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"

	// ProtoBuff
	pb "github.com/turnixxd/grpc-test/proto"
	"github.com/turnixxd/grpc-test/server/env"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

// Constructor
func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{}
}

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
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

func ProtoJsonMarshall(users_list *pb.UserList) {
	jsonBytes, err := protojson.Marshal(users_list)
	if err != nil {
		log.Fatalf("JSON Marshaling failed: %v", err)
	}
	if err := ioutil.WriteFile("users.json", jsonBytes, 0664); err != nil {
		log.Fatalf("Failed write to file: %v", err)
	}
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())
	readBytes, err := ioutil.ReadFile("users.json")
	var users_list *pb.UserList = &pb.UserList{}
	var user_id int32 = int32(rand.Intn(1000))
	created_user := &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: user_id}

	if err != nil {
		if os.IsNotExist(err) {
			log.Print("File not found. Creating a new one")
			users_list.Users = append(users_list.Users, created_user)
			ProtoJsonMarshall(users_list)
			return created_user, nil
		} else {
			log.Fatalln("Error reading file: ", err)
		}
	}

	if err := protojson.Unmarshal(readBytes, users_list); err != nil {
		log.Fatal("Failed to parse user list: %v", err)
	}

	users_list.Users = append(users_list.Users, created_user)
	ProtoJsonMarshall(users_list)

	return created_user, nil
}

func (s *UserManagementServer) GetUsers(ctx context.Context, in *pb.GetUsersParams) (*pb.UserList, error) {
	jsonBytes, err := ioutil.ReadFile("users.json")
	if err != nil {
		log.Fatal("failed read from file: %v", err)
	}

	var users_list *pb.UserList = &pb.UserList{}
	if err := protojson.Unmarshal(jsonBytes, users_list); err != nil {
		log.Fatalf("Unmarshaling failed: %v", err)
	}

	return users_list, nil
}

func main() {
	var user_mgmt_server *UserManagementServer = NewUserManagementServer()
	if err := user_mgmt_server.Run(); err != nil {
		log.Fatal("Failed to serve: %v", err)
	}
}

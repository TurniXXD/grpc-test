package main

import (
	"context"
	"log"
	"time"

	"github.com/turnixxd/grpc-test/client/env"
	// ProtoBuff
	pb "github.com/turnixxd/grpc-test/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(env.Process("ADDR"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Failed to connect %v", err)
	}
	defer conn.Close()

	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var new_users = make(map[string]int32)

	new_users["Alice"] = 43
	new_users["Bob"] = 30

	for name, age := range new_users {
		r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
		if err != nil {
			log.Fatal("Could not create user: %v", err)
		}
		log.Printf("\n\nUser Details: \nNAME: %s\nAGE: %d\nID: %d", r.GetName(), r.GetAge(), r.GetId())
	}

	params := &pb.GetUsersParams{}
	r, err := c.GetUsers(ctx, params)
	if err != nil {
		log.Fatal("Could not retrieve users: %v", err)
	}
	log.Printf("\n\nUser list: \n%v", r.GetUsers())
}

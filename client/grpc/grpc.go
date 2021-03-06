package grpc

import (
	"context"
	"log"
	"time"

	"github.com/turnixxd/grpc-test/client/env"
	// ProtoBuff
	pb "github.com/turnixxd/grpc-test/client/proto"
	"google.golang.org/grpc"
)

func CreateSetRequest(key, value string) {
	conn, err := grpc.Dial(env.Process("ADDR"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Failed to connect %v", err)
	}
	defer conn.Close()

	c := pb.NewBasicServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// var new_users = make(map[string]int32)

	// new_users["Alice"] = 43
	// new_users["Bob"] = 30

	// for name, age := range new_users {
	// 	r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
	// 	if err != nil {
	// 		log.Fatal("Could not create user: %v", err)
	// 	}
	// 	log.Printf("\n\nUser Details: \nNAME: %s\nAGE: %d\nID: %d", r.GetName(), r.GetAge(), r.GetId())
	// }

	r, err := c.Set(ctx, &pb.SetRequest{Key: key, Value: value})
	if err != nil {
		log.Fatal("Could not create: %v", err)
	}
	log.Printf("\n\nService request: \nState: %v\nValue: %v", r.GetSuccess(), r.GetValue())
}

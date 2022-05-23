package main

import (
	"time"

	client "github.com/turnixxd/grpc-test/client/grpc"
)

func main() {
	time.Sleep(2 * time.Second)
	client.CreateSetRequest("name", "Jakub")
	client.CreateSetRequest("name", "Tom")
	client.CreateSetRequest("name", "Jana")
}

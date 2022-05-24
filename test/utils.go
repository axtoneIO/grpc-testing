// +acceptance

package test

import (
	"log"

	rocket "github.com/axtoneIO/grpc-testing/protos/rocket/v1"
	"google.golang.org/grpc"
)

// GetClient - will get the client for testing purposes
func GetClient() rocket.RocketServiceClient {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:50051",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("unable to connect %s",err)
	}

	rocketClient := rocket.NewRocketServiceClient(conn)
	return rocketClient
}
package main

import (
	"log"

	"github.com/axtoneIO/grpc-testing/internal/db"
	"github.com/axtoneIO/grpc-testing/internal/rocket"
	"github.com/axtoneIO/grpc-testing/internal/transport/grpc"
)

func Run() error {
	// responsible for initializing and starting
	// our gRPC server
	rocketStore, err := db.New()
	if err != nil {
		return err
	}
	// responsible for the migration process
	err = rocketStore.Migrate()
	if err != nil {
		log.Println("Failed to run migrations")
		return err
	}

	rktService := rocket.New(rocketStore)
	rktHandler := grpc.New(rktService)

	if err := rktHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

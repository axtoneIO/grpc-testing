package grpc

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/axtoneIO/grpc-testing/internal/rocket"
	rkt "github.com/axtoneIO/grpc-testing/protos/rocket/v1"
	"google.golang.org/grpc"
)

type RocketService interface {
	GetRocket(ctx context.Context, id string) (rocket.Rocket, error)
	AddRocket(ctx context.Context, rkt rocket.Rocket) (rocket.Rocket, error)
	DeleteRocket(ctx context.Context, id string) (string, error)
}

// Handler - will handle incoming gRPC requests
type Handler struct {
	RocketService RocketService
}

func New(rktService RocketService) Handler {
	return Handler{
		RocketService: rktService,
	}
}

func (h Handler) Serve() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Print("could not listen on port 50051")
		return err
	}

	grpcServer := grpc.NewServer()
	rkt.RegisterRocketServiceServer(grpcServer, &h)

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve %v", err)
		return err
	}

	return nil
}

func (h Handler) GetRocket(ctx context.Context, req *rkt.GetRocketRequest) (*rkt.GetRocketResponse, error) {
	id := req.Id
	if id == "" {
		return &rkt.GetRocketResponse{}, errors.New("id parameter must be sent")
	}

	rocket, err := h.RocketService.GetRocket(ctx, req.Id)
	if err != nil {
		return &rkt.GetRocketResponse{}, err
	}

	return &rkt.GetRocketResponse{
		Rocket: &rkt.Rocket{
			ID:   rocket.ID,
			Name: rocket.Name,
			Type: rocket.Type,
		},
	}, nil
}

func (h Handler) AddRocket(ctx context.Context, req *rkt.AddRocketRequest) (*rkt.AddRocketResponse, error) {
	if req.Rocket == nil {
		return &rkt.AddRocketResponse{},errors.New("rocket parameter must be sent")
	}

	newRkt, err := h.RocketService.AddRocket(ctx, rocket.Rocket{
		ID:   req.Rocket.ID,
		Name: req.Rocket.Name,
		Type: req.Rocket.Type,
	})
	if err != nil {
		return nil, errors.New("failed to insert rocket into database")
	}
	fmt.Println(newRkt)
	return &rkt.AddRocketResponse{
		Rocket: &rkt.Rocket{
			ID:   newRkt.ID,
			Name: newRkt.Name,
			Type: newRkt.Type,
		},
	}, nil
}

func (h Handler) DeleteRocket(ctx context.Context, req *rkt.DeleteRocketRequest) (*rkt.DeleteRocketResponse, error) {
	if req.Id == "" {
		return &rkt.DeleteRocketResponse{}, errors.New("id parameter must be sent")
	}

	status, err := h.RocketService.DeleteRocket(ctx,req.Id)
	if err != nil{
		return nil, errors.New("failed to delete rocket from database")
	}

	return &rkt.DeleteRocketResponse{
		Status: status,
	}, nil
}

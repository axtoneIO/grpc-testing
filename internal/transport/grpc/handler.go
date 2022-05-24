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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RocketService interface {
	GetRocket(ctx context.Context, id int64) (rocket.Rocket, error)
	AddRocket(ctx context.Context, rkt rocket.Rocket) (rocket.Rocket, error)
	DeleteRocket(ctx context.Context, id int64) (string, error)
}

// Handler - will handle incoming gRPC requests
type Handler struct {
	RocketService RocketService
}

// New - represents the container func for the service chosen
func New(rktService RocketService) Handler {
	return Handler{
		RocketService: rktService,
	}
}

// Serve - will start the microservice 
func (h Handler) Serve() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Print("could not listen on port 50051")
		return err
	}

	grpcServer := grpc.NewServer()
	rkt.RegisterRocketServiceServer(grpcServer, &h)

	if err := grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}

// GetRocket - 
func (h Handler) GetRocket(ctx context.Context, req *rkt.GetRocketRequest) (*rkt.GetRocketResponse, error) {
	if req.Id == 0 {
		errorStatus := status.Error(codes.InvalidArgument, "id parameter is not valid")
		return &rkt.GetRocketResponse{}, errorStatus
	}

	rocket, err := h.RocketService.GetRocket(ctx, req.Id)
	if err != nil {
		return &rkt.GetRocketResponse{}, err
	}

	return &rkt.GetRocketResponse{
		Rocket: &rkt.Rocket{
			Id:   rocket.Id,
			Name: rocket.Name,
			Type: rocket.Type,
		},
	}, nil
}

func (h Handler) AddRocket(ctx context.Context, req *rkt.AddRocketRequest) (*rkt.AddRocketResponse, error) {
	if req.Rocket == nil {
		errorStatus := status.Error(codes.InvalidArgument, "rocket parameter is not valid")
		return &rkt.AddRocketResponse{}, errorStatus
	}

	newRkt, err := h.RocketService.AddRocket(ctx, rocket.Rocket{
		Id:   req.Rocket.Id,
		Name: req.Rocket.Name,
		Type: req.Rocket.Type,
	})
	if err != nil {
		return nil, errors.New("failed to insert rocket into database")
	}
	fmt.Println(newRkt)
	return &rkt.AddRocketResponse{
		Rocket: &rkt.Rocket{
			Id:   newRkt.Id,
			Name: newRkt.Name,
			Type: newRkt.Type,
		},
	}, nil
}

func (h Handler) DeleteRocket(ctx context.Context, req *rkt.DeleteRocketRequest) (*rkt.DeleteRocketResponse, error) {
	if req.Id == 0 {
		errorStatus := status.Error(codes.InvalidArgument, "id parameter is not valid")
		return &rkt.DeleteRocketResponse{}, errorStatus
	}

	status, err := h.RocketService.DeleteRocket(ctx, req.Id)
	if err != nil {
		return nil, errors.New("failed to delete rocket from database")
	}

	return &rkt.DeleteRocketResponse{
		Status: status,
	}, nil
}

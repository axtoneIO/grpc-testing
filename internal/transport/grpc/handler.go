package grpc

import (
	"context"
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
	rkt.RegisterRocketServiceServer(grpcServer,&h)

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve %v", err)
		return err
	}
	
	return nil
}

func (h Handler) GetRocket(ctx context.Context, req *rkt.GetRocketRequest) (*rkt.GetRocketResponse,error){
	log.Println("Get Rocket gRPC Endpoint Hit")

	rocket, err := h.RocketService.GetRocket(ctx,req.Id)
	if err != nil {
		return &rkt.GetRocketResponse{},err
	}
	
	return &rkt.GetRocketResponse{
		Rocket: &rkt.Rocket{
			ID: 	rocket.ID,
			Name:	rocket.Name,
			Type: 	rocket.Type,
		},
	},nil
}

func (h Handler) AddRocket(ctx context.Context, req *rkt.AddRocketRequest) (*rkt.AddRocketResponse,error){
	return &rkt.AddRocketResponse{},nil
}

func (h Handler) DeleteRocket(ctx context.Context, req *rkt.DeleteRocketRequest) (*rkt.DeleteRocketResponse,error){
	return &rkt.DeleteRocketResponse{},nil
}

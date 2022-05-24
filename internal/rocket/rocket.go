//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/axtoneIO/grpc-testing/internal/rocket Store
package rocket

import "context"

// Rocket - should contain the definition of our
// rocket
type Rocket struct {
	Id		int64
	Name	string
	Type    string
}

// store defines the interface we expect
// our database implementation to follow
type Store interface {
	GetRocket(id int64) (Rocket, error)
	AddRocket(rkt Rocket) (Rocket, error)
	DeleteRocket(id int64) (string,error)
}

// Service - our rocket service, responsible for
// updating the rocket inventory
type Service struct {
	Store Store
}

// New - returns a new instance of out rocket service
func New(store Store) Service {
	return Service{
		Store: store,
	}
}

// GetRocketById - retrieves a rocket by id from the store
func (s Service) GetRocket(ctx context.Context, id int64) (Rocket, error) {
	rkt, err := s.Store.GetRocket(id)

	if err != nil {
		return Rocket{}, err
	}

	return rkt, nil
}

// InsertRocket - inserts a new rocket into the store
func (s Service) AddRocket(ctx context.Context, rkt Rocket) (Rocket, error) {
	rkt, err := s.Store.AddRocket(rkt)

	if err != nil {
		return Rocket{}, err
	}

	return rkt, err
}

// DeleteRocket - deletes a rocket by id from the store
func (s Service) DeleteRocket(ctx context.Context, id int64) (string,error) {
	status,err := s.Store.DeleteRocket(id)

	if err != nil {
		return status,err
	}

	return status,nil
}

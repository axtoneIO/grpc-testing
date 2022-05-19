//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/axtoneIO/grpc-testing/internal/rocket Store
package rocket

import "context"

// Rocket - should contain the definition of our
// rocket
type Rocket struct {
	ID      string
	Name    string
	Type    string
	Flights int
}

// store defines the interface we expect
// our database implementation to follow
type Store interface {
	GetRocketById(id string) (Rocket, error)
	InsertRocket(rkt Rocket) (Rocket, error)
	DeleteRocket(id string) error
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
func (s Service) GetRocketById(ctx context.Context, id string) (Rocket, error) {
	rkt, err := s.Store.GetRocketById(id)

	if err != nil {
		return Rocket{}, err
	}

	return rkt, nil
}

// InsertRocket - inserts a new rocket into the store
func (s Service) InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error) {
	rkt, err := s.Store.InsertRocket(rkt)

	if err != nil {
		return Rocket{}, err
	}

	return rkt, err
}

// DeleteRocket - deletes a rocket by id from the store
func (s Service) DeleteRocket(id string) error {
	err := s.Store.DeleteRocket(id)

	if err != nil {
		return err
	}

	return nil
}

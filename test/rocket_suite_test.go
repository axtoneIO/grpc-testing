// +acceptance

package test

import (
	"context"
	"testing"

	rocket "github.com/axtoneIO/grpc-testing/protos/rocket/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RocketTestSuite struct {
	suite.Suite
}

func (s *RocketTestSuite) TestAddRocket() {
	s.T().Run("Adds a new rocket successfully", func(t *testing.T) {
		client := GetClient()
		resp, err := client.AddRocket(
			context.Background(),
			&rocket.AddRocketRequest{
				Rocket: &rocket.Rocket{
					ID:   "5",
					Name: "a",
					Type: "a",
				},
			},
		)
		assert.NoError(s.T(), err)
		assert.Equal(s.T(),rocket.Rocket{
			ID:   "5",
			Name: "a",
			Type: "a",
		},resp.Rocket)
	})
}

func TestRocketService(t *testing.T) {
	suite.Run(t, new(RocketTestSuite))
}

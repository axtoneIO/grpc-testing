// +acceptance

package test

import (
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

type RocketTestSuite struct {
	suite.Suite
}

func(s *RocketTestSuite) TestAddRocket() {
	log.Println("Hello World")
}

func TestRocketService(t *testing.T){
	suite.Run(t,new(RocketTestSuite))
}

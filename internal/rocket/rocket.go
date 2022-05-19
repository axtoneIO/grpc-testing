package rocket

// Rocket - should contain the definition of our
// rocket 
type Rocket struct {
	ID 		string
	Name 	string
	Type 	string
	Flights int
}

// Service - our rocket service, responsible for
// updating the rocket inventory
type Service struct {

}

// New - returns a new instance of out rocket service 
func New() Service {
	return Service{}
}
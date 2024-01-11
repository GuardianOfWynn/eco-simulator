package pathfinding

import "github.com/GuardianOfWynn/eco-simulator/territory"

type Pathfinder struct {
	From       string
	Target     string
	Claim      territory.Claim
	RouteStyle territory.RouteStyle
}

func (p *Pathfinder) route() []string {
	return []string{}
}

package territory

type Pathfinder struct {
	From       *Territory
	Target     *Territory
	Claim      Claim
	RouteStyle RouteStyle
}

// for cheapest: order conns by tax value
func (p *Pathfinder) Route() []string {

	//implement djikstra algorithm
	return []string{}
}

package pathfinding

import (
	"strings"

	"github.com/GuardianOfWynn/eco-simulator/territory"
	"github.com/victorbetoni/go-streams/streams"
)

type Pathfinder struct {
	From       *territory.Territory
	Target     *territory.Territory
	Claim      territory.Claim
	RouteStyle territory.RouteStyle
}

func (p *Pathfinder) Route() []territory.Territory {

	routes := [][]territory.Territory{}

	for _, conn := range p.From.Connections {
		connTerr := p.Claim.GetTerritory(conn)
		pfinder := Pathfinder{
			From:       connTerr,
			Target:     p.Target,
			Claim:      p.Claim,
			RouteStyle: p.RouteStyle,
		}
		result, vis := pfinder.expand([]territory.Territory{})
		if result {
			routes = append(routes, vis)
		}
	}

	if len(routes) == 0 {
		return []territory.Territory{}
	}

	currentRoute := routes[0]

	for _, route := range routes {
		if len(route) < len(currentRoute) {
			currentRoute = route
		}
	}

	return currentRoute
}

func (p *Pathfinder) expand(visited []territory.Territory) (bool, []territory.Territory) {
	if p.From == nil {
		return false, visited
	}

	visited = append(visited, *p.From)
	for _, conn := range p.From.Connections {
		if strings.ToLower(conn) == strings.ToLower(p.Target.Name) {
			return true, visited
		}
	}
	for _, conn := range p.From.Connections {
		if streams.StreamOf[territory.Territory](visited...).AnyMatch(func(e territory.Territory) bool {
			return strings.ToLower(e.Name) == strings.ToLower(conn)
		}) {
			continue
		}
		connTerr := p.Claim.GetTerritory(conn)
		pfinder := Pathfinder{
			From:       connTerr,
			Target:     p.Target,
			Claim:      p.Claim,
			RouteStyle: p.RouteStyle,
		}
		result, vis := pfinder.expand(visited)
		if result {
			return result, vis
		}
	}

	visited = visited[:len(visited)-1]
	return false, []territory.Territory{}
}

package territory

import (
	"strings"

	"github.com/victorbetoni/go-streams/streams"
)

type Pathfinder struct {
	From       *Territory
	Target     *Territory
	Claim      Claim
	RouteStyle RouteStyle
}

// for cheapest: order conns by tax value
func (p *Pathfinder) Route() []Territory {

	routes := [][]Territory{}

	for _, conn := range p.From.Connections {
		connTerr := p.Claim.GetTerritory(conn)
		pfinder := Pathfinder{
			From:       connTerr,
			Target:     p.Target,
			Claim:      p.Claim,
			RouteStyle: p.RouteStyle,
		}
		result, vis := pfinder.expand([]Territory{})
		if result {
			routes = append(routes, vis)
		}
	}

	if len(routes) == 0 {
		return []Territory{}
	}

	currentRoute := routes[0]

	for _, route := range routes {
		if len(route) < len(currentRoute) {
			currentRoute = route
		}
	}

	return currentRoute
}

// for cheapest: order conns by tax value
func (p *Pathfinder) expand(visited []Territory) (bool, []Territory) {
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
		if streams.StreamOf[Territory](visited...).AnyMatch(func(e Territory) bool {
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
	return false, []Territory{}
}

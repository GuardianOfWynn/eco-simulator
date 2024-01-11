package territory

import (
	"strings"

	"github.com/GuardianOfWynn/eco-simulator/territory"
	"github.com/victorbetoni/go-streams/optionals"
	"github.com/victorbetoni/go-streams/streams"
)

type Claim struct {
	GlobalTax     uint8
	AllyTax       uint8
	GlobalStyle   territory.RouteStyle
	GlobalBorders territory.BorderStyle
	Territories   []territory.Territory
}

func (c *Claim) GetTerritory(name string) *optionals.Optional[Territory] {
	return optionals.Of[Territory](
		*streams.StreamOf[Territory](c.Territories).Filter(func(e Territory) bool {
			return strings.ToLower(e.Name) == strings.ToLower(name)
		}).FindFirst(),
	)
}

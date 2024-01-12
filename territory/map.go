package territory

import (
	"strings"

	"github.com/victorbetoni/go-streams/streams"
)

type GuildMap struct {
	Claims      []*Claim
	Territories []*Territory
}

func (g *GuildMap) TransferTerritory(territory *Territory, claim *Claim) {
	if territory.Claim != nil {
		index := streams.StreamOf(claim.Territories...).FindIndex(func(e *Territory) bool {
			return strings.EqualFold(e.Name, territory.Name)
		})
		currentClaim := territory.Claim
		if index != -1 {
			currentClaim.Territories = append(currentClaim.Territories[:index], currentClaim.Territories[index+1:]...)
		}
	}
	claim.Territories = append(claim.Territories, territory)
}

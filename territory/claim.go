package territory

import (
	"strings"

	"github.com/google/uuid"
)

type Claim struct {
	GlobalTax     uint8
	AllyTax       uint8
	GlobalStyle   RouteStyle
	GlobalBorders BorderStyle
	Territories   []*Territory
}

func (c *Claim) GetTerritory(name string) *Territory {
	for _, v := range c.Territories {
		if strings.EqualFold(v.Name, name) {
			return v
		}
	}
	return nil
}

func (c *Claim) SetAsHQ(territory *Territory) {
	for _, t := range c.Territories {
		t.HQ = false
	}
	territory.HQ = true
}

func (c *Claim) GetHQ() *Territory {
	for _, t := range c.Territories {
		if t.HQ {
			return t
		}
	}
	return nil
}

func (c *Claim) AskForResources(asking *Territory, res Storage) {
	hq := c.GetHQ()
	hq.PassingResource = append(hq.PassingResource, ResourceTransference{
		Id:        uuid.NewString(),
		Direction: HQ_TO_TERRITORY,
		Storage:   res,
		Target:    asking,
	})
}

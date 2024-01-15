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

type ClaimPreset struct {
	Name          string      `json:"name"`
	HQ            string      `json:"hq"`
	GlobalStyle   RouteStyle  `json:"global_style"`
	GlobalBorders BorderStyle `json:"global_borders"`
	GlobalTax     float64     `json:"global_tax"`
	AllyTax       float64     `json:"ally_tax"`
	Territories   []struct {
		Territory string `json:"territory"`
		Boosts    struct {
			Bonuses  map[string]uint8 `json:"bonuses"`
			Upgrades map[string]uint8 `json:"upgrades"`
		} `json:"boosts"`
		Style    RouteStyle  `json:"style"`
		Border   BorderStyle `json:"borders"`
		Treasury Treasury    `json:"treasury"`
		Tax      float64     `json:"tax"`
		AllyTax  float64     `json:"ally_tax"`
	} `json:"territories"`
}

func (c *ClaimPreset) Parse(guildMap *GuildMap) *Claim {
	claim := &Claim{
		GlobalTax:     uint8(c.GlobalTax),
		AllyTax:       uint8(c.AllyTax),
		GlobalStyle:   c.GlobalStyle,
		GlobalBorders: c.GlobalBorders,
		Territories:   []*Territory{},
	}

	guildMap.Claims = append(guildMap.Claims, claim)

	for _, v := range c.Territories {
		territory := guildMap.GetTerritory(v.Territory)
		territory.Reset()
		territory.Claim = claim
		if strings.EqualFold(v.Territory, c.HQ) {
			territory.HQ = true
		}
		for k, v := range v.Boosts.Bonuses {
			territory.Bonuses[k] = v
		}
		for k, v := range v.Boosts.Upgrades {
			territory.Upgrades[k] = v
		}
		territory.Borders = v.Border
		territory.Treasury = v.Treasury
		territory.Tax = v.Tax
		territory.AllyTax = v.AllyTax
		guildMap.TransferTerritory(territory, claim)
	}

	return claim
}

func (c *Claim) Tick() {
	for _, v := range c.Territories {
		v.Tick()
	}
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

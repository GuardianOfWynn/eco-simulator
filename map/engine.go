package territory

import (
	"encoding/json"
	"time"
)

var EngineInstance *Engine = nil

type Engine struct {
	Map    *GuildMap
	ticker *time.Ticker
}

func CreateEngine(territoriesJson []byte, presetJson []byte) {
	var terrs []BaseTerritory
	_ = json.Unmarshal(territoriesJson, &terrs)
	for _, v := range terrs {
		BaseTerritoriesMap[v.Name] = v
	}

	var preset ClaimPreset
	_ = json.Unmarshal(presetJson, &preset)

	territories := []*Territory{}
	for _, a := range terrs {
		territories = append(territories, a.CreateTerritoryInstance())
	}

	guildMap := &GuildMap{
		Territories: territories,
		Claims:      []*Claim{},
	}
	preset.Parse(guildMap)
	EngineInstance = &Engine{
		Map: guildMap,
	}
}

func (e *Engine) Start() {

	for range time.Tick(time.Second * 1) {
		for _, terr := range e.Map.Territories {
			terr.Tick()
		}
	}

}

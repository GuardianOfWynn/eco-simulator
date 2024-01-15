package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"

	territory "github.com/GuardianOfWynn/eco-simulator/map"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func main() {
	fmt.Println(basepath)
	b, err := ioutil.ReadFile(basepath + "\\territories.json")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)

	var terrs []territory.BaseTerritory
	_ = json.Unmarshal([]byte(str), &terrs)
	for _, v := range terrs {
		territory.BaseTerritoriesMap[v.Name] = v
	}

	presetJson, err := ioutil.ReadFile(basepath + "\\presets\\sky.json")
	if err != nil {
		fmt.Print(err)
	}
	presetStr := string(presetJson)

	var preset territory.ClaimPreset
	err = json.Unmarshal([]byte(presetStr), &preset)
	if err != nil {
		fmt.Println(err.Error())
	}

	territories := []*territory.Territory{}
	for _, a := range terrs {
		territories = append(territories, a.CreateTerritoryInstance())
	}

	guildMap := &territory.GuildMap{
		Territories: territories,
		Claims:      []*territory.Claim{},
	}
	preset.Parse(guildMap)
	territory.EngineInstance = &territory.Engine{
		Map: guildMap,
	}
	territory.EngineInstance.Start()
}

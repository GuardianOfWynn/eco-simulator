package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"

	"github.com/GuardianOfWynn/eco-simulator/engine"
	"github.com/GuardianOfWynn/eco-simulator/territory"
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

	territories := []*territory.Territory{}
	for _, a := range terrs {
		territories = append(territories, a.CreateTerritoryInstance())
	}

	guildMap := &territory.GuildMap{
		Territories: territories,
		Claims:      []*territory.Claim{},
	}

	engine := engine.Engine{
		Map: guildMap,
	}
	engine.Start()
}

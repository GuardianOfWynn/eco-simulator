package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"

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
	claim := territory.Claim{
		GlobalTax:     0,
		AllyTax:       0,
		GlobalStyle:   territory.CHEAPEST,
		GlobalBorders: territory.OPEN,
	}

	for _, a := range terrs {
		territories = append(territories, a.CreateTerritoryInstance())
	}
	claim.Territories = territories
	index := 32
	for i, v := range claim.Territories {
		if v.Name == "Canyon Entrance Waterfall" {
			index = i
		}
	}

	finder := territory.Pathfinder{
		From:       claim.Territories[0],
		Target:     claim.Territories[index],
		Claim:      claim,
		RouteStyle: territory.CHEAPEST,
	}

	fmt.Println("From: ", finder.From.Name)
	fmt.Println("Target: ", finder.Target.Name)
	fmt.Println("")
	route := finder.Route()
	fmt.Println(len(route))
	for _, v := range finder.Route() {
		fmt.Println(v)
	}
}

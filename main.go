package main

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	territory "github.com/GuardianOfWynn/eco-simulator/map"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

//export startEngineTinyGo
func StartEngine(territories, preset string) {
	territory.CreateEngine([]byte(territories), []byte(preset))
	territory.EngineInstance.Start()
}

func main() {
	fmt.Println("Started")
	log.Default().Println("Booting WASM EcoEngine application...")
}

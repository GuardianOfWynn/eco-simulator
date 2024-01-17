package wasm

import (
	"syscall/js"

	territory "github.com/GuardianOfWynn/eco-simulator/map"
	"github.com/norunners/vert"
)

func GetEngine(val js.Value, inputs []js.Value) interface{} {
	return vert.ValueOf(*territory.EngineInstance)
}

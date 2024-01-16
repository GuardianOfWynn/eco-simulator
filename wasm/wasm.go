package wasm

import (
	"syscall/js"

	territory "github.com/GuardianOfWynn/eco-simulator/map"
	"github.com/norunners/vert"
)

func StartEngine(val js.Value, inputs []js.Value) interface{} {
	territories := inputs[0].String()
	preset := inputs[1].String()

	territory.CreateEngine([]byte(territories), []byte(preset))
	territory.EngineInstance.Start()
	return nil
}

func GetEngine(val js.Value, inputs []js.Value) interface{} {
	return vert.ValueOf(*territory.EngineInstance)
}

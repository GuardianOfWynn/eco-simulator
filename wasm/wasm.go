package wasm

import (
	"syscall/js"

	territory "github.com/GuardianOfWynn/eco-simulator/map"
)

func StartEngine(val js.Value, inputs []js.Value) interface{} {
	territories := inputs[0].String()
	preset := inputs[1].String()

	territory.CreateEngine([]byte(territories), []byte(preset))
	territory.EngineInstance.Start()
	return nil
}

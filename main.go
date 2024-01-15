package main

import (
	"path/filepath"
	"runtime"
	"syscall/js"

	territory "github.com/GuardianOfWynn/eco-simulator/map"
	"github.com/GuardianOfWynn/eco-simulator/wasm"
	"github.com/norunners/vert"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func main() {
	js.Global().Set("startEngine", js.FuncOf(wasm.StartEngine))
	js.Global().Set("engineInstance", vert.ValueOf(territory.EngineInstance))
	ch := make(chan struct{}, 0)
	<-ch
}

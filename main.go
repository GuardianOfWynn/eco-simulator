package main

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"syscall/js"

	"github.com/GuardianOfWynn/eco-simulator/wasm"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func main() {
	fmt.Println("Started")
	log.Default().Println("Booting WASM EcoEngine application...")
	ch := make(chan byte, 1)
	js.Global().Set("startEngine", js.FuncOf(wasm.StartEngine))
	js.Global().Set("getEngineInstance", js.FuncOf(wasm.GetEngine))

	<-ch
}

package main

import (
	"fmt"
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
	fmt.Println("hello")
	ch := make(chan byte, 1)
	js.Global().Set("startEngine", js.FuncOf(wasm.StartEngine))

	<-ch
}

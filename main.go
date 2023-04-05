//go:build wasm
// +build wasm

package main

import (
	"fmt"
	"math"
	"syscall/js"
)

func calc(f float64) float64 {
	return math.Exp(f)
}

func main() {

	var ff js.Func
	ff = js.FuncOf(func(this js.Value, args []js.Value) any {
		fmt.Println("start go wasm task")
		f := args[0].Float()
		res := calc(f)
		fmt.Println("done go wasm task")
		return res
	})
	js.Global().Set("go_task_func", ff)
	fmt.Println("----------------")
	<-make(chan bool)

}

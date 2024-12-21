//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"

	"golang-echo-wasm-example/internal/wasm/add"
	"golang-echo-wasm-example/internal/wasm/divide"
	"golang-echo-wasm-example/internal/wasm/multiply"
	"golang-echo-wasm-example/internal/wasm/subtract"
)

func main() {
	js.Global().Set("add", js.FuncOf(add.Add))
	js.Global().Set("subtract", js.FuncOf(subtract.Subtract))
	js.Global().Set("multiply", js.FuncOf(multiply.Multiply))
	js.Global().Set("divide", js.FuncOf(divide.Divide))

	select {}
}

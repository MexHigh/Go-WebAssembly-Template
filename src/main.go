package main 
/* 
	if there is an error in the package statement above,
	add the following to your IDEs settings.json (for
	VSCode, this is '.vscode/settings.json'):

	"go.testEnvVars": {
		"GOARCH":"wasm",
		"GOOS":"js"
	},
	"go.toolsEnvVars": {
		"GOARCH":"wasm",
		"GOOS":"js"
	}
	
*/

import (
	"syscall/js"
)

// This is the go func signature, that is needed to be used in js
func add(this js.Value, i []js.Value) interface{} {
	return js.ValueOf(
		i[0].Int() + i[1].Int(),
	)
}

func registerJSFunctions() {
	js.Global().Set("add", js.FuncOf(add)) // adds js function add(a, b) --> a + b
}

func main() {
	// Add h1 to document
	document := js.Global().Get("document")
	p := document.Call("createElement", "h1")
	p.Set("innerHTML", "Go Wasm is working :)")
	document.Get("body").Call("appendChild", p)

	// Add go functions as js functions (can be called from
	// console, event listener (e.g. button click), etc.)
	registerJSFunctions()

	// block programm execution so that 
	// callbacks can still be executed
	c := make(chan bool, 0)
	<- c
}
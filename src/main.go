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

func add(this js.Value, i []js.Value) interface{} {
	return js.ValueOf(
		i[0].Int() + i[1].Int(),
	)
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add)) // Call: add(2, 3); Returns: 5 (of course)
}

func main() {
	// Add h1 to document
	document := js.Global().Get("document")
	p := document.Call("createElement", "h1")
	p.Set("innerHTML", "Go Wasm is working :)")
	document.Get("body").Call("appendChild", p)

	// Add go functions as a js functions (can be called 
	//from console or from an event listener (e.g. button click))
	registerCallbacks()

	c := make(chan bool, 0)
	<- c
}
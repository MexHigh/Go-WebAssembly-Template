# How to use this?

## Short version
- `./build_wasm.sh`
- `go run devserver.go`
- Visit the webpage

## Long version

The Go code resides in the `src` folder. `./build_wasm.sh` complies the `main.wasm` file and puts it in the `webroot` directory. 

With `go run devserver.go` you can start a simple webserver, that serves all contents within the `webroot` directory. A special `.js` file instantiates the WebAssembly file and executes its code.

## Important !
As `.wasm` files need to bundle the whole execution runtime, the file can be fairly large (currently around 1.3M). WebAssembly therefore may not be performant enough right now for smaller projects. But there is an option right now to use a smaller go runtime (https://tinygo.org/webassembly/webassembly/).
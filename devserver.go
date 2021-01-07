package main

import (
	"fmt"
	"net/http"
)

const (
	listen = "0.0.0.0:3001"
)

func main() {
	fmt.Println("Listening on", listen)
	http.ListenAndServe(
		listen, 
		http.FileServer(http.Dir("./webroot")),
	)
}
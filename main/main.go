package main

import (
	"net/http"

	"github.com/nelthaarion/artesia"
)

func main() {
	hg := artesia.NewHandlerGenerator()
	mux := http.NewServeMux()
	newMux := hg(mux)

	http.ListenAndServe(":3000", newMux)

}

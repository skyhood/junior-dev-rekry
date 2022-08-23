// To run this, simply run `go run go-server.go` and visit localhost:3333/api?name=Jerry

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func respond(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	res := map[string]interface{}{
		"hello": name,
	}

	respond(w, r, res)
}

func main() {
	http.HandleFunc("/api", apiHandler)

	addr := ":3333"
	fmt.Printf("Running server on %s\n", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}

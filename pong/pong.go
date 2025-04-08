package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/pong", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Ping")
	})
	http.ListenAndServe(":1235", nil)
}

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	addr := ":8080"

	mux := http.NewServeMux()

	mux.HandleFunc("/Hi", func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		resp := Resp{
			Language:    "Eng",
			Translation: "Hi",
		}
		if err := enc.Encode(resp); err != nil {
			log.Panicf("unable to encode response %s", err)
		}
	})
	log.Printf("listening on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

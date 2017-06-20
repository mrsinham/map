package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/circle", http.HandlerFunc(circle))
	err := http.ListenAndServe(":2003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func circle(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "image/svg+xml")

	m := NewMap(686876788, 400, 400, 5000)

	s := &samplingUniform{}
	s.GetPoints(m)

	renderGrid(w, m)
}

package main

import (
	"fmt"
	"net/http"
)

var PORT = ":3000"

func doStuff(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Waffles are great")
}

func main() {
	http.HandleFunc("/", doStuff)
	fmt.Printf("Listening on port %s\n", PORT)
	http.ListenAndServe(PORT, nil)
}

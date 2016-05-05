package main

import (
	"log"
	"net/http"
)

func main() {
	var dir = http.Dir("site")
  http.Handle("/", http.FileServer(dir))
  log.Fatal(http.ListenAndServe(":3000", nil))
}

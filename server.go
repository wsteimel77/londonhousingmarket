package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port int

func init() {
	flag.IntVar(&port, "p", 3000, "Port")
	flag.Parse()
}

func main() {
	var dir = http.Dir("site")
	http.Handle("/", http.FileServer(dir))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

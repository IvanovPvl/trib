package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "3333", "port for listening")
	flag.Parse()
	log.Println("Service running on port: " + *port)
	log.Fatal(http.ListenAndServe(":"+*port, tribHandler()))
}

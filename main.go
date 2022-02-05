package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	flag.Parse()
	flag.Parse()
	resourceConsumerHandler := NewResourceConsumerHandler()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *Port), resourceConsumerHandler))
}

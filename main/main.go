package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

func main() {
	initServices()
	port := fmt.Sprintf(":%d", 8081)
	http.HandleFunc("/api", graphqlHandler)

	log.Println("Listening on port " + port)
	err := http.ListenAndServe(port, nil)
	err = errors.Wrap(err, "Error creating server")
	if err != nil {
		log.Println(err)
		return
	}
}

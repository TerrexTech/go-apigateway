package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

// graphqlHandler handles GraphQL requests.
func graphqlHandler(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
		)
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}

	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = errors.Wrap(err, "Error reading request body")
		log.Println(err)
		return
	}

	result := graphql.Do(graphql.Params{
		Schema:        Schema,
		RequestString: string(reqBytes),
		RootObject:    rootObject,
	})
	if len(result.Errors) > 0 {
		log.Printf("GQL Error: %v", result.Errors)
	}

	json.NewEncoder(w).Encode(result)
}

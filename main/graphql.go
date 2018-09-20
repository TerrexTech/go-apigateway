package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

func graphqlHandler(w http.ResponseWriter, r *http.Request) {
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
		log.Printf("gql error: %v", result.Errors)
	}

	json.NewEncoder(w).Encode(result)
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	graphql "github.com/Davd96/graphql-go"
	"github.com/Davd96/graphql-go-todo/configs"
	"github.com/Davd96/graphql-go-todo/graphql/resolver"
	"github.com/Davd96/graphql-go/relay"
)

func main() {

	if os.Getenv(configs.ENVIROMENT_VARIABLE_PSQL) == "" {
		panic(fmt.Sprintf("environment variable not found (%s)", configs.ENVIROMENT_VARIABLE_PSQL))
	}

	var schema string = schemaToString("./graphql/schema/schema.graphql")

	schemaParsed, err := graphql.ParseSchema(schema, &resolver.QueryResolver{}, &resolver.MutationResolver{})

	if err != nil {
		panic(err)
	}

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		http.ServeFile(w, r, "web/graphql.html")

	}))

	http.Handle("/query", &relay.Handler{Schema: schemaParsed})

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func schemaToString(schemaPath string) string {
	byteSchema, err := ioutil.ReadFile(schemaPath)
	if err != nil {
		log.Fatal(err)
	}

	return string(byteSchema)
}

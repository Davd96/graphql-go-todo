package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Davd96/graphql-go-todo/configs"
	"github.com/Davd96/graphql-go-todo/graphql/resolver"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {

	if os.Getenv(configs.ENVIROMENT_VARIABLE_PSQL) == "" {
		panic(fmt.Sprintf("environment variable not found (%s)", configs.ENVIROMENT_VARIABLE_PSQL))
	}

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		http.ServeFile(w, r, "web/graphql.html")

	}))

	schema := schemaToString("./graphql/schema/schema.graphql")

	schemaParsed := graphql.MustParseSchema(schema, &resolver.QueryResolver{})

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

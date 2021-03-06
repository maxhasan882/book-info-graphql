package main

import (
	"bookinfo/cmd/app/config"
	"bookinfo/graph"
	"bookinfo/graph/generated"
	db "bookinfo/internal/app/adapter/db/connections"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	initEnv()
}

//initEnv Load and initialize environment Also run migration is any available
func initEnv() {
	log.Printf("Loading environment settings.")

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("No local env file. Using global OS environment variables")
	}
	config.SetEnvironment()
	log.Printf("Connecting to database")
	con, err := db.GetDatabase()
	if err != nil {
		panic(err)
	}
	db.RunMigrations(con)
}

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	http.HandleFunc("/load-data", loadData)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func loadData(w http.ResponseWriter, _ *http.Request) {
	db.DataLoader()
	_, err := fmt.Fprintf(w, "hello")
	if err != nil {
		return
	}
}

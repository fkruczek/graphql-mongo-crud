package cmd

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/fk/gqlplayground/auth"
	"github.com/fk/gqlplayground/db"
	"github.com/fk/gqlplayground/graph"
	"github.com/fk/gqlplayground/graph/generated"
	"github.com/go-chi/chi"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serving gql server",
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := db.New()
		if err != nil {
			return err
		}
		defer db.Disconnect()

		router := chi.NewRouter()
		router.Use(auth.Middleware(db))
		srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Db: db}}))

		router.Handle("/", playground.Handler("GraphQL playground", "/query"))
		router.Handle("/query", srv)

		log.Printf("http://localhost:8080")
		err = http.ListenAndServe(":8080", router)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func Run() {
	serveCmd.Execute()
}

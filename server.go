package main

import (
	"context"
	"entgo.io/contrib/entgql"
	"fmt"
	"github.com/maxh/gqlgen-todos/orm/ent"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/maxh/gqlgen-todos/graphql"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/maxh/gqlgen-todos/orm/ent/runtime"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	client, err := ent.Open(
		"sqlite3",
		"file:ent?mode=memory&cache=shared&_fk=1",
	)
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal("running schema migration", err)
	}
	srv := handler.NewDefaultServer(graphql.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})

	_, err = CreateUser(context.TODO(), client)
	if err != nil {
		log.Fatal("unable to create user", err)
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	t, err := client.Tenant.Create().SetName("my tenant").Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating Tenant: %w", err)
	}
	log.Println("tenant was created: ", t)

	o, err := client.Organization.Create().SetName("my organization").SetTenant(t).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating organization: %w", err)
	}
	log.Println("organization was created: ", o)

	u, err := client.User.
		Create().
		SetOrganization(o).
		SetTenant(t).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	log.Println("user was created: ", u)
	return u, nil
}

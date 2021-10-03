package main

import (
	"context"
	"entgo.io/contrib/entgql"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/maxh/gqlgen-todos/auth"
	"github.com/maxh/gqlgen-todos/graphql"
	"github.com/maxh/gqlgen-todos/orm/ent"
	"github.com/maxh/gqlgen-todos/orm/ent/migrate"
	_ "github.com/maxh/gqlgen-todos/orm/ent/runtime"
	"github.com/maxh/gqlgen-todos/orm/revision"
	"github.com/maxh/gqlgen-todos/viewer"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	client, err := ent.Open(
		"postgres",
		"host=localhost port=5432 user=postgres dbname=gqlgen_todos_dev password=postgres sslmode=disable")
	//client, err := ent.Open(
	//	"sqlite3",
	//	"file:ent?mode=memory&cache=shared&_fk=1",
	//)
	if err != nil {
		log.Fatal("opening ent client", err)
	}

	client.Use(func(next ent.Mutator) ent.Mutator {
		return revision.AddRevision(next)
	})

	if err := client.Schema.Create(context.Background(), migrate.WithDropIndex(true),
		migrate.WithDropColumn(true)); err != nil {
		log.Fatal("running schema migration", err)
	}

	router := chi.NewRouter()

	router.Use(auth.Middleware(client))

	srv := handler.NewDefaultServer(graphql.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})

	err = wrapSeedDatabase(context.TODO(), client)
	if err != nil {
		log.Printf("error seeding db", err)
	}

	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func wrapSeedDatabase(ctx context.Context, client *ent.Client) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	txClient := tx.Client()
	txContext := ent.NewContext(ctx, txClient)
	if err := seedDatabase(txContext, txClient); err != nil {
		return rollback(tx, err)
	}
	return tx.Commit()
}

// rollback calls to tx.Rollback and wraps the given error with the rollback error if occurred.
func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}

func seedDatabase(ctx context.Context, client *ent.Client) error {
	ctx = viewer.NewContext(ctx, viewer.UserViewer{
		Role: viewer.Admin,
	})
	t, err := client.Tenant.Create().SetName("my tenant").Save(ctx)
	if err != nil {
		return fmt.Errorf("failed creating Tenant: %w", err)
	}

	o, err := client.Organization.Create().SetName("my organization").SetTenant(t).Save(ctx)
	if err != nil {
		return fmt.Errorf("failed creating organization: %w", err)
	}

	_, err = client.User.
		Create().
		SetOrganization(o).
		SetTenant(t).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed creating user: %w", err)
	}

	return nil
}

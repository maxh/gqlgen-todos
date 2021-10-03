package main

import (
	"context"
	"entgo.io/contrib/entgql"
	"fmt"
	"github.com/maxh/gqlgen-todos/auth"
	"github.com/maxh/gqlgen-todos/orm/ent"
	"github.com/maxh/gqlgen-todos/orm/ent/migrate"
	"github.com/maxh/gqlgen-todos/util"
	"github.com/maxh/gqlgen-todos/viewer"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/maxh/gqlgen-todos/graphql"

	_ "github.com/lib/pq"
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
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if m.Type() == ent.TypeEntityRevision {
				return next.Mutate(ctx, m)
			}
			start := time.Now()
			defer func() {
				log.Printf("Op=%s\tType=%s\tTime=%s\tConcreteType=%T\n", m.Op(), m.Type(), time.Since(start), m)
			}()

			er, err := client.EntityRevision.Create().
				SetEntityID("123").
				SetEntityRevision("456").
				SetEntityValue(&util.Any).
				Save(ctx)
			if err != nil {
				log.Printf("error creating rev", err)
			}
			log.Printf("created er", er)

			return next.Mutate(ctx, m)
		})
	})

	if err := client.Schema.Create(context.Background(), migrate.WithDropIndex(true),
		migrate.WithDropColumn(true)); err != nil {
		log.Fatal("running schema migration", err)
	}

	router := chi.NewRouter()

	router.Use(auth.Middleware(client))

	srv := handler.NewDefaultServer(graphql.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})

	_, err = CreateUser(context.TODO(), client)
	if err != nil {
		log.Fatal("unable to create user", err)
	}

	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	ctx = viewer.NewContext(ctx, viewer.UserViewer{
		Role: viewer.Admin,
	})
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

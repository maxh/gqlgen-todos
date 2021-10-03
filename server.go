package main

import (
	"context"
	"entgo.io/contrib/entgql"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/maxh/gqlgen-todos/auth"
	"github.com/maxh/gqlgen-todos/graphql"
	"github.com/maxh/gqlgen-todos/orm/ent"
	"github.com/maxh/gqlgen-todos/orm/ent/migrate"
	"github.com/maxh/gqlgen-todos/qid"
	"github.com/maxh/gqlgen-todos/util"
	"github.com/maxh/gqlgen-todos/viewer"
	"log"
	"net/http"
	"os"
	"strconv"

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
		return addRevision(next)
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

type EntityRevisioner interface {
	ID() (value qid.ID, exists bool)
}

func addRevision(next ent.Mutator) ent.Mutator {
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		if m.Type() == ent.TypeEntityRevision {
			// We don't want to store revisions for records for revisions themselves,
			// otherwise we'll end up in an infinite loop.
			return next.Mutate(ctx, m)
		}

		client := ent.FromContext(ctx)

		// Apply the mutation _before_ saving the audit.
		// (Other hooks may change the node before persistence, and we only want to save
		// the "final" revision from this transaction in the revisions table.)
		v, err := next.Mutate(ctx, m)
		if err != nil {
			// Don't save a revision if the mutation failed.
			return v, err
		}

		id := entityId(m)
		node := mutatedNode(ctx, m, client)
		value := entityValue(node)
		_, err = client.EntityRevision.Create().
			SetEntityID(string(id)).
			SetEntityRevision("456").
			SetEntityValue(&value).
			Save(ctx)
		if err != nil {
			return nil, err
		}

		return v, nil
	})
}

func entityValue(node *ent.Node) util.EntityValue {
	fieldMap := util.FieldMap{}
	for _, f := range node.Fields {
		st, err := strconv.Unquote(f.Value)
		if err != nil {
			// Booleans cannot be unquoted; it's not a problem
			// to fallback on the raw value.
			st = f.Value
		}
		fieldMap[f.Name] = st
	}
	edgeMap := util.EdgeMap{}
	for _, e := range node.Edges {
		edgeMap[e.Name] = e.IDs
	}
	value := util.EntityValue{
		Fields: fieldMap,
		Edges:  edgeMap,
	}
	return value
}

func mutatedNode(ctx context.Context, m ent.Mutation, client *ent.Client) *ent.Node {
	// All saved entities must have an ID.
	id := entityId(m)

	// We need to look up the table name because WithFixedNodeType expects it.
	tableName := ent.TablesByEntType[m.Type()]
	noder, err := client.Noder(ctx, id, ent.WithFixedNodeType(tableName))

	// At this point, we expect the node to exist in the transaction context.
	if err != nil {
		panic(err)
	}
	node, err := noder.Node(ctx)
	if err != nil {
		panic(err)
	}

	return node
}

func entityId(m ent.Mutation) qid.ID {
	rev, ok := m.(EntityRevisioner)
	if !ok {
		panic("no id method on mutated node")
	}
	id, exists := rev.ID()
	if !exists {
		panic("id does not exist on mutated node")
	}
	return id
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

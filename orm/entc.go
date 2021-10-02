// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension()
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	extensions := entc.Extensions(ex)
	config := &gen.Config{
		// IDType: &field.TypeInfo{Type: field.TypeString},
		Target: "./ent",
		Package: "github.com/maxh/gqlgen-todos/orm/ent",
	}
	if err := entc.Generate("./schema", config, extensions); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
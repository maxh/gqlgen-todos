//go:build ignore
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
	options := []entc.Option{
		entc.FeatureNames("privacy"),
		entc.Extensions(ex),
	}
	config := &gen.Config{
		// IDType: &field.TypeInfo{Type: field.TypeString},
		Target:  "./ent",
		Package: "github.com/maxh/gqlgen-todos/orm/ent",
	}
	err = entc.Generate("./schema", config, options...)
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}

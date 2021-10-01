// +build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/contrib/entgql"
)

func main() {
	//ex, err := entgql.NewExtension()
	//if err != nil {
	//	log.Fatalf("creating entgql extension: %v", err)
	//}
	c := &gen.Config{
		Target: "./orm/ent",
	}
	if err := entc.Generate("./schema", c); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
package main

import (
	"context"
	"log"
	"os"

	"ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"

	_ "github.com/lib/pq"
)

func main() {
	// We need a name for the new migration file.
	if len(os.Args) < 2 {
		log.Fatalln("no name given")
	}
	// Create a local migration directory.
	dir, err := migrate.NewLocalDir("migrations")
	if err != nil {
		log.Fatalln(err)
	}
	// Load the graph.
	graph, err := entc.LoadGraph("./ent/schema", &gen.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	tbls, err := graph.Tables()
	if err != nil {
		log.Fatalln(err)
	}
	// Open connection to the database.
	drv, err := sql.Open("postgres", "user=postgres password=nopass host=localhost port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	// Inspect the current database state and compare it with the graph.
	m, err := schema.NewMigrate(drv, schema.WithDir(dir))
	if err != nil {
		log.Fatalln(err)
	}
	if err := m.NamedDiff(context.Background(), os.Args[1], tbls...); err != nil {
		log.Fatalln(err)
	}
}

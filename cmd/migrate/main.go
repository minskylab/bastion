package main

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	// postgresql driver
	// _ "github.com/lib/pq"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	// client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=postgrespassword sslmode=disable")
	// if err != nil {
	// 	log.Fatalf("failed connecting to mysql: %v", err)
	// }
	// defer client.Close()

	// ctx := context.Background()

	// if err = client.Schema.Create(
	// 	ctx,
	// 	migrate.WithGlobalUniqueID(true),
	// 	migrate.WithDropIndex(true),
	// 	migrate.WithDropColumn(true),
	// ); err != nil {
	// 	log.Fatalf("failed creating schema resources: %v", err)
	// }

	dbURL := "postgres://postgres:postgrespassword@localhost:5432/postgres?sslmode=disable"

	client, db := Open(dbURL, false)

	if err := AutoMigration(context.Background(), client, db); err != nil {
		logrus.Fatalf("failed creating schema resources: %v", err)
	}

	fmt.Println("schema resources created")
}

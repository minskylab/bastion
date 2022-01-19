package main

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	// postgresql driver
	// _ "github.com/lib/pq"
	"github.com/gookit/config/v2"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	godotenv.Load()
	config.WithOptions(config.ParseEnv)

	dbURL := config.Getenv("DATABASE_URL")

	client, db := Open(dbURL, false)

	if err := AutoMigration(context.Background(), client, db); err != nil {
		logrus.Fatalf("failed creating schema resources: %v", err)
	}

	fmt.Println("schema resources created")
}

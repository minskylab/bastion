package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/minskylab/bastion/ent"
	"github.com/pkg/errors"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func hasUpdateAtField(cols []*schema.Column) bool {
	for _, c := range cols {
		if c.Name == "updated_at" || c.Name == "updatedAt" {
			return true
		}
	}
	return false
}

const pgTimestampUpdateTrigger = `
BEGIN;

DROP TRIGGER IF EXISTS set_public_%[1]s_updated_at ON %[1]s;

CREATE TRIGGER set_public_%[1]s_updated_at
    BEFORE UPDATE 
    ON %[1]s
    FOR EACH ROW
    EXECUTE PROCEDURE set_current_timestamp_updated_at();
    
COMMIT;
`

const pgSetTimestampFunction = `
CREATE OR REPLACE FUNCTION set_current_timestamp_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
`

func performPrelude(db *sql.DB) error {
	if _, err := db.Exec(pgSetTimestampFunction); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func Open(databaseURL string, debug bool) (*ent.Client, *sql.DB) {
	tries := 5

	var db *sql.DB
	var err error

	for i := 0; i < tries; i++ {
		db, err = sql.Open("pgx", databaseURL)
		if err != nil {
			log.Println("[ERROR]", err.Error())
		}
		time.Sleep(6 * time.Second)
	}

	if db == nil {
		log.Fatal(err)
	}

	drv := entsql.OpenDB(dialect.Postgres, db)

	client := ent.NewClient(ent.Driver(drv))

	if debug {
		client = client.Debug()
	}

	return client, db
}

func AutoMigration(ctx context.Context, client *ent.Client, db *sql.DB) error {
	if err := performPrelude(db); err != nil {
		return errors.WithStack(err)
	}

	if err := client.Schema.Create(
		ctx,
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
		// schema.WithGlobalUniqueID(true),
		// schema.WithFixture(true),
		schema.WithHooks(func(next schema.Creator) schema.Creator {
			return schema.CreateFunc(func(ctx context.Context, tables ...*schema.Table) error {
				if err := next.Create(ctx, tables...); err != nil {
					return errors.WithStack(err)
				}

				// after create tables, create updated_at triggers
				for _, t := range tables {
					if !hasUpdateAtField(t.Columns) {
						continue
					}

					qFunc := fmt.Sprintf(pgTimestampUpdateTrigger, t.Name)

					if _, err := db.Exec(qFunc); err != nil {
						log.Printf("error at create update_at timestamp trigger: %s\n", err.Error())
						continue
					}
				}

				return nil
			})
		}),
	); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

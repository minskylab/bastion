package core

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	// _ "github.com/jackc/pgx/v4/stdlib"
)

var IDField = field.UUID("id", uuid.UUID{}).
	Default(uuid.New).
	Immutable().
	Annotations(entsql.Annotation{
		Default: "gen_random_uuid()",
	})

var CreatedAtField = field.Time("createdAt").
	Default(time.Now).
	Annotations(entsql.Annotation{
		Default: "now()",
	})

var UpdatedAtField = field.Time("updatedAt").
	Default(time.Now).
	UpdateDefault(time.Now).
	Annotations(entsql.Annotation{
		Default: "now()",
	})

var SystemFields = []ent.Field{
	IDField,
	CreatedAtField,
	UpdatedAtField,
}

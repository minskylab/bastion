package core

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

var SystemFields = []ent.Field{
	field.UUID("id", uuid.UUID{}),
	field.Time("createdAt").Default(time.Now),
	field.Time("updatedAt").Default(time.Now),
}

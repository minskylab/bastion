package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/minskylab/bastion/core"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return append(core.SystemFields, []ent.Field{
		field.String("name"),
		field.String("description"),
		field.Time("startAt").Default(time.Now),
	}...)
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", Role.Type),
		edge.To("tasks", Task.Type),
		edge.From("organizations", Organization.Type).Ref("projects"),
	}
}

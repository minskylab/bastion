package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/minskylab/bastion/core"
)

// Organization holds the schema definition for the Organization entity.
type Organization struct {
	ent.Schema
}

// Fields of the Organization.
func (Organization) Fields() []ent.Field {
	return []ent.Field{
		core.IDField,
		core.CreatedAtField,
		core.UpdatedAtField,
		field.String("name"),
	}
}

// Edges of the Organization.
func (Organization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("projects", Project.Type),
		edge.To("developers", Member.Type),
		edge.To("managers", Member.Type),
	}
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/minskylab/bastion/core"
)

// Member holds the schema definition for the Member entity.
type Member struct {
	ent.Schema
}

// Fields of the Member.
func (Member) Fields() []ent.Field {
	return append(core.SystemFields, []ent.Field{
		field.String("name"),
		field.String("email"),
	}...)
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", Role.Type),
		edge.From("developerOf", Organization.Type).Ref("developers"),
		edge.From("managerOf", Organization.Type).Ref("managers"),
		edge.From("tasks", Task.Type).Ref("assignees"),
	}
}

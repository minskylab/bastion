package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/minskylab/bastion/core"
	hasura "github.com/minskylab/ent-hasura"
)

// Organization holds the schema definition for the Organization entity.
type Organization struct {
	ent.Schema
}

// Fields of the Organization.
func (Organization) Fields() []ent.Field {
	return append(core.SystemFields, []ent.Field{
		field.String("name"),
	}...)
}

// Edges of the Organization.
func (Organization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("projects", Project.Type),
		edge.To("members", Member.Type),
	}
}

func (Organization) Annotations() []schema.Annotation {
	return []schema.Annotation{
		core.MemberPermissions(
			&hasura.PermissionSelect{
				AllColumns:        true,
				Filter:            hasura.M{},
				AllowAggregations: true,
			},
			nil,
			nil,
			nil,
		),
	}
}

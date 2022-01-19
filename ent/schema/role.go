package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/minskylab/bastion/core"
	hasura "github.com/minskylab/ent-hasura"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return append(core.SystemFields, []ent.Field{
		field.String("name"),
	}...)
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("members", Member.Type).Ref("roles"),
		edge.From("projects", Project.Type).Ref("roles"),
	}
}

func (Role) Annotations() []schema.Annotation {
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

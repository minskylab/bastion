package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/minskylab/bastion/core"
	hasura "github.com/minskylab/ent-hasura"
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
		field.UUID("authID", uuid.UUID{}),
	}...)
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", Role.Type),
		edge.From("organizations", Organization.Type).Ref("members"),
		edge.From("tasks", Task.Type).Ref("assignees"),
	}
}

func (Member) Annotations() []schema.Annotation {
	isAuthUser := hasura.M{"auth_id": hasura.Eq("X-Hasura-User-Id")}

	return []schema.Annotation{
		core.MemberPermissions(
			&hasura.PermissionSelect{
				AllColumns:        true,
				Filter:            isAuthUser,
				AllowAggregations: true,
			},
			nil,
			&hasura.PermissionUpdate{
				AllColumns: true,
				Filter:     isAuthUser,
				Check:      isAuthUser,
			},
			nil,
		),
	}
}

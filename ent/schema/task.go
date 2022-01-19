package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/minskylab/bastion/core"
	hasura "github.com/minskylab/ent-hasura"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return append(core.SystemFields, []ent.Field{
		field.String("name"),
		field.String("url"),
	}...)
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).Ref("tasks"),
		edge.To("assignees", Member.Type),
	}
}

func (Task) Annotations() []schema.Annotation {
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

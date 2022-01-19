// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/minskylab/bastion/ent/member"
	"github.com/minskylab/bastion/ent/organization"
	"github.com/minskylab/bastion/ent/predicate"
	"github.com/minskylab/bastion/ent/role"
	"github.com/minskylab/bastion/ent/task"
)

// MemberUpdate is the builder for updating Member entities.
type MemberUpdate struct {
	config
	hooks    []Hook
	mutation *MemberMutation
}

// Where appends a list predicates to the MemberUpdate builder.
func (mu *MemberUpdate) Where(ps ...predicate.Member) *MemberUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetCreatedAt sets the "createdAt" field.
func (mu *MemberUpdate) SetCreatedAt(t time.Time) *MemberUpdate {
	mu.mutation.SetCreatedAt(t)
	return mu
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableCreatedAt(t *time.Time) *MemberUpdate {
	if t != nil {
		mu.SetCreatedAt(*t)
	}
	return mu
}

// SetUpdatedAt sets the "updatedAt" field.
func (mu *MemberUpdate) SetUpdatedAt(t time.Time) *MemberUpdate {
	mu.mutation.SetUpdatedAt(t)
	return mu
}

// SetName sets the "name" field.
func (mu *MemberUpdate) SetName(s string) *MemberUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetEmail sets the "email" field.
func (mu *MemberUpdate) SetEmail(s string) *MemberUpdate {
	mu.mutation.SetEmail(s)
	return mu
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (mu *MemberUpdate) AddRoleIDs(ids ...uuid.UUID) *MemberUpdate {
	mu.mutation.AddRoleIDs(ids...)
	return mu
}

// AddRoles adds the "roles" edges to the Role entity.
func (mu *MemberUpdate) AddRoles(r ...*Role) *MemberUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return mu.AddRoleIDs(ids...)
}

// AddDeveloperOfIDs adds the "developerOf" edge to the Organization entity by IDs.
func (mu *MemberUpdate) AddDeveloperOfIDs(ids ...uuid.UUID) *MemberUpdate {
	mu.mutation.AddDeveloperOfIDs(ids...)
	return mu
}

// AddDeveloperOf adds the "developerOf" edges to the Organization entity.
func (mu *MemberUpdate) AddDeveloperOf(o ...*Organization) *MemberUpdate {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return mu.AddDeveloperOfIDs(ids...)
}

// AddManagerOfIDs adds the "managerOf" edge to the Organization entity by IDs.
func (mu *MemberUpdate) AddManagerOfIDs(ids ...uuid.UUID) *MemberUpdate {
	mu.mutation.AddManagerOfIDs(ids...)
	return mu
}

// AddManagerOf adds the "managerOf" edges to the Organization entity.
func (mu *MemberUpdate) AddManagerOf(o ...*Organization) *MemberUpdate {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return mu.AddManagerOfIDs(ids...)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (mu *MemberUpdate) AddTaskIDs(ids ...uuid.UUID) *MemberUpdate {
	mu.mutation.AddTaskIDs(ids...)
	return mu
}

// AddTasks adds the "tasks" edges to the Task entity.
func (mu *MemberUpdate) AddTasks(t ...*Task) *MemberUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return mu.AddTaskIDs(ids...)
}

// Mutation returns the MemberMutation object of the builder.
func (mu *MemberUpdate) Mutation() *MemberMutation {
	return mu.mutation
}

// ClearRoles clears all "roles" edges to the Role entity.
func (mu *MemberUpdate) ClearRoles() *MemberUpdate {
	mu.mutation.ClearRoles()
	return mu
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (mu *MemberUpdate) RemoveRoleIDs(ids ...uuid.UUID) *MemberUpdate {
	mu.mutation.RemoveRoleIDs(ids...)
	return mu
}

// RemoveRoles removes "roles" edges to Role entities.
func (mu *MemberUpdate) RemoveRoles(r ...*Role) *MemberUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return mu.RemoveRoleIDs(ids...)
}

// ClearDeveloperOf clears all "developerOf" edges to the Organization entity.
func (mu *MemberUpdate) ClearDeveloperOf() *MemberUpdate {
	mu.mutation.ClearDeveloperOf()
	return mu
}

// RemoveDeveloperOfIDs removes the "developerOf" edge to Organization entities by IDs.
func (mu *MemberUpdate) RemoveDeveloperOfIDs(ids ...uuid.UUID) *MemberUpdate {
	mu.mutation.RemoveDeveloperOfIDs(ids...)
	return mu
}

// RemoveDeveloperOf removes "developerOf" edges to Organization entities.
func (mu *MemberUpdate) RemoveDeveloperOf(o ...*Organization) *MemberUpdate {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return mu.RemoveDeveloperOfIDs(ids...)
}

// ClearManagerOf clears all "managerOf" edges to the Organization entity.
func (mu *MemberUpdate) ClearManagerOf() *MemberUpdate {
	mu.mutation.ClearManagerOf()
	return mu
}

// RemoveManagerOfIDs removes the "managerOf" edge to Organization entities by IDs.
func (mu *MemberUpdate) RemoveManagerOfIDs(ids ...uuid.UUID) *MemberUpdate {
	mu.mutation.RemoveManagerOfIDs(ids...)
	return mu
}

// RemoveManagerOf removes "managerOf" edges to Organization entities.
func (mu *MemberUpdate) RemoveManagerOf(o ...*Organization) *MemberUpdate {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return mu.RemoveManagerOfIDs(ids...)
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (mu *MemberUpdate) ClearTasks() *MemberUpdate {
	mu.mutation.ClearTasks()
	return mu
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (mu *MemberUpdate) RemoveTaskIDs(ids ...uuid.UUID) *MemberUpdate {
	mu.mutation.RemoveTaskIDs(ids...)
	return mu
}

// RemoveTasks removes "tasks" edges to Task entities.
func (mu *MemberUpdate) RemoveTasks(t ...*Task) *MemberUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return mu.RemoveTaskIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MemberUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	mu.defaults()
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MemberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MemberUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MemberUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MemberUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MemberUpdate) defaults() {
	if _, ok := mu.mutation.UpdatedAt(); !ok {
		v := member.UpdateDefaultUpdatedAt()
		mu.mutation.SetUpdatedAt(v)
	}
}

func (mu *MemberUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   member.Table,
			Columns: member.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: member.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: member.FieldCreatedAt,
		})
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: member.FieldUpdatedAt,
		})
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldName,
		})
	}
	if value, ok := mu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldEmail,
		})
	}
	if mu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   member.RolesTable,
			Columns: member.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: role.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedRolesIDs(); len(nodes) > 0 && !mu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   member.RolesTable,
			Columns: member.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   member.RolesTable,
			Columns: member.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.DeveloperOfCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   member.DeveloperOfTable,
			Columns: member.DeveloperOfPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: organization.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedDeveloperOfIDs(); len(nodes) > 0 && !mu.mutation.DeveloperOfCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   member.DeveloperOfTable,
			Columns: member.DeveloperOfPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: organization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.DeveloperOfIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   member.DeveloperOfTable,
			Columns: member.DeveloperOfPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: organization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.ManagerOfCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   member.ManagerOfTable,
			Columns: member.ManagerOfPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: organization.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedManagerOfIDs(); len(nodes) > 0 && !mu.mutation.ManagerOfCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   member.ManagerOfTable,
			Columns: member.ManagerOfPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: organization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ManagerOfIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   member.ManagerOfTable,
			Columns: member.ManagerOfPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: organization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   member.TasksTable,
			Columns: member.TasksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: task.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedTasksIDs(); len(nodes) > 0 && !mu.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   member.TasksTable,
			Columns: member.TasksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   member.TasksTable,
			Columns: member.TasksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{member.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// MemberUpdateOne is the builder for updating a single Member entity.
type MemberUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MemberMutation
}

// SetCreatedAt sets the "createdAt" field.
func (muo *MemberUpdateOne) SetCreatedAt(t time.Time) *MemberUpdateOne {
	muo.mutation.SetCreatedAt(t)
	return muo
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableCreatedAt(t *time.Time) *MemberUpdateOne {
	if t != nil {
		muo.SetCreatedAt(*t)
	}
	return muo
}

// SetUpdatedAt sets the "updatedAt" field.
func (muo *MemberUpdateOne) SetUpdatedAt(t time.Time) *MemberUpdateOne {
	muo.mutation.SetUpdatedAt(t)
	return muo
}

// SetName sets the "name" field.
func (muo *MemberUpdateOne) SetName(s string) *MemberUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetEmail sets the "email" field.
func (muo *MemberUpdateOne) SetEmail(s string) *MemberUpdateOne {
	muo.mutation.SetEmail(s)
	return muo
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (muo *MemberUpdateOne) AddRoleIDs(ids ...uuid.UUID) *MemberUpdateOne {
	muo.mutation.AddRoleIDs(ids...)
	return muo
}

// AddRoles adds the "roles" edges to the Role entity.
func (muo *MemberUpdateOne) AddRoles(r ...*Role) *MemberUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return muo.AddRoleIDs(ids...)
}

// AddDeveloperOfIDs adds the "developerOf" edge to the Organization entity by IDs.
func (muo *MemberUpdateOne) AddDeveloperOfIDs(ids ...uuid.UUID) *MemberUpdateOne {
	muo.mutation.AddDeveloperOfIDs(ids...)
	return muo
}

// AddDeveloperOf adds the "developerOf" edges to the Organization entity.
func (muo *MemberUpdateOne) AddDeveloperOf(o ...*Organization) *MemberUpdateOne {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return muo.AddDeveloperOfIDs(ids...)
}

// AddManagerOfIDs adds the "managerOf" edge to the Organization entity by IDs.
func (muo *MemberUpdateOne) AddManagerOfIDs(ids ...uuid.UUID) *MemberUpdateOne {
	muo.mutation.AddManagerOfIDs(ids...)
	return muo
}

// AddManagerOf adds the "managerOf" edges to the Organization entity.
func (muo *MemberUpdateOne) AddManagerOf(o ...*Organization) *MemberUpdateOne {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return muo.AddManagerOfIDs(ids...)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (muo *MemberUpdateOne) AddTaskIDs(ids ...uuid.UUID) *MemberUpdateOne {
	muo.mutation.AddTaskIDs(ids...)
	return muo
}

// AddTasks adds the "tasks" edges to the Task entity.
func (muo *MemberUpdateOne) AddTasks(t ...*Task) *MemberUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return muo.AddTaskIDs(ids...)
}

// Mutation returns the MemberMutation object of the builder.
func (muo *MemberUpdateOne) Mutation() *MemberMutation {
	return muo.mutation
}

// ClearRoles clears all "roles" edges to the Role entity.
func (muo *MemberUpdateOne) ClearRoles() *MemberUpdateOne {
	muo.mutation.ClearRoles()
	return muo
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (muo *MemberUpdateOne) RemoveRoleIDs(ids ...uuid.UUID) *MemberUpdateOne {
	muo.mutation.RemoveRoleIDs(ids...)
	return muo
}

// RemoveRoles removes "roles" edges to Role entities.
func (muo *MemberUpdateOne) RemoveRoles(r ...*Role) *MemberUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return muo.RemoveRoleIDs(ids...)
}

// ClearDeveloperOf clears all "developerOf" edges to the Organization entity.
func (muo *MemberUpdateOne) ClearDeveloperOf() *MemberUpdateOne {
	muo.mutation.ClearDeveloperOf()
	return muo
}

// RemoveDeveloperOfIDs removes the "developerOf" edge to Organization entities by IDs.
func (muo *MemberUpdateOne) RemoveDeveloperOfIDs(ids ...uuid.UUID) *MemberUpdateOne {
	muo.mutation.RemoveDeveloperOfIDs(ids...)
	return muo
}

// RemoveDeveloperOf removes "developerOf" edges to Organization entities.
func (muo *MemberUpdateOne) RemoveDeveloperOf(o ...*Organization) *MemberUpdateOne {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return muo.RemoveDeveloperOfIDs(ids...)
}

// ClearManagerOf clears all "managerOf" edges to the Organization entity.
func (muo *MemberUpdateOne) ClearManagerOf() *MemberUpdateOne {
	muo.mutation.ClearManagerOf()
	return muo
}

// RemoveManagerOfIDs removes the "managerOf" edge to Organization entities by IDs.
func (muo *MemberUpdateOne) RemoveManagerOfIDs(ids ...uuid.UUID) *MemberUpdateOne {
	muo.mutation.RemoveManagerOfIDs(ids...)
	return muo
}

// RemoveManagerOf removes "managerOf" edges to Organization entities.
func (muo *MemberUpdateOne) RemoveManagerOf(o ...*Organization) *MemberUpdateOne {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return muo.RemoveManagerOfIDs(ids...)
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (muo *MemberUpdateOne) ClearTasks() *MemberUpdateOne {
	muo.mutation.ClearTasks()
	return muo
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (muo *MemberUpdateOne) RemoveTaskIDs(ids ...uuid.UUID) *MemberUpdateOne {
	muo.mutation.RemoveTaskIDs(ids...)
	return muo
}

// RemoveTasks removes "tasks" edges to Task entities.
func (muo *MemberUpdateOne) RemoveTasks(t ...*Task) *MemberUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return muo.RemoveTaskIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MemberUpdateOne) Select(field string, fields ...string) *MemberUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Member entity.
func (muo *MemberUpdateOne) Save(ctx context.Context) (*Member, error) {
	var (
		err  error
		node *Member
	)
	muo.defaults()
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MemberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MemberUpdateOne) SaveX(ctx context.Context) *Member {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MemberUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MemberUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MemberUpdateOne) defaults() {
	if _, ok := muo.mutation.UpdatedAt(); !ok {
		v := member.UpdateDefaultUpdatedAt()
		muo.mutation.SetUpdatedAt(v)
	}
}

func (muo *MemberUpdateOne) sqlSave(ctx context.Context) (_node *Member, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   member.Table,
			Columns: member.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: member.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Member.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, member.FieldID)
		for _, f := range fields {
			if !member.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != member.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: member.FieldCreatedAt,
		})
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: member.FieldUpdatedAt,
		})
	}
	if value, ok := muo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldName,
		})
	}
	if value, ok := muo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldEmail,
		})
	}
	if muo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   member.RolesTable,
			Columns: member.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: role.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedRolesIDs(); len(nodes) > 0 && !muo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   member.RolesTable,
			Columns: member.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   member.RolesTable,
			Columns: member.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.DeveloperOfCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   member.DeveloperOfTable,
			Columns: member.DeveloperOfPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: organization.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedDeveloperOfIDs(); len(nodes) > 0 && !muo.mutation.DeveloperOfCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   member.DeveloperOfTable,
			Columns: member.DeveloperOfPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: organization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.DeveloperOfIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   member.DeveloperOfTable,
			Columns: member.DeveloperOfPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: organization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.ManagerOfCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   member.ManagerOfTable,
			Columns: member.ManagerOfPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: organization.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedManagerOfIDs(); len(nodes) > 0 && !muo.mutation.ManagerOfCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   member.ManagerOfTable,
			Columns: member.ManagerOfPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: organization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ManagerOfIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   member.ManagerOfTable,
			Columns: member.ManagerOfPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: organization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   member.TasksTable,
			Columns: member.TasksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: task.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedTasksIDs(); len(nodes) > 0 && !muo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   member.TasksTable,
			Columns: member.TasksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   member.TasksTable,
			Columns: member.TasksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Member{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{member.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

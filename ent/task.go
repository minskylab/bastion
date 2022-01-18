// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/minskylab/bastion/ent/task"
)

// Task is the model entity for the Task schema.
type Task struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TaskQuery when eager-loading is set.
	Edges TaskEdges `json:"edges"`
}

// TaskEdges holds the relations/edges for other nodes in the graph.
type TaskEdges struct {
	// Project holds the value of the project edge.
	Project []*Project `json:"project,omitempty"`
	// Assignees holds the value of the assignees edge.
	Assignees []*Member `json:"assignees,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading.
func (e TaskEdges) ProjectOrErr() ([]*Project, error) {
	if e.loadedTypes[0] {
		return e.Project, nil
	}
	return nil, &NotLoadedError{edge: "project"}
}

// AssigneesOrErr returns the Assignees value or an error if the edge
// was not loaded in eager-loading.
func (e TaskEdges) AssigneesOrErr() ([]*Member, error) {
	if e.loadedTypes[1] {
		return e.Assignees, nil
	}
	return nil, &NotLoadedError{edge: "assignees"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Task) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case task.FieldName, task.FieldURL:
			values[i] = new(sql.NullString)
		case task.FieldCreatedAt, task.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case task.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Task", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Task fields.
func (t *Task) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case task.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case task.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case task.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		case task.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case task.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				t.URL = value.String
			}
		}
	}
	return nil
}

// QueryProject queries the "project" edge of the Task entity.
func (t *Task) QueryProject() *ProjectQuery {
	return (&TaskClient{config: t.config}).QueryProject(t)
}

// QueryAssignees queries the "assignees" edge of the Task entity.
func (t *Task) QueryAssignees() *MemberQuery {
	return (&TaskClient{config: t.config}).QueryAssignees(t)
}

// Update returns a builder for updating this Task.
// Note that you need to call Task.Unwrap() before calling this method if this Task
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Task) Update() *TaskUpdateOne {
	return (&TaskClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Task entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Task) Unwrap() *Task {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Task is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Task) String() string {
	var builder strings.Builder
	builder.WriteString("Task(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", createdAt=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(t.Name)
	builder.WriteString(", url=")
	builder.WriteString(t.URL)
	builder.WriteByte(')')
	return builder.String()
}

// Tasks is a parsable slice of Task.
type Tasks []*Task

func (t Tasks) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}

// Code generated by entc, DO NOT EDIT.

package organization

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the organization type in the database.
	Label = "organization"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeProjects holds the string denoting the projects edge name in mutations.
	EdgeProjects = "projects"
	// EdgeDevelopers holds the string denoting the developers edge name in mutations.
	EdgeDevelopers = "developers"
	// EdgeManagers holds the string denoting the managers edge name in mutations.
	EdgeManagers = "managers"
	// Table holds the table name of the organization in the database.
	Table = "organizations"
	// ProjectsTable is the table that holds the projects relation/edge. The primary key declared below.
	ProjectsTable = "organization_projects"
	// ProjectsInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectsInverseTable = "projects"
	// DevelopersTable is the table that holds the developers relation/edge. The primary key declared below.
	DevelopersTable = "organization_developers"
	// DevelopersInverseTable is the table name for the Member entity.
	// It exists in this package in order to avoid circular dependency with the "member" package.
	DevelopersInverseTable = "members"
	// ManagersTable is the table that holds the managers relation/edge. The primary key declared below.
	ManagersTable = "organization_managers"
	// ManagersInverseTable is the table name for the Member entity.
	// It exists in this package in order to avoid circular dependency with the "member" package.
	ManagersInverseTable = "members"
)

// Columns holds all SQL columns for organization fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
}

var (
	// ProjectsPrimaryKey and ProjectsColumn2 are the table columns denoting the
	// primary key for the projects relation (M2M).
	ProjectsPrimaryKey = []string{"organization_id", "project_id"}
	// DevelopersPrimaryKey and DevelopersColumn2 are the table columns denoting the
	// primary key for the developers relation (M2M).
	DevelopersPrimaryKey = []string{"organization_id", "member_id"}
	// ManagersPrimaryKey and ManagersColumn2 are the table columns denoting the
	// primary key for the managers relation (M2M).
	ManagersPrimaryKey = []string{"organization_id", "member_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updatedAt" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updatedAt" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

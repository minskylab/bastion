// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// MembersColumns holds the columns for the "members" table.
	MembersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Default: "gen_random_uuid()"},
		{Name: "created_at", Type: field.TypeTime, Default: "now()"},
		{Name: "updated_at", Type: field.TypeTime, Default: "now()"},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "auth_id", Type: field.TypeUUID},
	}
	// MembersTable holds the schema information for the "members" table.
	MembersTable = &schema.Table{
		Name:       "members",
		Columns:    MembersColumns,
		PrimaryKey: []*schema.Column{MembersColumns[0]},
	}
	// OrganizationsColumns holds the columns for the "organizations" table.
	OrganizationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Default: "gen_random_uuid()"},
		{Name: "created_at", Type: field.TypeTime, Default: "now()"},
		{Name: "updated_at", Type: field.TypeTime, Default: "now()"},
		{Name: "name", Type: field.TypeString},
	}
	// OrganizationsTable holds the schema information for the "organizations" table.
	OrganizationsTable = &schema.Table{
		Name:       "organizations",
		Columns:    OrganizationsColumns,
		PrimaryKey: []*schema.Column{OrganizationsColumns[0]},
	}
	// ProjectsColumns holds the columns for the "projects" table.
	ProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Default: "gen_random_uuid()"},
		{Name: "created_at", Type: field.TypeTime, Default: "now()"},
		{Name: "updated_at", Type: field.TypeTime, Default: "now()"},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "start_at", Type: field.TypeTime},
	}
	// ProjectsTable holds the schema information for the "projects" table.
	ProjectsTable = &schema.Table{
		Name:       "projects",
		Columns:    ProjectsColumns,
		PrimaryKey: []*schema.Column{ProjectsColumns[0]},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Default: "gen_random_uuid()"},
		{Name: "created_at", Type: field.TypeTime, Default: "now()"},
		{Name: "updated_at", Type: field.TypeTime, Default: "now()"},
		{Name: "name", Type: field.TypeString},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
	}
	// TasksColumns holds the columns for the "tasks" table.
	TasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Default: "gen_random_uuid()"},
		{Name: "created_at", Type: field.TypeTime, Default: "now()"},
		{Name: "updated_at", Type: field.TypeTime, Default: "now()"},
		{Name: "name", Type: field.TypeString},
		{Name: "url", Type: field.TypeString},
	}
	// TasksTable holds the schema information for the "tasks" table.
	TasksTable = &schema.Table{
		Name:       "tasks",
		Columns:    TasksColumns,
		PrimaryKey: []*schema.Column{TasksColumns[0]},
	}
	// MemberRolesColumns holds the columns for the "member_roles" table.
	MemberRolesColumns = []*schema.Column{
		{Name: "member_id", Type: field.TypeUUID},
		{Name: "role_id", Type: field.TypeUUID},
	}
	// MemberRolesTable holds the schema information for the "member_roles" table.
	MemberRolesTable = &schema.Table{
		Name:       "member_roles",
		Columns:    MemberRolesColumns,
		PrimaryKey: []*schema.Column{MemberRolesColumns[0], MemberRolesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "member_roles_member_id",
				Columns:    []*schema.Column{MemberRolesColumns[0]},
				RefColumns: []*schema.Column{MembersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "member_roles_role_id",
				Columns:    []*schema.Column{MemberRolesColumns[1]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// OrganizationProjectsColumns holds the columns for the "organization_projects" table.
	OrganizationProjectsColumns = []*schema.Column{
		{Name: "organization_id", Type: field.TypeUUID},
		{Name: "project_id", Type: field.TypeUUID},
	}
	// OrganizationProjectsTable holds the schema information for the "organization_projects" table.
	OrganizationProjectsTable = &schema.Table{
		Name:       "organization_projects",
		Columns:    OrganizationProjectsColumns,
		PrimaryKey: []*schema.Column{OrganizationProjectsColumns[0], OrganizationProjectsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "organization_projects_organization_id",
				Columns:    []*schema.Column{OrganizationProjectsColumns[0]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "organization_projects_project_id",
				Columns:    []*schema.Column{OrganizationProjectsColumns[1]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// OrganizationMembersColumns holds the columns for the "organization_members" table.
	OrganizationMembersColumns = []*schema.Column{
		{Name: "organization_id", Type: field.TypeUUID},
		{Name: "member_id", Type: field.TypeUUID},
	}
	// OrganizationMembersTable holds the schema information for the "organization_members" table.
	OrganizationMembersTable = &schema.Table{
		Name:       "organization_members",
		Columns:    OrganizationMembersColumns,
		PrimaryKey: []*schema.Column{OrganizationMembersColumns[0], OrganizationMembersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "organization_members_organization_id",
				Columns:    []*schema.Column{OrganizationMembersColumns[0]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "organization_members_member_id",
				Columns:    []*schema.Column{OrganizationMembersColumns[1]},
				RefColumns: []*schema.Column{MembersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ProjectRolesColumns holds the columns for the "project_roles" table.
	ProjectRolesColumns = []*schema.Column{
		{Name: "project_id", Type: field.TypeUUID},
		{Name: "role_id", Type: field.TypeUUID},
	}
	// ProjectRolesTable holds the schema information for the "project_roles" table.
	ProjectRolesTable = &schema.Table{
		Name:       "project_roles",
		Columns:    ProjectRolesColumns,
		PrimaryKey: []*schema.Column{ProjectRolesColumns[0], ProjectRolesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "project_roles_project_id",
				Columns:    []*schema.Column{ProjectRolesColumns[0]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "project_roles_role_id",
				Columns:    []*schema.Column{ProjectRolesColumns[1]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ProjectTasksColumns holds the columns for the "project_tasks" table.
	ProjectTasksColumns = []*schema.Column{
		{Name: "project_id", Type: field.TypeUUID},
		{Name: "task_id", Type: field.TypeUUID},
	}
	// ProjectTasksTable holds the schema information for the "project_tasks" table.
	ProjectTasksTable = &schema.Table{
		Name:       "project_tasks",
		Columns:    ProjectTasksColumns,
		PrimaryKey: []*schema.Column{ProjectTasksColumns[0], ProjectTasksColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "project_tasks_project_id",
				Columns:    []*schema.Column{ProjectTasksColumns[0]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "project_tasks_task_id",
				Columns:    []*schema.Column{ProjectTasksColumns[1]},
				RefColumns: []*schema.Column{TasksColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TaskAssigneesColumns holds the columns for the "task_assignees" table.
	TaskAssigneesColumns = []*schema.Column{
		{Name: "task_id", Type: field.TypeUUID},
		{Name: "member_id", Type: field.TypeUUID},
	}
	// TaskAssigneesTable holds the schema information for the "task_assignees" table.
	TaskAssigneesTable = &schema.Table{
		Name:       "task_assignees",
		Columns:    TaskAssigneesColumns,
		PrimaryKey: []*schema.Column{TaskAssigneesColumns[0], TaskAssigneesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "task_assignees_task_id",
				Columns:    []*schema.Column{TaskAssigneesColumns[0]},
				RefColumns: []*schema.Column{TasksColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "task_assignees_member_id",
				Columns:    []*schema.Column{TaskAssigneesColumns[1]},
				RefColumns: []*schema.Column{MembersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		MembersTable,
		OrganizationsTable,
		ProjectsTable,
		RolesTable,
		TasksTable,
		MemberRolesTable,
		OrganizationProjectsTable,
		OrganizationMembersTable,
		ProjectRolesTable,
		ProjectTasksTable,
		TaskAssigneesTable,
	}
)

func init() {
	MemberRolesTable.ForeignKeys[0].RefTable = MembersTable
	MemberRolesTable.ForeignKeys[1].RefTable = RolesTable
	OrganizationProjectsTable.ForeignKeys[0].RefTable = OrganizationsTable
	OrganizationProjectsTable.ForeignKeys[1].RefTable = ProjectsTable
	OrganizationMembersTable.ForeignKeys[0].RefTable = OrganizationsTable
	OrganizationMembersTable.ForeignKeys[1].RefTable = MembersTable
	ProjectRolesTable.ForeignKeys[0].RefTable = ProjectsTable
	ProjectRolesTable.ForeignKeys[1].RefTable = RolesTable
	ProjectTasksTable.ForeignKeys[0].RefTable = ProjectsTable
	ProjectTasksTable.ForeignKeys[1].RefTable = TasksTable
	TaskAssigneesTable.ForeignKeys[0].RefTable = TasksTable
	TaskAssigneesTable.ForeignKeys[1].RefTable = MembersTable
}

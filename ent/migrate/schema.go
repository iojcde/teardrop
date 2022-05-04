// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DeploymentsColumns holds the columns for the "deployments" table.
	DeploymentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "branch", Type: field.TypeString},
		{Name: "address", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "project_deployments", Type: field.TypeInt, Nullable: true},
	}
	// DeploymentsTable holds the schema information for the "deployments" table.
	DeploymentsTable = &schema.Table{
		Name:       "deployments",
		Columns:    DeploymentsColumns,
		PrimaryKey: []*schema.Column{DeploymentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "deployments_projects_deployments",
				Columns:    []*schema.Column{DeploymentsColumns[4]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DomainsColumns holds the columns for the "domains" table.
	DomainsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "project_id", Type: field.TypeInt},
		{Name: "domain", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
		{Name: "deployment_domains", Type: field.TypeString, Nullable: true},
	}
	// DomainsTable holds the schema information for the "domains" table.
	DomainsTable = &schema.Table{
		Name:       "domains",
		Columns:    DomainsColumns,
		PrimaryKey: []*schema.Column{DomainsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "domains_deployments_domains",
				Columns:    []*schema.Column{DomainsColumns[5]},
				RefColumns: []*schema.Column{DeploymentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProjectsColumns holds the columns for the "projects" table.
	ProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "git", Type: field.TypeString},
		{Name: "default_branch", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
		{Name: "deployment_projects", Type: field.TypeString, Nullable: true},
	}
	// ProjectsTable holds the schema information for the "projects" table.
	ProjectsTable = &schema.Table{
		Name:       "projects",
		Columns:    ProjectsColumns,
		PrimaryKey: []*schema.Column{ProjectsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "projects_deployments_projects",
				Columns:    []*schema.Column{ProjectsColumns[6]},
				RefColumns: []*schema.Column{DeploymentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(30)"}},
		{Name: "email", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(69)"}},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// ProjectUsersColumns holds the columns for the "project_users" table.
	ProjectUsersColumns = []*schema.Column{
		{Name: "project_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// ProjectUsersTable holds the schema information for the "project_users" table.
	ProjectUsersTable = &schema.Table{
		Name:       "project_users",
		Columns:    ProjectUsersColumns,
		PrimaryKey: []*schema.Column{ProjectUsersColumns[0], ProjectUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "project_users_project_id",
				Columns:    []*schema.Column{ProjectUsersColumns[0]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "project_users_user_id",
				Columns:    []*schema.Column{ProjectUsersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DeploymentsTable,
		DomainsTable,
		ProjectsTable,
		UsersTable,
		ProjectUsersTable,
	}
)

func init() {
	DeploymentsTable.ForeignKeys[0].RefTable = ProjectsTable
	DomainsTable.ForeignKeys[0].RefTable = DeploymentsTable
	ProjectsTable.ForeignKeys[0].RefTable = DeploymentsTable
	ProjectUsersTable.ForeignKeys[0].RefTable = ProjectsTable
	ProjectUsersTable.ForeignKeys[1].RefTable = UsersTable
}

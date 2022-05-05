// Code generated by entc, DO NOT EDIT.

package domain

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the domain type in the database.
	Label = "domain"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDomain holds the string denoting the domain field in the database.
	FieldDomain = "domain"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// EdgeDeployment holds the string denoting the deployment edge name in mutations.
	EdgeDeployment = "deployment"
	// EdgeProject holds the string denoting the project edge name in mutations.
	EdgeProject = "project"
	// Table holds the table name of the domain in the database.
	Table = "domains"
	// DeploymentTable is the table that holds the deployment relation/edge.
	DeploymentTable = "domains"
	// DeploymentInverseTable is the table name for the Deployment entity.
	// It exists in this package in order to avoid circular dependency with the "deployment" package.
	DeploymentInverseTable = "deployments"
	// DeploymentColumn is the table column denoting the deployment relation/edge.
	DeploymentColumn = "deployment_domains"
	// ProjectTable is the table that holds the project relation/edge.
	ProjectTable = "domains"
	// ProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectInverseTable = "projects"
	// ProjectColumn is the table column denoting the project relation/edge.
	ProjectColumn = "project_domains"
)

// Columns holds all SQL columns for domain fields.
var Columns = []string{
	FieldID,
	FieldDomain,
	FieldCreateAt,
	FieldUpdateAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "domains"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"deployment_domains",
	"project_domains",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() time.Time
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() time.Time
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

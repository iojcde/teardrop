// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/fosshostorg/teardrop/ent/migrate"
	"github.com/google/uuid"

	"github.com/fosshostorg/teardrop/ent/deployment"
	"github.com/fosshostorg/teardrop/ent/domain"
	"github.com/fosshostorg/teardrop/ent/project"
	"github.com/fosshostorg/teardrop/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Deployment is the client for interacting with the Deployment builders.
	Deployment *DeploymentClient
	// Domain is the client for interacting with the Domain builders.
	Domain *DomainClient
	// Project is the client for interacting with the Project builders.
	Project *ProjectClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Deployment = NewDeploymentClient(c.config)
	c.Domain = NewDomainClient(c.config)
	c.Project = NewProjectClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Deployment: NewDeploymentClient(cfg),
		Domain:     NewDomainClient(cfg),
		Project:    NewProjectClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Deployment: NewDeploymentClient(cfg),
		Domain:     NewDomainClient(cfg),
		Project:    NewProjectClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Deployment.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Deployment.Use(hooks...)
	c.Domain.Use(hooks...)
	c.Project.Use(hooks...)
	c.User.Use(hooks...)
}

// DeploymentClient is a client for the Deployment schema.
type DeploymentClient struct {
	config
}

// NewDeploymentClient returns a client for the Deployment from the given config.
func NewDeploymentClient(c config) *DeploymentClient {
	return &DeploymentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `deployment.Hooks(f(g(h())))`.
func (c *DeploymentClient) Use(hooks ...Hook) {
	c.hooks.Deployment = append(c.hooks.Deployment, hooks...)
}

// Create returns a create builder for Deployment.
func (c *DeploymentClient) Create() *DeploymentCreate {
	mutation := newDeploymentMutation(c.config, OpCreate)
	return &DeploymentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Deployment entities.
func (c *DeploymentClient) CreateBulk(builders ...*DeploymentCreate) *DeploymentCreateBulk {
	return &DeploymentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Deployment.
func (c *DeploymentClient) Update() *DeploymentUpdate {
	mutation := newDeploymentMutation(c.config, OpUpdate)
	return &DeploymentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DeploymentClient) UpdateOne(d *Deployment) *DeploymentUpdateOne {
	mutation := newDeploymentMutation(c.config, OpUpdateOne, withDeployment(d))
	return &DeploymentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DeploymentClient) UpdateOneID(id uuid.UUID) *DeploymentUpdateOne {
	mutation := newDeploymentMutation(c.config, OpUpdateOne, withDeploymentID(id))
	return &DeploymentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Deployment.
func (c *DeploymentClient) Delete() *DeploymentDelete {
	mutation := newDeploymentMutation(c.config, OpDelete)
	return &DeploymentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DeploymentClient) DeleteOne(d *Deployment) *DeploymentDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DeploymentClient) DeleteOneID(id uuid.UUID) *DeploymentDeleteOne {
	builder := c.Delete().Where(deployment.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DeploymentDeleteOne{builder}
}

// Query returns a query builder for Deployment.
func (c *DeploymentClient) Query() *DeploymentQuery {
	return &DeploymentQuery{
		config: c.config,
	}
}

// Get returns a Deployment entity by its id.
func (c *DeploymentClient) Get(ctx context.Context, id uuid.UUID) (*Deployment, error) {
	return c.Query().Where(deployment.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DeploymentClient) GetX(ctx context.Context, id uuid.UUID) *Deployment {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryProject queries the project edge of a Deployment.
func (c *DeploymentClient) QueryProject(d *Deployment) *ProjectQuery {
	query := &ProjectQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(deployment.Table, deployment.FieldID, id),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, deployment.ProjectTable, deployment.ProjectColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDomains queries the domains edge of a Deployment.
func (c *DeploymentClient) QueryDomains(d *Deployment) *DomainQuery {
	query := &DomainQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(deployment.Table, deployment.FieldID, id),
			sqlgraph.To(domain.Table, domain.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, deployment.DomainsTable, deployment.DomainsColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DeploymentClient) Hooks() []Hook {
	return c.hooks.Deployment
}

// DomainClient is a client for the Domain schema.
type DomainClient struct {
	config
}

// NewDomainClient returns a client for the Domain from the given config.
func NewDomainClient(c config) *DomainClient {
	return &DomainClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `domain.Hooks(f(g(h())))`.
func (c *DomainClient) Use(hooks ...Hook) {
	c.hooks.Domain = append(c.hooks.Domain, hooks...)
}

// Create returns a create builder for Domain.
func (c *DomainClient) Create() *DomainCreate {
	mutation := newDomainMutation(c.config, OpCreate)
	return &DomainCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Domain entities.
func (c *DomainClient) CreateBulk(builders ...*DomainCreate) *DomainCreateBulk {
	return &DomainCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Domain.
func (c *DomainClient) Update() *DomainUpdate {
	mutation := newDomainMutation(c.config, OpUpdate)
	return &DomainUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DomainClient) UpdateOne(d *Domain) *DomainUpdateOne {
	mutation := newDomainMutation(c.config, OpUpdateOne, withDomain(d))
	return &DomainUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DomainClient) UpdateOneID(id uuid.UUID) *DomainUpdateOne {
	mutation := newDomainMutation(c.config, OpUpdateOne, withDomainID(id))
	return &DomainUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Domain.
func (c *DomainClient) Delete() *DomainDelete {
	mutation := newDomainMutation(c.config, OpDelete)
	return &DomainDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DomainClient) DeleteOne(d *Domain) *DomainDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DomainClient) DeleteOneID(id uuid.UUID) *DomainDeleteOne {
	builder := c.Delete().Where(domain.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DomainDeleteOne{builder}
}

// Query returns a query builder for Domain.
func (c *DomainClient) Query() *DomainQuery {
	return &DomainQuery{
		config: c.config,
	}
}

// Get returns a Domain entity by its id.
func (c *DomainClient) Get(ctx context.Context, id uuid.UUID) (*Domain, error) {
	return c.Query().Where(domain.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DomainClient) GetX(ctx context.Context, id uuid.UUID) *Domain {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDeployment queries the deployment edge of a Domain.
func (c *DomainClient) QueryDeployment(d *Domain) *DeploymentQuery {
	query := &DeploymentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(domain.Table, domain.FieldID, id),
			sqlgraph.To(deployment.Table, deployment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, domain.DeploymentTable, domain.DeploymentColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryProject queries the project edge of a Domain.
func (c *DomainClient) QueryProject(d *Domain) *ProjectQuery {
	query := &ProjectQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(domain.Table, domain.FieldID, id),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, domain.ProjectTable, domain.ProjectColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DomainClient) Hooks() []Hook {
	return c.hooks.Domain
}

// ProjectClient is a client for the Project schema.
type ProjectClient struct {
	config
}

// NewProjectClient returns a client for the Project from the given config.
func NewProjectClient(c config) *ProjectClient {
	return &ProjectClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `project.Hooks(f(g(h())))`.
func (c *ProjectClient) Use(hooks ...Hook) {
	c.hooks.Project = append(c.hooks.Project, hooks...)
}

// Create returns a create builder for Project.
func (c *ProjectClient) Create() *ProjectCreate {
	mutation := newProjectMutation(c.config, OpCreate)
	return &ProjectCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Project entities.
func (c *ProjectClient) CreateBulk(builders ...*ProjectCreate) *ProjectCreateBulk {
	return &ProjectCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Project.
func (c *ProjectClient) Update() *ProjectUpdate {
	mutation := newProjectMutation(c.config, OpUpdate)
	return &ProjectUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProjectClient) UpdateOne(pr *Project) *ProjectUpdateOne {
	mutation := newProjectMutation(c.config, OpUpdateOne, withProject(pr))
	return &ProjectUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProjectClient) UpdateOneID(id uuid.UUID) *ProjectUpdateOne {
	mutation := newProjectMutation(c.config, OpUpdateOne, withProjectID(id))
	return &ProjectUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Project.
func (c *ProjectClient) Delete() *ProjectDelete {
	mutation := newProjectMutation(c.config, OpDelete)
	return &ProjectDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ProjectClient) DeleteOne(pr *Project) *ProjectDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ProjectClient) DeleteOneID(id uuid.UUID) *ProjectDeleteOne {
	builder := c.Delete().Where(project.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProjectDeleteOne{builder}
}

// Query returns a query builder for Project.
func (c *ProjectClient) Query() *ProjectQuery {
	return &ProjectQuery{
		config: c.config,
	}
}

// Get returns a Project entity by its id.
func (c *ProjectClient) Get(ctx context.Context, id uuid.UUID) (*Project, error) {
	return c.Query().Where(project.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProjectClient) GetX(ctx context.Context, id uuid.UUID) *Project {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUsers queries the users edge of a Project.
func (c *ProjectClient) QueryUsers(pr *Project) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(project.Table, project.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, project.UsersTable, project.UsersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDeployments queries the deployments edge of a Project.
func (c *ProjectClient) QueryDeployments(pr *Project) *DeploymentQuery {
	query := &DeploymentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(project.Table, project.FieldID, id),
			sqlgraph.To(deployment.Table, deployment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, project.DeploymentsTable, project.DeploymentsColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDomains queries the domains edge of a Project.
func (c *ProjectClient) QueryDomains(pr *Project) *DomainQuery {
	query := &DomainQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(project.Table, project.FieldID, id),
			sqlgraph.To(domain.Table, domain.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, project.DomainsTable, project.DomainsColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ProjectClient) Hooks() []Hook {
	return c.hooks.Project
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id uuid.UUID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id uuid.UUID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uuid.UUID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryProjects queries the projects edge of a User.
func (c *UserClient) QueryProjects(u *User) *ProjectQuery {
	query := &ProjectQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, user.ProjectsTable, user.ProjectsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
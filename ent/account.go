// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/fosshostorg/teardrop/ent/account"
	"github.com/fosshostorg/teardrop/ent/user"
	"github.com/google/uuid"
)

// Account is the model entity for the Account schema.
type Account struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// AccessToken holds the value of the "access_token" field.
	AccessToken string `json:"access_token,omitempty"`
	// RefreshToken holds the value of the "refresh_token" field.
	RefreshToken string `json:"refresh_token,omitempty"`
	// TokenType holds the value of the "token_type" field.
	TokenType string `json:"token_type,omitempty"`
	// ExpiresAt holds the value of the "expires_at" field.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// ProviderAccountId holds the value of the "providerAccountId" field.
	ProviderAccountId string `json:"providerAccountId,omitempty"`
	// Scope holds the value of the "scope" field.
	Scope string `json:"scope,omitempty"`
	// Provider holds the value of the "provider" field.
	Provider string `json:"provider,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AccountQuery when eager-loading is set.
	Edges         AccountEdges `json:"edges"`
	user_accounts *uuid.UUID
}

// AccountEdges holds the relations/edges for other nodes in the graph.
type AccountEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccountEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Account) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case account.FieldAccessToken, account.FieldRefreshToken, account.FieldTokenType, account.FieldProviderAccountId, account.FieldScope, account.FieldProvider:
			values[i] = new(sql.NullString)
		case account.FieldExpiresAt:
			values[i] = new(sql.NullTime)
		case account.FieldID:
			values[i] = new(uuid.UUID)
		case account.ForeignKeys[0]: // user_accounts
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Account", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Account fields.
func (a *Account) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case account.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case account.FieldAccessToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field access_token", values[i])
			} else if value.Valid {
				a.AccessToken = value.String
			}
		case account.FieldRefreshToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field refresh_token", values[i])
			} else if value.Valid {
				a.RefreshToken = value.String
			}
		case account.FieldTokenType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token_type", values[i])
			} else if value.Valid {
				a.TokenType = value.String
			}
		case account.FieldExpiresAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expires_at", values[i])
			} else if value.Valid {
				a.ExpiresAt = value.Time
			}
		case account.FieldProviderAccountId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field providerAccountId", values[i])
			} else if value.Valid {
				a.ProviderAccountId = value.String
			}
		case account.FieldScope:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field scope", values[i])
			} else if value.Valid {
				a.Scope = value.String
			}
		case account.FieldProvider:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provider", values[i])
			} else if value.Valid {
				a.Provider = value.String
			}
		case account.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_accounts", values[i])
			} else if value.Valid {
				a.user_accounts = new(uuid.UUID)
				*a.user_accounts = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Account entity.
func (a *Account) QueryUser() *UserQuery {
	return (&AccountClient{config: a.config}).QueryUser(a)
}

// Update returns a builder for updating this Account.
// Note that you need to call Account.Unwrap() before calling this method if this Account
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Account) Update() *AccountUpdateOne {
	return (&AccountClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Account entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Account) Unwrap() *Account {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Account is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Account) String() string {
	var builder strings.Builder
	builder.WriteString("Account(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", access_token=")
	builder.WriteString(a.AccessToken)
	builder.WriteString(", refresh_token=")
	builder.WriteString(a.RefreshToken)
	builder.WriteString(", token_type=")
	builder.WriteString(a.TokenType)
	builder.WriteString(", expires_at=")
	builder.WriteString(a.ExpiresAt.Format(time.ANSIC))
	builder.WriteString(", providerAccountId=")
	builder.WriteString(a.ProviderAccountId)
	builder.WriteString(", scope=")
	builder.WriteString(a.Scope)
	builder.WriteString(", provider=")
	builder.WriteString(a.Provider)
	builder.WriteByte(')')
	return builder.String()
}

// Accounts is a parsable slice of Account.
type Accounts []*Account

func (a Accounts) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}

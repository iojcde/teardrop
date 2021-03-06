package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Domain holds the schema definition for the Domain entity.
type Domain struct {
	ent.Schema
}

// Fields of the Domain.
func (Domain) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("domain"),
		field.Time("create_at").Default(time.Now),
		field.Time("update_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Domain.
func (Domain) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("deployment", Deployment.Type).
			Ref("domains").
			Unique(),
		edge.From("project", Project.Type).
			Ref("domains").
			Unique(),
	}
}

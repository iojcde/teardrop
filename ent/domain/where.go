// Code generated by entc, DO NOT EDIT.

package domain

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/fosshostorg/teardrop/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Domain applies equality check predicate on the "domain" field. It's identical to DomainEQ.
func Domain(v string) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDomain), v))
	})
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v time.Time) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v time.Time) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// DomainEQ applies the EQ predicate on the "domain" field.
func DomainEQ(v string) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDomain), v))
	})
}

// DomainNEQ applies the NEQ predicate on the "domain" field.
func DomainNEQ(v string) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDomain), v))
	})
}

// DomainIn applies the In predicate on the "domain" field.
func DomainIn(vs ...string) predicate.Domain {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Domain(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDomain), v...))
	})
}

// DomainNotIn applies the NotIn predicate on the "domain" field.
func DomainNotIn(vs ...string) predicate.Domain {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Domain(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDomain), v...))
	})
}

// DomainGT applies the GT predicate on the "domain" field.
func DomainGT(v string) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDomain), v))
	})
}

// DomainGTE applies the GTE predicate on the "domain" field.
func DomainGTE(v string) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDomain), v))
	})
}

// DomainLT applies the LT predicate on the "domain" field.
func DomainLT(v string) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDomain), v))
	})
}

// DomainLTE applies the LTE predicate on the "domain" field.
func DomainLTE(v string) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDomain), v))
	})
}

// DomainContains applies the Contains predicate on the "domain" field.
func DomainContains(v string) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDomain), v))
	})
}

// DomainHasPrefix applies the HasPrefix predicate on the "domain" field.
func DomainHasPrefix(v string) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDomain), v))
	})
}

// DomainHasSuffix applies the HasSuffix predicate on the "domain" field.
func DomainHasSuffix(v string) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDomain), v))
	})
}

// DomainEqualFold applies the EqualFold predicate on the "domain" field.
func DomainEqualFold(v string) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDomain), v))
	})
}

// DomainContainsFold applies the ContainsFold predicate on the "domain" field.
func DomainContainsFold(v string) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDomain), v))
	})
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v time.Time) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v time.Time) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...time.Time) predicate.Domain {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Domain(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateAt), v...))
	})
}

// CreateAtNotIn applies the NotIn predicate on the "create_at" field.
func CreateAtNotIn(vs ...time.Time) predicate.Domain {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Domain(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateAt), v...))
	})
}

// CreateAtGT applies the GT predicate on the "create_at" field.
func CreateAtGT(v time.Time) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateAt), v))
	})
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v time.Time) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateAt), v))
	})
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v time.Time) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateAt), v))
	})
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v time.Time) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateAt), v))
	})
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v time.Time) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v time.Time) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...time.Time) predicate.Domain {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Domain(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtNotIn applies the NotIn predicate on the "update_at" field.
func UpdateAtNotIn(vs ...time.Time) predicate.Domain {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Domain(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtGT applies the GT predicate on the "update_at" field.
func UpdateAtGT(v time.Time) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v time.Time) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v time.Time) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v time.Time) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateAt), v))
	})
}

// HasDeployment applies the HasEdge predicate on the "deployment" edge.
func HasDeployment() predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DeploymentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DeploymentTable, DeploymentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDeploymentWith applies the HasEdge predicate on the "deployment" edge with a given conditions (other predicates).
func HasDeploymentWith(preds ...predicate.Deployment) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DeploymentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DeploymentTable, DeploymentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProject applies the HasEdge predicate on the "project" edge.
func HasProject() predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProjectTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectWith applies the HasEdge predicate on the "project" edge with a given conditions (other predicates).
func HasProjectWith(preds ...predicate.Project) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProjectInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Domain) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Domain) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Domain) predicate.Domain {
	return predicate.Domain(func(s *sql.Selector) {
		p(s.Not())
	})
}
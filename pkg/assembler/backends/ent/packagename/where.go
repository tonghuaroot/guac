// Code generated by ent, DO NOT EDIT.

package packagename

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.PackageName {
	return predicate.PackageName(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.PackageName {
	return predicate.PackageName(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.PackageName {
	return predicate.PackageName(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.PackageName {
	return predicate.PackageName(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.PackageName {
	return predicate.PackageName(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.PackageName {
	return predicate.PackageName(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.PackageName {
	return predicate.PackageName(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.PackageName {
	return predicate.PackageName(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.PackageName {
	return predicate.PackageName(sql.FieldLTE(FieldID, id))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldEQ(FieldType, v))
}

// Namespace applies equality check predicate on the "namespace" field. It's identical to NamespaceEQ.
func Namespace(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldEQ(FieldNamespace, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldEQ(FieldName, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.PackageName {
	return predicate.PackageName(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.PackageName {
	return predicate.PackageName(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldContainsFold(FieldType, v))
}

// NamespaceEQ applies the EQ predicate on the "namespace" field.
func NamespaceEQ(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldEQ(FieldNamespace, v))
}

// NamespaceNEQ applies the NEQ predicate on the "namespace" field.
func NamespaceNEQ(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldNEQ(FieldNamespace, v))
}

// NamespaceIn applies the In predicate on the "namespace" field.
func NamespaceIn(vs ...string) predicate.PackageName {
	return predicate.PackageName(sql.FieldIn(FieldNamespace, vs...))
}

// NamespaceNotIn applies the NotIn predicate on the "namespace" field.
func NamespaceNotIn(vs ...string) predicate.PackageName {
	return predicate.PackageName(sql.FieldNotIn(FieldNamespace, vs...))
}

// NamespaceGT applies the GT predicate on the "namespace" field.
func NamespaceGT(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldGT(FieldNamespace, v))
}

// NamespaceGTE applies the GTE predicate on the "namespace" field.
func NamespaceGTE(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldGTE(FieldNamespace, v))
}

// NamespaceLT applies the LT predicate on the "namespace" field.
func NamespaceLT(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldLT(FieldNamespace, v))
}

// NamespaceLTE applies the LTE predicate on the "namespace" field.
func NamespaceLTE(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldLTE(FieldNamespace, v))
}

// NamespaceContains applies the Contains predicate on the "namespace" field.
func NamespaceContains(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldContains(FieldNamespace, v))
}

// NamespaceHasPrefix applies the HasPrefix predicate on the "namespace" field.
func NamespaceHasPrefix(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldHasPrefix(FieldNamespace, v))
}

// NamespaceHasSuffix applies the HasSuffix predicate on the "namespace" field.
func NamespaceHasSuffix(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldHasSuffix(FieldNamespace, v))
}

// NamespaceEqualFold applies the EqualFold predicate on the "namespace" field.
func NamespaceEqualFold(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldEqualFold(FieldNamespace, v))
}

// NamespaceContainsFold applies the ContainsFold predicate on the "namespace" field.
func NamespaceContainsFold(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldContainsFold(FieldNamespace, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.PackageName {
	return predicate.PackageName(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.PackageName {
	return predicate.PackageName(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.PackageName {
	return predicate.PackageName(sql.FieldContainsFold(FieldName, v))
}

// HasVersions applies the HasEdge predicate on the "versions" edge.
func HasVersions() predicate.PackageName {
	return predicate.PackageName(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, VersionsTable, VersionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVersionsWith applies the HasEdge predicate on the "versions" edge with a given conditions (other predicates).
func HasVersionsWith(preds ...predicate.PackageVersion) predicate.PackageName {
	return predicate.PackageName(func(s *sql.Selector) {
		step := newVersionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasHasSourceAt applies the HasEdge predicate on the "has_source_at" edge.
func HasHasSourceAt() predicate.PackageName {
	return predicate.PackageName(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, HasSourceAtTable, HasSourceAtColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHasSourceAtWith applies the HasEdge predicate on the "has_source_at" edge with a given conditions (other predicates).
func HasHasSourceAtWith(preds ...predicate.HasSourceAt) predicate.PackageName {
	return predicate.PackageName(func(s *sql.Selector) {
		step := newHasSourceAtStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCertification applies the HasEdge predicate on the "certification" edge.
func HasCertification() predicate.PackageName {
	return predicate.PackageName(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, CertificationTable, CertificationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCertificationWith applies the HasEdge predicate on the "certification" edge with a given conditions (other predicates).
func HasCertificationWith(preds ...predicate.Certification) predicate.PackageName {
	return predicate.PackageName(func(s *sql.Selector) {
		step := newCertificationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMetadata applies the HasEdge predicate on the "metadata" edge.
func HasMetadata() predicate.PackageName {
	return predicate.PackageName(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, MetadataTable, MetadataColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMetadataWith applies the HasEdge predicate on the "metadata" edge with a given conditions (other predicates).
func HasMetadataWith(preds ...predicate.HasMetadata) predicate.PackageName {
	return predicate.PackageName(func(s *sql.Selector) {
		step := newMetadataStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPoc applies the HasEdge predicate on the "poc" edge.
func HasPoc() predicate.PackageName {
	return predicate.PackageName(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, PocTable, PocColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPocWith applies the HasEdge predicate on the "poc" edge with a given conditions (other predicates).
func HasPocWith(preds ...predicate.PointOfContact) predicate.PackageName {
	return predicate.PackageName(func(s *sql.Selector) {
		step := newPocStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PackageName) predicate.PackageName {
	return predicate.PackageName(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PackageName) predicate.PackageName {
	return predicate.PackageName(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PackageName) predicate.PackageName {
	return predicate.PackageName(sql.NotPredicates(p))
}

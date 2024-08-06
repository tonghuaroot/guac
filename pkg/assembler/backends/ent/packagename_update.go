// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/certification"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/hasmetadata"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/hassourceat"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packagename"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packageversion"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/pointofcontact"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
)

// PackageNameUpdate is the builder for updating PackageName entities.
type PackageNameUpdate struct {
	config
	hooks    []Hook
	mutation *PackageNameMutation
}

// Where appends a list predicates to the PackageNameUpdate builder.
func (pnu *PackageNameUpdate) Where(ps ...predicate.PackageName) *PackageNameUpdate {
	pnu.mutation.Where(ps...)
	return pnu
}

// SetType sets the "type" field.
func (pnu *PackageNameUpdate) SetType(s string) *PackageNameUpdate {
	pnu.mutation.SetType(s)
	return pnu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (pnu *PackageNameUpdate) SetNillableType(s *string) *PackageNameUpdate {
	if s != nil {
		pnu.SetType(*s)
	}
	return pnu
}

// SetNamespace sets the "namespace" field.
func (pnu *PackageNameUpdate) SetNamespace(s string) *PackageNameUpdate {
	pnu.mutation.SetNamespace(s)
	return pnu
}

// SetNillableNamespace sets the "namespace" field if the given value is not nil.
func (pnu *PackageNameUpdate) SetNillableNamespace(s *string) *PackageNameUpdate {
	if s != nil {
		pnu.SetNamespace(*s)
	}
	return pnu
}

// SetName sets the "name" field.
func (pnu *PackageNameUpdate) SetName(s string) *PackageNameUpdate {
	pnu.mutation.SetName(s)
	return pnu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pnu *PackageNameUpdate) SetNillableName(s *string) *PackageNameUpdate {
	if s != nil {
		pnu.SetName(*s)
	}
	return pnu
}

// AddVersionIDs adds the "versions" edge to the PackageVersion entity by IDs.
func (pnu *PackageNameUpdate) AddVersionIDs(ids ...uuid.UUID) *PackageNameUpdate {
	pnu.mutation.AddVersionIDs(ids...)
	return pnu
}

// AddVersions adds the "versions" edges to the PackageVersion entity.
func (pnu *PackageNameUpdate) AddVersions(p ...*PackageVersion) *PackageNameUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pnu.AddVersionIDs(ids...)
}

// AddHasSourceAtIDs adds the "has_source_at" edge to the HasSourceAt entity by IDs.
func (pnu *PackageNameUpdate) AddHasSourceAtIDs(ids ...uuid.UUID) *PackageNameUpdate {
	pnu.mutation.AddHasSourceAtIDs(ids...)
	return pnu
}

// AddHasSourceAt adds the "has_source_at" edges to the HasSourceAt entity.
func (pnu *PackageNameUpdate) AddHasSourceAt(h ...*HasSourceAt) *PackageNameUpdate {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return pnu.AddHasSourceAtIDs(ids...)
}

// AddCertificationIDs adds the "certification" edge to the Certification entity by IDs.
func (pnu *PackageNameUpdate) AddCertificationIDs(ids ...uuid.UUID) *PackageNameUpdate {
	pnu.mutation.AddCertificationIDs(ids...)
	return pnu
}

// AddCertification adds the "certification" edges to the Certification entity.
func (pnu *PackageNameUpdate) AddCertification(c ...*Certification) *PackageNameUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pnu.AddCertificationIDs(ids...)
}

// AddMetadatumIDs adds the "metadata" edge to the HasMetadata entity by IDs.
func (pnu *PackageNameUpdate) AddMetadatumIDs(ids ...uuid.UUID) *PackageNameUpdate {
	pnu.mutation.AddMetadatumIDs(ids...)
	return pnu
}

// AddMetadata adds the "metadata" edges to the HasMetadata entity.
func (pnu *PackageNameUpdate) AddMetadata(h ...*HasMetadata) *PackageNameUpdate {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return pnu.AddMetadatumIDs(ids...)
}

// AddPocIDs adds the "poc" edge to the PointOfContact entity by IDs.
func (pnu *PackageNameUpdate) AddPocIDs(ids ...uuid.UUID) *PackageNameUpdate {
	pnu.mutation.AddPocIDs(ids...)
	return pnu
}

// AddPoc adds the "poc" edges to the PointOfContact entity.
func (pnu *PackageNameUpdate) AddPoc(p ...*PointOfContact) *PackageNameUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pnu.AddPocIDs(ids...)
}

// Mutation returns the PackageNameMutation object of the builder.
func (pnu *PackageNameUpdate) Mutation() *PackageNameMutation {
	return pnu.mutation
}

// ClearVersions clears all "versions" edges to the PackageVersion entity.
func (pnu *PackageNameUpdate) ClearVersions() *PackageNameUpdate {
	pnu.mutation.ClearVersions()
	return pnu
}

// RemoveVersionIDs removes the "versions" edge to PackageVersion entities by IDs.
func (pnu *PackageNameUpdate) RemoveVersionIDs(ids ...uuid.UUID) *PackageNameUpdate {
	pnu.mutation.RemoveVersionIDs(ids...)
	return pnu
}

// RemoveVersions removes "versions" edges to PackageVersion entities.
func (pnu *PackageNameUpdate) RemoveVersions(p ...*PackageVersion) *PackageNameUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pnu.RemoveVersionIDs(ids...)
}

// ClearHasSourceAt clears all "has_source_at" edges to the HasSourceAt entity.
func (pnu *PackageNameUpdate) ClearHasSourceAt() *PackageNameUpdate {
	pnu.mutation.ClearHasSourceAt()
	return pnu
}

// RemoveHasSourceAtIDs removes the "has_source_at" edge to HasSourceAt entities by IDs.
func (pnu *PackageNameUpdate) RemoveHasSourceAtIDs(ids ...uuid.UUID) *PackageNameUpdate {
	pnu.mutation.RemoveHasSourceAtIDs(ids...)
	return pnu
}

// RemoveHasSourceAt removes "has_source_at" edges to HasSourceAt entities.
func (pnu *PackageNameUpdate) RemoveHasSourceAt(h ...*HasSourceAt) *PackageNameUpdate {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return pnu.RemoveHasSourceAtIDs(ids...)
}

// ClearCertification clears all "certification" edges to the Certification entity.
func (pnu *PackageNameUpdate) ClearCertification() *PackageNameUpdate {
	pnu.mutation.ClearCertification()
	return pnu
}

// RemoveCertificationIDs removes the "certification" edge to Certification entities by IDs.
func (pnu *PackageNameUpdate) RemoveCertificationIDs(ids ...uuid.UUID) *PackageNameUpdate {
	pnu.mutation.RemoveCertificationIDs(ids...)
	return pnu
}

// RemoveCertification removes "certification" edges to Certification entities.
func (pnu *PackageNameUpdate) RemoveCertification(c ...*Certification) *PackageNameUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pnu.RemoveCertificationIDs(ids...)
}

// ClearMetadata clears all "metadata" edges to the HasMetadata entity.
func (pnu *PackageNameUpdate) ClearMetadata() *PackageNameUpdate {
	pnu.mutation.ClearMetadata()
	return pnu
}

// RemoveMetadatumIDs removes the "metadata" edge to HasMetadata entities by IDs.
func (pnu *PackageNameUpdate) RemoveMetadatumIDs(ids ...uuid.UUID) *PackageNameUpdate {
	pnu.mutation.RemoveMetadatumIDs(ids...)
	return pnu
}

// RemoveMetadata removes "metadata" edges to HasMetadata entities.
func (pnu *PackageNameUpdate) RemoveMetadata(h ...*HasMetadata) *PackageNameUpdate {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return pnu.RemoveMetadatumIDs(ids...)
}

// ClearPoc clears all "poc" edges to the PointOfContact entity.
func (pnu *PackageNameUpdate) ClearPoc() *PackageNameUpdate {
	pnu.mutation.ClearPoc()
	return pnu
}

// RemovePocIDs removes the "poc" edge to PointOfContact entities by IDs.
func (pnu *PackageNameUpdate) RemovePocIDs(ids ...uuid.UUID) *PackageNameUpdate {
	pnu.mutation.RemovePocIDs(ids...)
	return pnu
}

// RemovePoc removes "poc" edges to PointOfContact entities.
func (pnu *PackageNameUpdate) RemovePoc(p ...*PointOfContact) *PackageNameUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pnu.RemovePocIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pnu *PackageNameUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pnu.sqlSave, pnu.mutation, pnu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pnu *PackageNameUpdate) SaveX(ctx context.Context) int {
	affected, err := pnu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pnu *PackageNameUpdate) Exec(ctx context.Context) error {
	_, err := pnu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pnu *PackageNameUpdate) ExecX(ctx context.Context) {
	if err := pnu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pnu *PackageNameUpdate) check() error {
	if v, ok := pnu.mutation.GetType(); ok {
		if err := packagename.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "PackageName.type": %w`, err)}
		}
	}
	if v, ok := pnu.mutation.Name(); ok {
		if err := packagename.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "PackageName.name": %w`, err)}
		}
	}
	return nil
}

func (pnu *PackageNameUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pnu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(packagename.Table, packagename.Columns, sqlgraph.NewFieldSpec(packagename.FieldID, field.TypeUUID))
	if ps := pnu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pnu.mutation.GetType(); ok {
		_spec.SetField(packagename.FieldType, field.TypeString, value)
	}
	if value, ok := pnu.mutation.Namespace(); ok {
		_spec.SetField(packagename.FieldNamespace, field.TypeString, value)
	}
	if value, ok := pnu.mutation.Name(); ok {
		_spec.SetField(packagename.FieldName, field.TypeString, value)
	}
	if pnu.mutation.VersionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   packagename.VersionsTable,
			Columns: []string{packagename.VersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnu.mutation.RemovedVersionsIDs(); len(nodes) > 0 && !pnu.mutation.VersionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   packagename.VersionsTable,
			Columns: []string{packagename.VersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnu.mutation.VersionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   packagename.VersionsTable,
			Columns: []string{packagename.VersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pnu.mutation.HasSourceAtCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   packagename.HasSourceAtTable,
			Columns: []string{packagename.HasSourceAtColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hassourceat.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnu.mutation.RemovedHasSourceAtIDs(); len(nodes) > 0 && !pnu.mutation.HasSourceAtCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   packagename.HasSourceAtTable,
			Columns: []string{packagename.HasSourceAtColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hassourceat.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnu.mutation.HasSourceAtIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   packagename.HasSourceAtTable,
			Columns: []string{packagename.HasSourceAtColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hassourceat.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pnu.mutation.CertificationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   packagename.CertificationTable,
			Columns: []string{packagename.CertificationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(certification.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnu.mutation.RemovedCertificationIDs(); len(nodes) > 0 && !pnu.mutation.CertificationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   packagename.CertificationTable,
			Columns: []string{packagename.CertificationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(certification.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnu.mutation.CertificationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   packagename.CertificationTable,
			Columns: []string{packagename.CertificationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(certification.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pnu.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   packagename.MetadataTable,
			Columns: []string{packagename.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hasmetadata.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnu.mutation.RemovedMetadataIDs(); len(nodes) > 0 && !pnu.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   packagename.MetadataTable,
			Columns: []string{packagename.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hasmetadata.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnu.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   packagename.MetadataTable,
			Columns: []string{packagename.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hasmetadata.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pnu.mutation.PocCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   packagename.PocTable,
			Columns: []string{packagename.PocColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pointofcontact.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnu.mutation.RemovedPocIDs(); len(nodes) > 0 && !pnu.mutation.PocCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   packagename.PocTable,
			Columns: []string{packagename.PocColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pointofcontact.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnu.mutation.PocIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   packagename.PocTable,
			Columns: []string{packagename.PocColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pointofcontact.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pnu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{packagename.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pnu.mutation.done = true
	return n, nil
}

// PackageNameUpdateOne is the builder for updating a single PackageName entity.
type PackageNameUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PackageNameMutation
}

// SetType sets the "type" field.
func (pnuo *PackageNameUpdateOne) SetType(s string) *PackageNameUpdateOne {
	pnuo.mutation.SetType(s)
	return pnuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (pnuo *PackageNameUpdateOne) SetNillableType(s *string) *PackageNameUpdateOne {
	if s != nil {
		pnuo.SetType(*s)
	}
	return pnuo
}

// SetNamespace sets the "namespace" field.
func (pnuo *PackageNameUpdateOne) SetNamespace(s string) *PackageNameUpdateOne {
	pnuo.mutation.SetNamespace(s)
	return pnuo
}

// SetNillableNamespace sets the "namespace" field if the given value is not nil.
func (pnuo *PackageNameUpdateOne) SetNillableNamespace(s *string) *PackageNameUpdateOne {
	if s != nil {
		pnuo.SetNamespace(*s)
	}
	return pnuo
}

// SetName sets the "name" field.
func (pnuo *PackageNameUpdateOne) SetName(s string) *PackageNameUpdateOne {
	pnuo.mutation.SetName(s)
	return pnuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pnuo *PackageNameUpdateOne) SetNillableName(s *string) *PackageNameUpdateOne {
	if s != nil {
		pnuo.SetName(*s)
	}
	return pnuo
}

// AddVersionIDs adds the "versions" edge to the PackageVersion entity by IDs.
func (pnuo *PackageNameUpdateOne) AddVersionIDs(ids ...uuid.UUID) *PackageNameUpdateOne {
	pnuo.mutation.AddVersionIDs(ids...)
	return pnuo
}

// AddVersions adds the "versions" edges to the PackageVersion entity.
func (pnuo *PackageNameUpdateOne) AddVersions(p ...*PackageVersion) *PackageNameUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pnuo.AddVersionIDs(ids...)
}

// AddHasSourceAtIDs adds the "has_source_at" edge to the HasSourceAt entity by IDs.
func (pnuo *PackageNameUpdateOne) AddHasSourceAtIDs(ids ...uuid.UUID) *PackageNameUpdateOne {
	pnuo.mutation.AddHasSourceAtIDs(ids...)
	return pnuo
}

// AddHasSourceAt adds the "has_source_at" edges to the HasSourceAt entity.
func (pnuo *PackageNameUpdateOne) AddHasSourceAt(h ...*HasSourceAt) *PackageNameUpdateOne {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return pnuo.AddHasSourceAtIDs(ids...)
}

// AddCertificationIDs adds the "certification" edge to the Certification entity by IDs.
func (pnuo *PackageNameUpdateOne) AddCertificationIDs(ids ...uuid.UUID) *PackageNameUpdateOne {
	pnuo.mutation.AddCertificationIDs(ids...)
	return pnuo
}

// AddCertification adds the "certification" edges to the Certification entity.
func (pnuo *PackageNameUpdateOne) AddCertification(c ...*Certification) *PackageNameUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pnuo.AddCertificationIDs(ids...)
}

// AddMetadatumIDs adds the "metadata" edge to the HasMetadata entity by IDs.
func (pnuo *PackageNameUpdateOne) AddMetadatumIDs(ids ...uuid.UUID) *PackageNameUpdateOne {
	pnuo.mutation.AddMetadatumIDs(ids...)
	return pnuo
}

// AddMetadata adds the "metadata" edges to the HasMetadata entity.
func (pnuo *PackageNameUpdateOne) AddMetadata(h ...*HasMetadata) *PackageNameUpdateOne {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return pnuo.AddMetadatumIDs(ids...)
}

// AddPocIDs adds the "poc" edge to the PointOfContact entity by IDs.
func (pnuo *PackageNameUpdateOne) AddPocIDs(ids ...uuid.UUID) *PackageNameUpdateOne {
	pnuo.mutation.AddPocIDs(ids...)
	return pnuo
}

// AddPoc adds the "poc" edges to the PointOfContact entity.
func (pnuo *PackageNameUpdateOne) AddPoc(p ...*PointOfContact) *PackageNameUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pnuo.AddPocIDs(ids...)
}

// Mutation returns the PackageNameMutation object of the builder.
func (pnuo *PackageNameUpdateOne) Mutation() *PackageNameMutation {
	return pnuo.mutation
}

// ClearVersions clears all "versions" edges to the PackageVersion entity.
func (pnuo *PackageNameUpdateOne) ClearVersions() *PackageNameUpdateOne {
	pnuo.mutation.ClearVersions()
	return pnuo
}

// RemoveVersionIDs removes the "versions" edge to PackageVersion entities by IDs.
func (pnuo *PackageNameUpdateOne) RemoveVersionIDs(ids ...uuid.UUID) *PackageNameUpdateOne {
	pnuo.mutation.RemoveVersionIDs(ids...)
	return pnuo
}

// RemoveVersions removes "versions" edges to PackageVersion entities.
func (pnuo *PackageNameUpdateOne) RemoveVersions(p ...*PackageVersion) *PackageNameUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pnuo.RemoveVersionIDs(ids...)
}

// ClearHasSourceAt clears all "has_source_at" edges to the HasSourceAt entity.
func (pnuo *PackageNameUpdateOne) ClearHasSourceAt() *PackageNameUpdateOne {
	pnuo.mutation.ClearHasSourceAt()
	return pnuo
}

// RemoveHasSourceAtIDs removes the "has_source_at" edge to HasSourceAt entities by IDs.
func (pnuo *PackageNameUpdateOne) RemoveHasSourceAtIDs(ids ...uuid.UUID) *PackageNameUpdateOne {
	pnuo.mutation.RemoveHasSourceAtIDs(ids...)
	return pnuo
}

// RemoveHasSourceAt removes "has_source_at" edges to HasSourceAt entities.
func (pnuo *PackageNameUpdateOne) RemoveHasSourceAt(h ...*HasSourceAt) *PackageNameUpdateOne {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return pnuo.RemoveHasSourceAtIDs(ids...)
}

// ClearCertification clears all "certification" edges to the Certification entity.
func (pnuo *PackageNameUpdateOne) ClearCertification() *PackageNameUpdateOne {
	pnuo.mutation.ClearCertification()
	return pnuo
}

// RemoveCertificationIDs removes the "certification" edge to Certification entities by IDs.
func (pnuo *PackageNameUpdateOne) RemoveCertificationIDs(ids ...uuid.UUID) *PackageNameUpdateOne {
	pnuo.mutation.RemoveCertificationIDs(ids...)
	return pnuo
}

// RemoveCertification removes "certification" edges to Certification entities.
func (pnuo *PackageNameUpdateOne) RemoveCertification(c ...*Certification) *PackageNameUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pnuo.RemoveCertificationIDs(ids...)
}

// ClearMetadata clears all "metadata" edges to the HasMetadata entity.
func (pnuo *PackageNameUpdateOne) ClearMetadata() *PackageNameUpdateOne {
	pnuo.mutation.ClearMetadata()
	return pnuo
}

// RemoveMetadatumIDs removes the "metadata" edge to HasMetadata entities by IDs.
func (pnuo *PackageNameUpdateOne) RemoveMetadatumIDs(ids ...uuid.UUID) *PackageNameUpdateOne {
	pnuo.mutation.RemoveMetadatumIDs(ids...)
	return pnuo
}

// RemoveMetadata removes "metadata" edges to HasMetadata entities.
func (pnuo *PackageNameUpdateOne) RemoveMetadata(h ...*HasMetadata) *PackageNameUpdateOne {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return pnuo.RemoveMetadatumIDs(ids...)
}

// ClearPoc clears all "poc" edges to the PointOfContact entity.
func (pnuo *PackageNameUpdateOne) ClearPoc() *PackageNameUpdateOne {
	pnuo.mutation.ClearPoc()
	return pnuo
}

// RemovePocIDs removes the "poc" edge to PointOfContact entities by IDs.
func (pnuo *PackageNameUpdateOne) RemovePocIDs(ids ...uuid.UUID) *PackageNameUpdateOne {
	pnuo.mutation.RemovePocIDs(ids...)
	return pnuo
}

// RemovePoc removes "poc" edges to PointOfContact entities.
func (pnuo *PackageNameUpdateOne) RemovePoc(p ...*PointOfContact) *PackageNameUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pnuo.RemovePocIDs(ids...)
}

// Where appends a list predicates to the PackageNameUpdate builder.
func (pnuo *PackageNameUpdateOne) Where(ps ...predicate.PackageName) *PackageNameUpdateOne {
	pnuo.mutation.Where(ps...)
	return pnuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pnuo *PackageNameUpdateOne) Select(field string, fields ...string) *PackageNameUpdateOne {
	pnuo.fields = append([]string{field}, fields...)
	return pnuo
}

// Save executes the query and returns the updated PackageName entity.
func (pnuo *PackageNameUpdateOne) Save(ctx context.Context) (*PackageName, error) {
	return withHooks(ctx, pnuo.sqlSave, pnuo.mutation, pnuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pnuo *PackageNameUpdateOne) SaveX(ctx context.Context) *PackageName {
	node, err := pnuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pnuo *PackageNameUpdateOne) Exec(ctx context.Context) error {
	_, err := pnuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pnuo *PackageNameUpdateOne) ExecX(ctx context.Context) {
	if err := pnuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pnuo *PackageNameUpdateOne) check() error {
	if v, ok := pnuo.mutation.GetType(); ok {
		if err := packagename.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "PackageName.type": %w`, err)}
		}
	}
	if v, ok := pnuo.mutation.Name(); ok {
		if err := packagename.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "PackageName.name": %w`, err)}
		}
	}
	return nil
}

func (pnuo *PackageNameUpdateOne) sqlSave(ctx context.Context) (_node *PackageName, err error) {
	if err := pnuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(packagename.Table, packagename.Columns, sqlgraph.NewFieldSpec(packagename.FieldID, field.TypeUUID))
	id, ok := pnuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PackageName.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pnuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, packagename.FieldID)
		for _, f := range fields {
			if !packagename.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != packagename.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pnuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pnuo.mutation.GetType(); ok {
		_spec.SetField(packagename.FieldType, field.TypeString, value)
	}
	if value, ok := pnuo.mutation.Namespace(); ok {
		_spec.SetField(packagename.FieldNamespace, field.TypeString, value)
	}
	if value, ok := pnuo.mutation.Name(); ok {
		_spec.SetField(packagename.FieldName, field.TypeString, value)
	}
	if pnuo.mutation.VersionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   packagename.VersionsTable,
			Columns: []string{packagename.VersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnuo.mutation.RemovedVersionsIDs(); len(nodes) > 0 && !pnuo.mutation.VersionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   packagename.VersionsTable,
			Columns: []string{packagename.VersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnuo.mutation.VersionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   packagename.VersionsTable,
			Columns: []string{packagename.VersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pnuo.mutation.HasSourceAtCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   packagename.HasSourceAtTable,
			Columns: []string{packagename.HasSourceAtColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hassourceat.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnuo.mutation.RemovedHasSourceAtIDs(); len(nodes) > 0 && !pnuo.mutation.HasSourceAtCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   packagename.HasSourceAtTable,
			Columns: []string{packagename.HasSourceAtColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hassourceat.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnuo.mutation.HasSourceAtIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   packagename.HasSourceAtTable,
			Columns: []string{packagename.HasSourceAtColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hassourceat.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pnuo.mutation.CertificationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   packagename.CertificationTable,
			Columns: []string{packagename.CertificationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(certification.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnuo.mutation.RemovedCertificationIDs(); len(nodes) > 0 && !pnuo.mutation.CertificationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   packagename.CertificationTable,
			Columns: []string{packagename.CertificationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(certification.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnuo.mutation.CertificationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   packagename.CertificationTable,
			Columns: []string{packagename.CertificationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(certification.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pnuo.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   packagename.MetadataTable,
			Columns: []string{packagename.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hasmetadata.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnuo.mutation.RemovedMetadataIDs(); len(nodes) > 0 && !pnuo.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   packagename.MetadataTable,
			Columns: []string{packagename.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hasmetadata.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnuo.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   packagename.MetadataTable,
			Columns: []string{packagename.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hasmetadata.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pnuo.mutation.PocCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   packagename.PocTable,
			Columns: []string{packagename.PocColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pointofcontact.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnuo.mutation.RemovedPocIDs(); len(nodes) > 0 && !pnuo.mutation.PocCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   packagename.PocTable,
			Columns: []string{packagename.PocColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pointofcontact.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pnuo.mutation.PocIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   packagename.PocTable,
			Columns: []string{packagename.PocColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pointofcontact.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PackageName{config: pnuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pnuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{packagename.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pnuo.mutation.done = true
	return _node, nil
}

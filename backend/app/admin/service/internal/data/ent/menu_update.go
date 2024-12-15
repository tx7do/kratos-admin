// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	servicev1 "kratos-admin/api/gen/go/system/service/v1"
	"kratos-admin/app/admin/service/internal/data/ent/menu"
	"kratos-admin/app/admin/service/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MenuUpdate is the builder for updating Menu entities.
type MenuUpdate struct {
	config
	hooks     []Hook
	mutation  *MenuMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the MenuUpdate builder.
func (mu *MenuUpdate) Where(ps ...predicate.Menu) *MenuUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetStatus sets the "status" field.
func (mu *MenuUpdate) SetStatus(m menu.Status) *MenuUpdate {
	mu.mutation.SetStatus(m)
	return mu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableStatus(m *menu.Status) *MenuUpdate {
	if m != nil {
		mu.SetStatus(*m)
	}
	return mu
}

// ClearStatus clears the value of the "status" field.
func (mu *MenuUpdate) ClearStatus() *MenuUpdate {
	mu.mutation.ClearStatus()
	return mu
}

// SetUpdateTime sets the "update_time" field.
func (mu *MenuUpdate) SetUpdateTime(t time.Time) *MenuUpdate {
	mu.mutation.SetUpdateTime(t)
	return mu
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableUpdateTime(t *time.Time) *MenuUpdate {
	if t != nil {
		mu.SetUpdateTime(*t)
	}
	return mu
}

// ClearUpdateTime clears the value of the "update_time" field.
func (mu *MenuUpdate) ClearUpdateTime() *MenuUpdate {
	mu.mutation.ClearUpdateTime()
	return mu
}

// SetDeleteTime sets the "delete_time" field.
func (mu *MenuUpdate) SetDeleteTime(t time.Time) *MenuUpdate {
	mu.mutation.SetDeleteTime(t)
	return mu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableDeleteTime(t *time.Time) *MenuUpdate {
	if t != nil {
		mu.SetDeleteTime(*t)
	}
	return mu
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (mu *MenuUpdate) ClearDeleteTime() *MenuUpdate {
	mu.mutation.ClearDeleteTime()
	return mu
}

// SetCreateBy sets the "create_by" field.
func (mu *MenuUpdate) SetCreateBy(u uint32) *MenuUpdate {
	mu.mutation.ResetCreateBy()
	mu.mutation.SetCreateBy(u)
	return mu
}

// SetNillableCreateBy sets the "create_by" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableCreateBy(u *uint32) *MenuUpdate {
	if u != nil {
		mu.SetCreateBy(*u)
	}
	return mu
}

// AddCreateBy adds u to the "create_by" field.
func (mu *MenuUpdate) AddCreateBy(u int32) *MenuUpdate {
	mu.mutation.AddCreateBy(u)
	return mu
}

// ClearCreateBy clears the value of the "create_by" field.
func (mu *MenuUpdate) ClearCreateBy() *MenuUpdate {
	mu.mutation.ClearCreateBy()
	return mu
}

// SetParentID sets the "parent_id" field.
func (mu *MenuUpdate) SetParentID(i int32) *MenuUpdate {
	mu.mutation.SetParentID(i)
	return mu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableParentID(i *int32) *MenuUpdate {
	if i != nil {
		mu.SetParentID(*i)
	}
	return mu
}

// ClearParentID clears the value of the "parent_id" field.
func (mu *MenuUpdate) ClearParentID() *MenuUpdate {
	mu.mutation.ClearParentID()
	return mu
}

// SetType sets the "type" field.
func (mu *MenuUpdate) SetType(m menu.Type) *MenuUpdate {
	mu.mutation.SetType(m)
	return mu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableType(m *menu.Type) *MenuUpdate {
	if m != nil {
		mu.SetType(*m)
	}
	return mu
}

// ClearType clears the value of the "type" field.
func (mu *MenuUpdate) ClearType() *MenuUpdate {
	mu.mutation.ClearType()
	return mu
}

// SetPath sets the "path" field.
func (mu *MenuUpdate) SetPath(s string) *MenuUpdate {
	mu.mutation.SetPath(s)
	return mu
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (mu *MenuUpdate) SetNillablePath(s *string) *MenuUpdate {
	if s != nil {
		mu.SetPath(*s)
	}
	return mu
}

// ClearPath clears the value of the "path" field.
func (mu *MenuUpdate) ClearPath() *MenuUpdate {
	mu.mutation.ClearPath()
	return mu
}

// SetRedirect sets the "redirect" field.
func (mu *MenuUpdate) SetRedirect(s string) *MenuUpdate {
	mu.mutation.SetRedirect(s)
	return mu
}

// SetNillableRedirect sets the "redirect" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableRedirect(s *string) *MenuUpdate {
	if s != nil {
		mu.SetRedirect(*s)
	}
	return mu
}

// ClearRedirect clears the value of the "redirect" field.
func (mu *MenuUpdate) ClearRedirect() *MenuUpdate {
	mu.mutation.ClearRedirect()
	return mu
}

// SetAlias sets the "alias" field.
func (mu *MenuUpdate) SetAlias(s string) *MenuUpdate {
	mu.mutation.SetAlias(s)
	return mu
}

// SetNillableAlias sets the "alias" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableAlias(s *string) *MenuUpdate {
	if s != nil {
		mu.SetAlias(*s)
	}
	return mu
}

// ClearAlias clears the value of the "alias" field.
func (mu *MenuUpdate) ClearAlias() *MenuUpdate {
	mu.mutation.ClearAlias()
	return mu
}

// SetName sets the "name" field.
func (mu *MenuUpdate) SetName(s string) *MenuUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableName(s *string) *MenuUpdate {
	if s != nil {
		mu.SetName(*s)
	}
	return mu
}

// ClearName clears the value of the "name" field.
func (mu *MenuUpdate) ClearName() *MenuUpdate {
	mu.mutation.ClearName()
	return mu
}

// SetComponent sets the "component" field.
func (mu *MenuUpdate) SetComponent(s string) *MenuUpdate {
	mu.mutation.SetComponent(s)
	return mu
}

// SetNillableComponent sets the "component" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableComponent(s *string) *MenuUpdate {
	if s != nil {
		mu.SetComponent(*s)
	}
	return mu
}

// ClearComponent clears the value of the "component" field.
func (mu *MenuUpdate) ClearComponent() *MenuUpdate {
	mu.mutation.ClearComponent()
	return mu
}

// SetMeta sets the "meta" field.
func (mu *MenuUpdate) SetMeta(sm *servicev1.RouteMeta) *MenuUpdate {
	mu.mutation.SetMeta(sm)
	return mu
}

// ClearMeta clears the value of the "meta" field.
func (mu *MenuUpdate) ClearMeta() *MenuUpdate {
	mu.mutation.ClearMeta()
	return mu
}

// SetParent sets the "parent" edge to the Menu entity.
func (mu *MenuUpdate) SetParent(m *Menu) *MenuUpdate {
	return mu.SetParentID(m.ID)
}

// AddChildIDs adds the "children" edge to the Menu entity by IDs.
func (mu *MenuUpdate) AddChildIDs(ids ...int32) *MenuUpdate {
	mu.mutation.AddChildIDs(ids...)
	return mu
}

// AddChildren adds the "children" edges to the Menu entity.
func (mu *MenuUpdate) AddChildren(m ...*Menu) *MenuUpdate {
	ids := make([]int32, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.AddChildIDs(ids...)
}

// Mutation returns the MenuMutation object of the builder.
func (mu *MenuUpdate) Mutation() *MenuMutation {
	return mu.mutation
}

// ClearParent clears the "parent" edge to the Menu entity.
func (mu *MenuUpdate) ClearParent() *MenuUpdate {
	mu.mutation.ClearParent()
	return mu
}

// ClearChildren clears all "children" edges to the Menu entity.
func (mu *MenuUpdate) ClearChildren() *MenuUpdate {
	mu.mutation.ClearChildren()
	return mu
}

// RemoveChildIDs removes the "children" edge to Menu entities by IDs.
func (mu *MenuUpdate) RemoveChildIDs(ids ...int32) *MenuUpdate {
	mu.mutation.RemoveChildIDs(ids...)
	return mu
}

// RemoveChildren removes "children" edges to Menu entities.
func (mu *MenuUpdate) RemoveChildren(m ...*Menu) *MenuUpdate {
	ids := make([]int32, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.RemoveChildIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MenuUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MenuUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MenuUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MenuUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MenuUpdate) check() error {
	if v, ok := mu.mutation.Status(); ok {
		if err := menu.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Menu.status": %w`, err)}
		}
	}
	if v, ok := mu.mutation.GetType(); ok {
		if err := menu.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Menu.type": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Meta(); ok {
		if err := v.Validate(); err != nil {
			return &ValidationError{Name: "meta", err: fmt.Errorf(`ent: validator failed for field "Menu.meta": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mu *MenuUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MenuUpdate {
	mu.modifiers = append(mu.modifiers, modifiers...)
	return mu
}

func (mu *MenuUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(menu.Table, menu.Columns, sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt32))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Status(); ok {
		_spec.SetField(menu.FieldStatus, field.TypeEnum, value)
	}
	if mu.mutation.StatusCleared() {
		_spec.ClearField(menu.FieldStatus, field.TypeEnum)
	}
	if mu.mutation.CreateTimeCleared() {
		_spec.ClearField(menu.FieldCreateTime, field.TypeTime)
	}
	if value, ok := mu.mutation.UpdateTime(); ok {
		_spec.SetField(menu.FieldUpdateTime, field.TypeTime, value)
	}
	if mu.mutation.UpdateTimeCleared() {
		_spec.ClearField(menu.FieldUpdateTime, field.TypeTime)
	}
	if value, ok := mu.mutation.DeleteTime(); ok {
		_spec.SetField(menu.FieldDeleteTime, field.TypeTime, value)
	}
	if mu.mutation.DeleteTimeCleared() {
		_spec.ClearField(menu.FieldDeleteTime, field.TypeTime)
	}
	if value, ok := mu.mutation.CreateBy(); ok {
		_spec.SetField(menu.FieldCreateBy, field.TypeUint32, value)
	}
	if value, ok := mu.mutation.AddedCreateBy(); ok {
		_spec.AddField(menu.FieldCreateBy, field.TypeUint32, value)
	}
	if mu.mutation.CreateByCleared() {
		_spec.ClearField(menu.FieldCreateBy, field.TypeUint32)
	}
	if value, ok := mu.mutation.GetType(); ok {
		_spec.SetField(menu.FieldType, field.TypeEnum, value)
	}
	if mu.mutation.TypeCleared() {
		_spec.ClearField(menu.FieldType, field.TypeEnum)
	}
	if value, ok := mu.mutation.Path(); ok {
		_spec.SetField(menu.FieldPath, field.TypeString, value)
	}
	if mu.mutation.PathCleared() {
		_spec.ClearField(menu.FieldPath, field.TypeString)
	}
	if value, ok := mu.mutation.Redirect(); ok {
		_spec.SetField(menu.FieldRedirect, field.TypeString, value)
	}
	if mu.mutation.RedirectCleared() {
		_spec.ClearField(menu.FieldRedirect, field.TypeString)
	}
	if value, ok := mu.mutation.Alias(); ok {
		_spec.SetField(menu.FieldAlias, field.TypeString, value)
	}
	if mu.mutation.AliasCleared() {
		_spec.ClearField(menu.FieldAlias, field.TypeString)
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.SetField(menu.FieldName, field.TypeString, value)
	}
	if mu.mutation.NameCleared() {
		_spec.ClearField(menu.FieldName, field.TypeString)
	}
	if value, ok := mu.mutation.Component(); ok {
		_spec.SetField(menu.FieldComponent, field.TypeString, value)
	}
	if mu.mutation.ComponentCleared() {
		_spec.ClearField(menu.FieldComponent, field.TypeString)
	}
	if value, ok := mu.mutation.Meta(); ok {
		_spec.SetField(menu.FieldMeta, field.TypeJSON, value)
	}
	if mu.mutation.MetaCleared() {
		_spec.ClearField(menu.FieldMeta, field.TypeJSON)
	}
	if mu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menu.ParentTable,
			Columns: []string{menu.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menu.ParentTable,
			Columns: []string{menu.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.ChildrenTable,
			Columns: []string{menu.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !mu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.ChildrenTable,
			Columns: []string{menu.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.ChildrenTable,
			Columns: []string{menu.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(mu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{menu.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MenuUpdateOne is the builder for updating a single Menu entity.
type MenuUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *MenuMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetStatus sets the "status" field.
func (muo *MenuUpdateOne) SetStatus(m menu.Status) *MenuUpdateOne {
	muo.mutation.SetStatus(m)
	return muo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableStatus(m *menu.Status) *MenuUpdateOne {
	if m != nil {
		muo.SetStatus(*m)
	}
	return muo
}

// ClearStatus clears the value of the "status" field.
func (muo *MenuUpdateOne) ClearStatus() *MenuUpdateOne {
	muo.mutation.ClearStatus()
	return muo
}

// SetUpdateTime sets the "update_time" field.
func (muo *MenuUpdateOne) SetUpdateTime(t time.Time) *MenuUpdateOne {
	muo.mutation.SetUpdateTime(t)
	return muo
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableUpdateTime(t *time.Time) *MenuUpdateOne {
	if t != nil {
		muo.SetUpdateTime(*t)
	}
	return muo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (muo *MenuUpdateOne) ClearUpdateTime() *MenuUpdateOne {
	muo.mutation.ClearUpdateTime()
	return muo
}

// SetDeleteTime sets the "delete_time" field.
func (muo *MenuUpdateOne) SetDeleteTime(t time.Time) *MenuUpdateOne {
	muo.mutation.SetDeleteTime(t)
	return muo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableDeleteTime(t *time.Time) *MenuUpdateOne {
	if t != nil {
		muo.SetDeleteTime(*t)
	}
	return muo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (muo *MenuUpdateOne) ClearDeleteTime() *MenuUpdateOne {
	muo.mutation.ClearDeleteTime()
	return muo
}

// SetCreateBy sets the "create_by" field.
func (muo *MenuUpdateOne) SetCreateBy(u uint32) *MenuUpdateOne {
	muo.mutation.ResetCreateBy()
	muo.mutation.SetCreateBy(u)
	return muo
}

// SetNillableCreateBy sets the "create_by" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableCreateBy(u *uint32) *MenuUpdateOne {
	if u != nil {
		muo.SetCreateBy(*u)
	}
	return muo
}

// AddCreateBy adds u to the "create_by" field.
func (muo *MenuUpdateOne) AddCreateBy(u int32) *MenuUpdateOne {
	muo.mutation.AddCreateBy(u)
	return muo
}

// ClearCreateBy clears the value of the "create_by" field.
func (muo *MenuUpdateOne) ClearCreateBy() *MenuUpdateOne {
	muo.mutation.ClearCreateBy()
	return muo
}

// SetParentID sets the "parent_id" field.
func (muo *MenuUpdateOne) SetParentID(i int32) *MenuUpdateOne {
	muo.mutation.SetParentID(i)
	return muo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableParentID(i *int32) *MenuUpdateOne {
	if i != nil {
		muo.SetParentID(*i)
	}
	return muo
}

// ClearParentID clears the value of the "parent_id" field.
func (muo *MenuUpdateOne) ClearParentID() *MenuUpdateOne {
	muo.mutation.ClearParentID()
	return muo
}

// SetType sets the "type" field.
func (muo *MenuUpdateOne) SetType(m menu.Type) *MenuUpdateOne {
	muo.mutation.SetType(m)
	return muo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableType(m *menu.Type) *MenuUpdateOne {
	if m != nil {
		muo.SetType(*m)
	}
	return muo
}

// ClearType clears the value of the "type" field.
func (muo *MenuUpdateOne) ClearType() *MenuUpdateOne {
	muo.mutation.ClearType()
	return muo
}

// SetPath sets the "path" field.
func (muo *MenuUpdateOne) SetPath(s string) *MenuUpdateOne {
	muo.mutation.SetPath(s)
	return muo
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillablePath(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetPath(*s)
	}
	return muo
}

// ClearPath clears the value of the "path" field.
func (muo *MenuUpdateOne) ClearPath() *MenuUpdateOne {
	muo.mutation.ClearPath()
	return muo
}

// SetRedirect sets the "redirect" field.
func (muo *MenuUpdateOne) SetRedirect(s string) *MenuUpdateOne {
	muo.mutation.SetRedirect(s)
	return muo
}

// SetNillableRedirect sets the "redirect" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableRedirect(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetRedirect(*s)
	}
	return muo
}

// ClearRedirect clears the value of the "redirect" field.
func (muo *MenuUpdateOne) ClearRedirect() *MenuUpdateOne {
	muo.mutation.ClearRedirect()
	return muo
}

// SetAlias sets the "alias" field.
func (muo *MenuUpdateOne) SetAlias(s string) *MenuUpdateOne {
	muo.mutation.SetAlias(s)
	return muo
}

// SetNillableAlias sets the "alias" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableAlias(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetAlias(*s)
	}
	return muo
}

// ClearAlias clears the value of the "alias" field.
func (muo *MenuUpdateOne) ClearAlias() *MenuUpdateOne {
	muo.mutation.ClearAlias()
	return muo
}

// SetName sets the "name" field.
func (muo *MenuUpdateOne) SetName(s string) *MenuUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableName(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetName(*s)
	}
	return muo
}

// ClearName clears the value of the "name" field.
func (muo *MenuUpdateOne) ClearName() *MenuUpdateOne {
	muo.mutation.ClearName()
	return muo
}

// SetComponent sets the "component" field.
func (muo *MenuUpdateOne) SetComponent(s string) *MenuUpdateOne {
	muo.mutation.SetComponent(s)
	return muo
}

// SetNillableComponent sets the "component" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableComponent(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetComponent(*s)
	}
	return muo
}

// ClearComponent clears the value of the "component" field.
func (muo *MenuUpdateOne) ClearComponent() *MenuUpdateOne {
	muo.mutation.ClearComponent()
	return muo
}

// SetMeta sets the "meta" field.
func (muo *MenuUpdateOne) SetMeta(sm *servicev1.RouteMeta) *MenuUpdateOne {
	muo.mutation.SetMeta(sm)
	return muo
}

// ClearMeta clears the value of the "meta" field.
func (muo *MenuUpdateOne) ClearMeta() *MenuUpdateOne {
	muo.mutation.ClearMeta()
	return muo
}

// SetParent sets the "parent" edge to the Menu entity.
func (muo *MenuUpdateOne) SetParent(m *Menu) *MenuUpdateOne {
	return muo.SetParentID(m.ID)
}

// AddChildIDs adds the "children" edge to the Menu entity by IDs.
func (muo *MenuUpdateOne) AddChildIDs(ids ...int32) *MenuUpdateOne {
	muo.mutation.AddChildIDs(ids...)
	return muo
}

// AddChildren adds the "children" edges to the Menu entity.
func (muo *MenuUpdateOne) AddChildren(m ...*Menu) *MenuUpdateOne {
	ids := make([]int32, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.AddChildIDs(ids...)
}

// Mutation returns the MenuMutation object of the builder.
func (muo *MenuUpdateOne) Mutation() *MenuMutation {
	return muo.mutation
}

// ClearParent clears the "parent" edge to the Menu entity.
func (muo *MenuUpdateOne) ClearParent() *MenuUpdateOne {
	muo.mutation.ClearParent()
	return muo
}

// ClearChildren clears all "children" edges to the Menu entity.
func (muo *MenuUpdateOne) ClearChildren() *MenuUpdateOne {
	muo.mutation.ClearChildren()
	return muo
}

// RemoveChildIDs removes the "children" edge to Menu entities by IDs.
func (muo *MenuUpdateOne) RemoveChildIDs(ids ...int32) *MenuUpdateOne {
	muo.mutation.RemoveChildIDs(ids...)
	return muo
}

// RemoveChildren removes "children" edges to Menu entities.
func (muo *MenuUpdateOne) RemoveChildren(m ...*Menu) *MenuUpdateOne {
	ids := make([]int32, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.RemoveChildIDs(ids...)
}

// Where appends a list predicates to the MenuUpdate builder.
func (muo *MenuUpdateOne) Where(ps ...predicate.Menu) *MenuUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MenuUpdateOne) Select(field string, fields ...string) *MenuUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Menu entity.
func (muo *MenuUpdateOne) Save(ctx context.Context) (*Menu, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MenuUpdateOne) SaveX(ctx context.Context) *Menu {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MenuUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MenuUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MenuUpdateOne) check() error {
	if v, ok := muo.mutation.Status(); ok {
		if err := menu.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Menu.status": %w`, err)}
		}
	}
	if v, ok := muo.mutation.GetType(); ok {
		if err := menu.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Menu.type": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Meta(); ok {
		if err := v.Validate(); err != nil {
			return &ValidationError{Name: "meta", err: fmt.Errorf(`ent: validator failed for field "Menu.meta": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (muo *MenuUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MenuUpdateOne {
	muo.modifiers = append(muo.modifiers, modifiers...)
	return muo
}

func (muo *MenuUpdateOne) sqlSave(ctx context.Context) (_node *Menu, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(menu.Table, menu.Columns, sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt32))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Menu.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, menu.FieldID)
		for _, f := range fields {
			if !menu.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != menu.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Status(); ok {
		_spec.SetField(menu.FieldStatus, field.TypeEnum, value)
	}
	if muo.mutation.StatusCleared() {
		_spec.ClearField(menu.FieldStatus, field.TypeEnum)
	}
	if muo.mutation.CreateTimeCleared() {
		_spec.ClearField(menu.FieldCreateTime, field.TypeTime)
	}
	if value, ok := muo.mutation.UpdateTime(); ok {
		_spec.SetField(menu.FieldUpdateTime, field.TypeTime, value)
	}
	if muo.mutation.UpdateTimeCleared() {
		_spec.ClearField(menu.FieldUpdateTime, field.TypeTime)
	}
	if value, ok := muo.mutation.DeleteTime(); ok {
		_spec.SetField(menu.FieldDeleteTime, field.TypeTime, value)
	}
	if muo.mutation.DeleteTimeCleared() {
		_spec.ClearField(menu.FieldDeleteTime, field.TypeTime)
	}
	if value, ok := muo.mutation.CreateBy(); ok {
		_spec.SetField(menu.FieldCreateBy, field.TypeUint32, value)
	}
	if value, ok := muo.mutation.AddedCreateBy(); ok {
		_spec.AddField(menu.FieldCreateBy, field.TypeUint32, value)
	}
	if muo.mutation.CreateByCleared() {
		_spec.ClearField(menu.FieldCreateBy, field.TypeUint32)
	}
	if value, ok := muo.mutation.GetType(); ok {
		_spec.SetField(menu.FieldType, field.TypeEnum, value)
	}
	if muo.mutation.TypeCleared() {
		_spec.ClearField(menu.FieldType, field.TypeEnum)
	}
	if value, ok := muo.mutation.Path(); ok {
		_spec.SetField(menu.FieldPath, field.TypeString, value)
	}
	if muo.mutation.PathCleared() {
		_spec.ClearField(menu.FieldPath, field.TypeString)
	}
	if value, ok := muo.mutation.Redirect(); ok {
		_spec.SetField(menu.FieldRedirect, field.TypeString, value)
	}
	if muo.mutation.RedirectCleared() {
		_spec.ClearField(menu.FieldRedirect, field.TypeString)
	}
	if value, ok := muo.mutation.Alias(); ok {
		_spec.SetField(menu.FieldAlias, field.TypeString, value)
	}
	if muo.mutation.AliasCleared() {
		_spec.ClearField(menu.FieldAlias, field.TypeString)
	}
	if value, ok := muo.mutation.Name(); ok {
		_spec.SetField(menu.FieldName, field.TypeString, value)
	}
	if muo.mutation.NameCleared() {
		_spec.ClearField(menu.FieldName, field.TypeString)
	}
	if value, ok := muo.mutation.Component(); ok {
		_spec.SetField(menu.FieldComponent, field.TypeString, value)
	}
	if muo.mutation.ComponentCleared() {
		_spec.ClearField(menu.FieldComponent, field.TypeString)
	}
	if value, ok := muo.mutation.Meta(); ok {
		_spec.SetField(menu.FieldMeta, field.TypeJSON, value)
	}
	if muo.mutation.MetaCleared() {
		_spec.ClearField(menu.FieldMeta, field.TypeJSON)
	}
	if muo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menu.ParentTable,
			Columns: []string{menu.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menu.ParentTable,
			Columns: []string{menu.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.ChildrenTable,
			Columns: []string{menu.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !muo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.ChildrenTable,
			Columns: []string{menu.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.ChildrenTable,
			Columns: []string{menu.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(muo.modifiers...)
	_node = &Menu{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{menu.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
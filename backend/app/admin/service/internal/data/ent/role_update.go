// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratos-admin/app/admin/service/internal/data/ent/predicate"
	"kratos-admin/app/admin/service/internal/data/ent/role"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoleUpdate is the builder for updating Role entities.
type RoleUpdate struct {
	config
	hooks     []Hook
	mutation  *RoleMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the RoleUpdate builder.
func (ru *RoleUpdate) Where(ps ...predicate.Role) *RoleUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetUpdateTime sets the "update_time" field.
func (ru *RoleUpdate) SetUpdateTime(t time.Time) *RoleUpdate {
	ru.mutation.SetUpdateTime(t)
	return ru
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableUpdateTime(t *time.Time) *RoleUpdate {
	if t != nil {
		ru.SetUpdateTime(*t)
	}
	return ru
}

// ClearUpdateTime clears the value of the "update_time" field.
func (ru *RoleUpdate) ClearUpdateTime() *RoleUpdate {
	ru.mutation.ClearUpdateTime()
	return ru
}

// SetDeleteTime sets the "delete_time" field.
func (ru *RoleUpdate) SetDeleteTime(t time.Time) *RoleUpdate {
	ru.mutation.SetDeleteTime(t)
	return ru
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableDeleteTime(t *time.Time) *RoleUpdate {
	if t != nil {
		ru.SetDeleteTime(*t)
	}
	return ru
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (ru *RoleUpdate) ClearDeleteTime() *RoleUpdate {
	ru.mutation.ClearDeleteTime()
	return ru
}

// SetStatus sets the "status" field.
func (ru *RoleUpdate) SetStatus(r role.Status) *RoleUpdate {
	ru.mutation.SetStatus(r)
	return ru
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableStatus(r *role.Status) *RoleUpdate {
	if r != nil {
		ru.SetStatus(*r)
	}
	return ru
}

// ClearStatus clears the value of the "status" field.
func (ru *RoleUpdate) ClearStatus() *RoleUpdate {
	ru.mutation.ClearStatus()
	return ru
}

// SetCreateBy sets the "create_by" field.
func (ru *RoleUpdate) SetCreateBy(u uint32) *RoleUpdate {
	ru.mutation.ResetCreateBy()
	ru.mutation.SetCreateBy(u)
	return ru
}

// SetNillableCreateBy sets the "create_by" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableCreateBy(u *uint32) *RoleUpdate {
	if u != nil {
		ru.SetCreateBy(*u)
	}
	return ru
}

// AddCreateBy adds u to the "create_by" field.
func (ru *RoleUpdate) AddCreateBy(u int32) *RoleUpdate {
	ru.mutation.AddCreateBy(u)
	return ru
}

// ClearCreateBy clears the value of the "create_by" field.
func (ru *RoleUpdate) ClearCreateBy() *RoleUpdate {
	ru.mutation.ClearCreateBy()
	return ru
}

// SetRemark sets the "remark" field.
func (ru *RoleUpdate) SetRemark(s string) *RoleUpdate {
	ru.mutation.SetRemark(s)
	return ru
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableRemark(s *string) *RoleUpdate {
	if s != nil {
		ru.SetRemark(*s)
	}
	return ru
}

// ClearRemark clears the value of the "remark" field.
func (ru *RoleUpdate) ClearRemark() *RoleUpdate {
	ru.mutation.ClearRemark()
	return ru
}

// SetName sets the "name" field.
func (ru *RoleUpdate) SetName(s string) *RoleUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableName(s *string) *RoleUpdate {
	if s != nil {
		ru.SetName(*s)
	}
	return ru
}

// ClearName clears the value of the "name" field.
func (ru *RoleUpdate) ClearName() *RoleUpdate {
	ru.mutation.ClearName()
	return ru
}

// SetCode sets the "code" field.
func (ru *RoleUpdate) SetCode(s string) *RoleUpdate {
	ru.mutation.SetCode(s)
	return ru
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableCode(s *string) *RoleUpdate {
	if s != nil {
		ru.SetCode(*s)
	}
	return ru
}

// ClearCode clears the value of the "code" field.
func (ru *RoleUpdate) ClearCode() *RoleUpdate {
	ru.mutation.ClearCode()
	return ru
}

// SetParentID sets the "parent_id" field.
func (ru *RoleUpdate) SetParentID(u uint32) *RoleUpdate {
	ru.mutation.SetParentID(u)
	return ru
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableParentID(u *uint32) *RoleUpdate {
	if u != nil {
		ru.SetParentID(*u)
	}
	return ru
}

// ClearParentID clears the value of the "parent_id" field.
func (ru *RoleUpdate) ClearParentID() *RoleUpdate {
	ru.mutation.ClearParentID()
	return ru
}

// SetOrderNo sets the "order_no" field.
func (ru *RoleUpdate) SetOrderNo(i int32) *RoleUpdate {
	ru.mutation.ResetOrderNo()
	ru.mutation.SetOrderNo(i)
	return ru
}

// SetNillableOrderNo sets the "order_no" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableOrderNo(i *int32) *RoleUpdate {
	if i != nil {
		ru.SetOrderNo(*i)
	}
	return ru
}

// AddOrderNo adds i to the "order_no" field.
func (ru *RoleUpdate) AddOrderNo(i int32) *RoleUpdate {
	ru.mutation.AddOrderNo(i)
	return ru
}

// ClearOrderNo clears the value of the "order_no" field.
func (ru *RoleUpdate) ClearOrderNo() *RoleUpdate {
	ru.mutation.ClearOrderNo()
	return ru
}

// SetParent sets the "parent" edge to the Role entity.
func (ru *RoleUpdate) SetParent(r *Role) *RoleUpdate {
	return ru.SetParentID(r.ID)
}

// AddChildIDs adds the "children" edge to the Role entity by IDs.
func (ru *RoleUpdate) AddChildIDs(ids ...uint32) *RoleUpdate {
	ru.mutation.AddChildIDs(ids...)
	return ru
}

// AddChildren adds the "children" edges to the Role entity.
func (ru *RoleUpdate) AddChildren(r ...*Role) *RoleUpdate {
	ids := make([]uint32, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.AddChildIDs(ids...)
}

// Mutation returns the RoleMutation object of the builder.
func (ru *RoleUpdate) Mutation() *RoleMutation {
	return ru.mutation
}

// ClearParent clears the "parent" edge to the Role entity.
func (ru *RoleUpdate) ClearParent() *RoleUpdate {
	ru.mutation.ClearParent()
	return ru
}

// ClearChildren clears all "children" edges to the Role entity.
func (ru *RoleUpdate) ClearChildren() *RoleUpdate {
	ru.mutation.ClearChildren()
	return ru
}

// RemoveChildIDs removes the "children" edge to Role entities by IDs.
func (ru *RoleUpdate) RemoveChildIDs(ids ...uint32) *RoleUpdate {
	ru.mutation.RemoveChildIDs(ids...)
	return ru
}

// RemoveChildren removes "children" edges to Role entities.
func (ru *RoleUpdate) RemoveChildren(r ...*Role) *RoleUpdate {
	ids := make([]uint32, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.RemoveChildIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RoleUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RoleUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RoleUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RoleUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RoleUpdate) check() error {
	if v, ok := ru.mutation.Status(); ok {
		if err := role.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Role.status": %w`, err)}
		}
	}
	if v, ok := ru.mutation.Name(); ok {
		if err := role.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Role.name": %w`, err)}
		}
	}
	if v, ok := ru.mutation.Code(); ok {
		if err := role.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Role.code": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ru *RoleUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *RoleUpdate {
	ru.modifiers = append(ru.modifiers, modifiers...)
	return ru
}

func (ru *RoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(role.Table, role.Columns, sqlgraph.NewFieldSpec(role.FieldID, field.TypeUint32))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ru.mutation.CreateTimeCleared() {
		_spec.ClearField(role.FieldCreateTime, field.TypeTime)
	}
	if value, ok := ru.mutation.UpdateTime(); ok {
		_spec.SetField(role.FieldUpdateTime, field.TypeTime, value)
	}
	if ru.mutation.UpdateTimeCleared() {
		_spec.ClearField(role.FieldUpdateTime, field.TypeTime)
	}
	if value, ok := ru.mutation.DeleteTime(); ok {
		_spec.SetField(role.FieldDeleteTime, field.TypeTime, value)
	}
	if ru.mutation.DeleteTimeCleared() {
		_spec.ClearField(role.FieldDeleteTime, field.TypeTime)
	}
	if value, ok := ru.mutation.Status(); ok {
		_spec.SetField(role.FieldStatus, field.TypeEnum, value)
	}
	if ru.mutation.StatusCleared() {
		_spec.ClearField(role.FieldStatus, field.TypeEnum)
	}
	if value, ok := ru.mutation.CreateBy(); ok {
		_spec.SetField(role.FieldCreateBy, field.TypeUint32, value)
	}
	if value, ok := ru.mutation.AddedCreateBy(); ok {
		_spec.AddField(role.FieldCreateBy, field.TypeUint32, value)
	}
	if ru.mutation.CreateByCleared() {
		_spec.ClearField(role.FieldCreateBy, field.TypeUint32)
	}
	if value, ok := ru.mutation.Remark(); ok {
		_spec.SetField(role.FieldRemark, field.TypeString, value)
	}
	if ru.mutation.RemarkCleared() {
		_spec.ClearField(role.FieldRemark, field.TypeString)
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.SetField(role.FieldName, field.TypeString, value)
	}
	if ru.mutation.NameCleared() {
		_spec.ClearField(role.FieldName, field.TypeString)
	}
	if value, ok := ru.mutation.Code(); ok {
		_spec.SetField(role.FieldCode, field.TypeString, value)
	}
	if ru.mutation.CodeCleared() {
		_spec.ClearField(role.FieldCode, field.TypeString)
	}
	if value, ok := ru.mutation.OrderNo(); ok {
		_spec.SetField(role.FieldOrderNo, field.TypeInt32, value)
	}
	if value, ok := ru.mutation.AddedOrderNo(); ok {
		_spec.AddField(role.FieldOrderNo, field.TypeInt32, value)
	}
	if ru.mutation.OrderNoCleared() {
		_spec.ClearField(role.FieldOrderNo, field.TypeInt32)
	}
	if ru.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   role.ParentTable,
			Columns: []string{role.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   role.ParentTable,
			Columns: []string{role.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.ChildrenTable,
			Columns: []string{role.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !ru.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.ChildrenTable,
			Columns: []string{role.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.ChildrenTable,
			Columns: []string{role.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(ru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RoleUpdateOne is the builder for updating a single Role entity.
type RoleUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *RoleMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdateTime sets the "update_time" field.
func (ruo *RoleUpdateOne) SetUpdateTime(t time.Time) *RoleUpdateOne {
	ruo.mutation.SetUpdateTime(t)
	return ruo
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableUpdateTime(t *time.Time) *RoleUpdateOne {
	if t != nil {
		ruo.SetUpdateTime(*t)
	}
	return ruo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (ruo *RoleUpdateOne) ClearUpdateTime() *RoleUpdateOne {
	ruo.mutation.ClearUpdateTime()
	return ruo
}

// SetDeleteTime sets the "delete_time" field.
func (ruo *RoleUpdateOne) SetDeleteTime(t time.Time) *RoleUpdateOne {
	ruo.mutation.SetDeleteTime(t)
	return ruo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableDeleteTime(t *time.Time) *RoleUpdateOne {
	if t != nil {
		ruo.SetDeleteTime(*t)
	}
	return ruo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (ruo *RoleUpdateOne) ClearDeleteTime() *RoleUpdateOne {
	ruo.mutation.ClearDeleteTime()
	return ruo
}

// SetStatus sets the "status" field.
func (ruo *RoleUpdateOne) SetStatus(r role.Status) *RoleUpdateOne {
	ruo.mutation.SetStatus(r)
	return ruo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableStatus(r *role.Status) *RoleUpdateOne {
	if r != nil {
		ruo.SetStatus(*r)
	}
	return ruo
}

// ClearStatus clears the value of the "status" field.
func (ruo *RoleUpdateOne) ClearStatus() *RoleUpdateOne {
	ruo.mutation.ClearStatus()
	return ruo
}

// SetCreateBy sets the "create_by" field.
func (ruo *RoleUpdateOne) SetCreateBy(u uint32) *RoleUpdateOne {
	ruo.mutation.ResetCreateBy()
	ruo.mutation.SetCreateBy(u)
	return ruo
}

// SetNillableCreateBy sets the "create_by" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableCreateBy(u *uint32) *RoleUpdateOne {
	if u != nil {
		ruo.SetCreateBy(*u)
	}
	return ruo
}

// AddCreateBy adds u to the "create_by" field.
func (ruo *RoleUpdateOne) AddCreateBy(u int32) *RoleUpdateOne {
	ruo.mutation.AddCreateBy(u)
	return ruo
}

// ClearCreateBy clears the value of the "create_by" field.
func (ruo *RoleUpdateOne) ClearCreateBy() *RoleUpdateOne {
	ruo.mutation.ClearCreateBy()
	return ruo
}

// SetRemark sets the "remark" field.
func (ruo *RoleUpdateOne) SetRemark(s string) *RoleUpdateOne {
	ruo.mutation.SetRemark(s)
	return ruo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableRemark(s *string) *RoleUpdateOne {
	if s != nil {
		ruo.SetRemark(*s)
	}
	return ruo
}

// ClearRemark clears the value of the "remark" field.
func (ruo *RoleUpdateOne) ClearRemark() *RoleUpdateOne {
	ruo.mutation.ClearRemark()
	return ruo
}

// SetName sets the "name" field.
func (ruo *RoleUpdateOne) SetName(s string) *RoleUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableName(s *string) *RoleUpdateOne {
	if s != nil {
		ruo.SetName(*s)
	}
	return ruo
}

// ClearName clears the value of the "name" field.
func (ruo *RoleUpdateOne) ClearName() *RoleUpdateOne {
	ruo.mutation.ClearName()
	return ruo
}

// SetCode sets the "code" field.
func (ruo *RoleUpdateOne) SetCode(s string) *RoleUpdateOne {
	ruo.mutation.SetCode(s)
	return ruo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableCode(s *string) *RoleUpdateOne {
	if s != nil {
		ruo.SetCode(*s)
	}
	return ruo
}

// ClearCode clears the value of the "code" field.
func (ruo *RoleUpdateOne) ClearCode() *RoleUpdateOne {
	ruo.mutation.ClearCode()
	return ruo
}

// SetParentID sets the "parent_id" field.
func (ruo *RoleUpdateOne) SetParentID(u uint32) *RoleUpdateOne {
	ruo.mutation.SetParentID(u)
	return ruo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableParentID(u *uint32) *RoleUpdateOne {
	if u != nil {
		ruo.SetParentID(*u)
	}
	return ruo
}

// ClearParentID clears the value of the "parent_id" field.
func (ruo *RoleUpdateOne) ClearParentID() *RoleUpdateOne {
	ruo.mutation.ClearParentID()
	return ruo
}

// SetOrderNo sets the "order_no" field.
func (ruo *RoleUpdateOne) SetOrderNo(i int32) *RoleUpdateOne {
	ruo.mutation.ResetOrderNo()
	ruo.mutation.SetOrderNo(i)
	return ruo
}

// SetNillableOrderNo sets the "order_no" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableOrderNo(i *int32) *RoleUpdateOne {
	if i != nil {
		ruo.SetOrderNo(*i)
	}
	return ruo
}

// AddOrderNo adds i to the "order_no" field.
func (ruo *RoleUpdateOne) AddOrderNo(i int32) *RoleUpdateOne {
	ruo.mutation.AddOrderNo(i)
	return ruo
}

// ClearOrderNo clears the value of the "order_no" field.
func (ruo *RoleUpdateOne) ClearOrderNo() *RoleUpdateOne {
	ruo.mutation.ClearOrderNo()
	return ruo
}

// SetParent sets the "parent" edge to the Role entity.
func (ruo *RoleUpdateOne) SetParent(r *Role) *RoleUpdateOne {
	return ruo.SetParentID(r.ID)
}

// AddChildIDs adds the "children" edge to the Role entity by IDs.
func (ruo *RoleUpdateOne) AddChildIDs(ids ...uint32) *RoleUpdateOne {
	ruo.mutation.AddChildIDs(ids...)
	return ruo
}

// AddChildren adds the "children" edges to the Role entity.
func (ruo *RoleUpdateOne) AddChildren(r ...*Role) *RoleUpdateOne {
	ids := make([]uint32, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.AddChildIDs(ids...)
}

// Mutation returns the RoleMutation object of the builder.
func (ruo *RoleUpdateOne) Mutation() *RoleMutation {
	return ruo.mutation
}

// ClearParent clears the "parent" edge to the Role entity.
func (ruo *RoleUpdateOne) ClearParent() *RoleUpdateOne {
	ruo.mutation.ClearParent()
	return ruo
}

// ClearChildren clears all "children" edges to the Role entity.
func (ruo *RoleUpdateOne) ClearChildren() *RoleUpdateOne {
	ruo.mutation.ClearChildren()
	return ruo
}

// RemoveChildIDs removes the "children" edge to Role entities by IDs.
func (ruo *RoleUpdateOne) RemoveChildIDs(ids ...uint32) *RoleUpdateOne {
	ruo.mutation.RemoveChildIDs(ids...)
	return ruo
}

// RemoveChildren removes "children" edges to Role entities.
func (ruo *RoleUpdateOne) RemoveChildren(r ...*Role) *RoleUpdateOne {
	ids := make([]uint32, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.RemoveChildIDs(ids...)
}

// Where appends a list predicates to the RoleUpdate builder.
func (ruo *RoleUpdateOne) Where(ps ...predicate.Role) *RoleUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RoleUpdateOne) Select(field string, fields ...string) *RoleUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Role entity.
func (ruo *RoleUpdateOne) Save(ctx context.Context) (*Role, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RoleUpdateOne) SaveX(ctx context.Context) *Role {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RoleUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RoleUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RoleUpdateOne) check() error {
	if v, ok := ruo.mutation.Status(); ok {
		if err := role.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Role.status": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.Name(); ok {
		if err := role.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Role.name": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.Code(); ok {
		if err := role.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Role.code": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ruo *RoleUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *RoleUpdateOne {
	ruo.modifiers = append(ruo.modifiers, modifiers...)
	return ruo
}

func (ruo *RoleUpdateOne) sqlSave(ctx context.Context) (_node *Role, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(role.Table, role.Columns, sqlgraph.NewFieldSpec(role.FieldID, field.TypeUint32))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Role.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, role.FieldID)
		for _, f := range fields {
			if !role.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != role.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ruo.mutation.CreateTimeCleared() {
		_spec.ClearField(role.FieldCreateTime, field.TypeTime)
	}
	if value, ok := ruo.mutation.UpdateTime(); ok {
		_spec.SetField(role.FieldUpdateTime, field.TypeTime, value)
	}
	if ruo.mutation.UpdateTimeCleared() {
		_spec.ClearField(role.FieldUpdateTime, field.TypeTime)
	}
	if value, ok := ruo.mutation.DeleteTime(); ok {
		_spec.SetField(role.FieldDeleteTime, field.TypeTime, value)
	}
	if ruo.mutation.DeleteTimeCleared() {
		_spec.ClearField(role.FieldDeleteTime, field.TypeTime)
	}
	if value, ok := ruo.mutation.Status(); ok {
		_spec.SetField(role.FieldStatus, field.TypeEnum, value)
	}
	if ruo.mutation.StatusCleared() {
		_spec.ClearField(role.FieldStatus, field.TypeEnum)
	}
	if value, ok := ruo.mutation.CreateBy(); ok {
		_spec.SetField(role.FieldCreateBy, field.TypeUint32, value)
	}
	if value, ok := ruo.mutation.AddedCreateBy(); ok {
		_spec.AddField(role.FieldCreateBy, field.TypeUint32, value)
	}
	if ruo.mutation.CreateByCleared() {
		_spec.ClearField(role.FieldCreateBy, field.TypeUint32)
	}
	if value, ok := ruo.mutation.Remark(); ok {
		_spec.SetField(role.FieldRemark, field.TypeString, value)
	}
	if ruo.mutation.RemarkCleared() {
		_spec.ClearField(role.FieldRemark, field.TypeString)
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.SetField(role.FieldName, field.TypeString, value)
	}
	if ruo.mutation.NameCleared() {
		_spec.ClearField(role.FieldName, field.TypeString)
	}
	if value, ok := ruo.mutation.Code(); ok {
		_spec.SetField(role.FieldCode, field.TypeString, value)
	}
	if ruo.mutation.CodeCleared() {
		_spec.ClearField(role.FieldCode, field.TypeString)
	}
	if value, ok := ruo.mutation.OrderNo(); ok {
		_spec.SetField(role.FieldOrderNo, field.TypeInt32, value)
	}
	if value, ok := ruo.mutation.AddedOrderNo(); ok {
		_spec.AddField(role.FieldOrderNo, field.TypeInt32, value)
	}
	if ruo.mutation.OrderNoCleared() {
		_spec.ClearField(role.FieldOrderNo, field.TypeInt32)
	}
	if ruo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   role.ParentTable,
			Columns: []string{role.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   role.ParentTable,
			Columns: []string{role.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.ChildrenTable,
			Columns: []string{role.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !ruo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.ChildrenTable,
			Columns: []string{role.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.ChildrenTable,
			Columns: []string{role.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(ruo.modifiers...)
	_node = &Role{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}

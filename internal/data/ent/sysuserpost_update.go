// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/liangtengfei/beer-admin/internal/data/ent/predicate"
	"github.com/liangtengfei/beer-admin/internal/data/ent/syspost"
	"github.com/liangtengfei/beer-admin/internal/data/ent/sysuser"
	"github.com/liangtengfei/beer-admin/internal/data/ent/sysuserpost"
)

// SysUserPostUpdate is the builder for updating SysUserPost entities.
type SysUserPostUpdate struct {
	config
	hooks    []Hook
	mutation *SysUserPostMutation
}

// Where appends a list predicates to the SysUserPostUpdate builder.
func (supu *SysUserPostUpdate) Where(ps ...predicate.SysUserPost) *SysUserPostUpdate {
	supu.mutation.Where(ps...)
	return supu
}

// SetSysUserID sets the "sys_user_id" field.
func (supu *SysUserPostUpdate) SetSysUserID(u uint64) *SysUserPostUpdate {
	supu.mutation.SetSysUserID(u)
	return supu
}

// SetSysPostID sets the "sys_post_id" field.
func (supu *SysUserPostUpdate) SetSysPostID(u uint64) *SysUserPostUpdate {
	supu.mutation.SetSysPostID(u)
	return supu
}

// SetCreatedAt sets the "created_at" field.
func (supu *SysUserPostUpdate) SetCreatedAt(t time.Time) *SysUserPostUpdate {
	supu.mutation.SetCreatedAt(t)
	return supu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (supu *SysUserPostUpdate) SetNillableCreatedAt(t *time.Time) *SysUserPostUpdate {
	if t != nil {
		supu.SetCreatedAt(*t)
	}
	return supu
}

// SetUserID sets the "user" edge to the SysUser entity by ID.
func (supu *SysUserPostUpdate) SetUserID(id uint64) *SysUserPostUpdate {
	supu.mutation.SetUserID(id)
	return supu
}

// SetUser sets the "user" edge to the SysUser entity.
func (supu *SysUserPostUpdate) SetUser(s *SysUser) *SysUserPostUpdate {
	return supu.SetUserID(s.ID)
}

// SetPostID sets the "post" edge to the SysPost entity by ID.
func (supu *SysUserPostUpdate) SetPostID(id uint64) *SysUserPostUpdate {
	supu.mutation.SetPostID(id)
	return supu
}

// SetPost sets the "post" edge to the SysPost entity.
func (supu *SysUserPostUpdate) SetPost(s *SysPost) *SysUserPostUpdate {
	return supu.SetPostID(s.ID)
}

// Mutation returns the SysUserPostMutation object of the builder.
func (supu *SysUserPostUpdate) Mutation() *SysUserPostMutation {
	return supu.mutation
}

// ClearUser clears the "user" edge to the SysUser entity.
func (supu *SysUserPostUpdate) ClearUser() *SysUserPostUpdate {
	supu.mutation.ClearUser()
	return supu
}

// ClearPost clears the "post" edge to the SysPost entity.
func (supu *SysUserPostUpdate) ClearPost() *SysUserPostUpdate {
	supu.mutation.ClearPost()
	return supu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (supu *SysUserPostUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, SysUserPostMutation](ctx, supu.sqlSave, supu.mutation, supu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (supu *SysUserPostUpdate) SaveX(ctx context.Context) int {
	affected, err := supu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (supu *SysUserPostUpdate) Exec(ctx context.Context) error {
	_, err := supu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (supu *SysUserPostUpdate) ExecX(ctx context.Context) {
	if err := supu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (supu *SysUserPostUpdate) check() error {
	if _, ok := supu.mutation.UserID(); supu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "SysUserPost.user"`)
	}
	if _, ok := supu.mutation.PostID(); supu.mutation.PostCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "SysUserPost.post"`)
	}
	return nil
}

func (supu *SysUserPostUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := supu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(sysuserpost.Table, sysuserpost.Columns, sqlgraph.NewFieldSpec(sysuserpost.FieldSysUserID, field.TypeUint64), sqlgraph.NewFieldSpec(sysuserpost.FieldSysPostID, field.TypeUint64))
	if ps := supu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := supu.mutation.CreatedAt(); ok {
		_spec.SetField(sysuserpost.FieldCreatedAt, field.TypeTime, value)
	}
	if supu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sysuserpost.UserTable,
			Columns: []string{sysuserpost.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysuser.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := supu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sysuserpost.UserTable,
			Columns: []string{sysuserpost.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysuser.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if supu.mutation.PostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sysuserpost.PostTable,
			Columns: []string{sysuserpost.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(syspost.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := supu.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sysuserpost.PostTable,
			Columns: []string{sysuserpost.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(syspost.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, supu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysuserpost.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	supu.mutation.done = true
	return n, nil
}

// SysUserPostUpdateOne is the builder for updating a single SysUserPost entity.
type SysUserPostUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SysUserPostMutation
}

// SetSysUserID sets the "sys_user_id" field.
func (supuo *SysUserPostUpdateOne) SetSysUserID(u uint64) *SysUserPostUpdateOne {
	supuo.mutation.SetSysUserID(u)
	return supuo
}

// SetSysPostID sets the "sys_post_id" field.
func (supuo *SysUserPostUpdateOne) SetSysPostID(u uint64) *SysUserPostUpdateOne {
	supuo.mutation.SetSysPostID(u)
	return supuo
}

// SetCreatedAt sets the "created_at" field.
func (supuo *SysUserPostUpdateOne) SetCreatedAt(t time.Time) *SysUserPostUpdateOne {
	supuo.mutation.SetCreatedAt(t)
	return supuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (supuo *SysUserPostUpdateOne) SetNillableCreatedAt(t *time.Time) *SysUserPostUpdateOne {
	if t != nil {
		supuo.SetCreatedAt(*t)
	}
	return supuo
}

// SetUserID sets the "user" edge to the SysUser entity by ID.
func (supuo *SysUserPostUpdateOne) SetUserID(id uint64) *SysUserPostUpdateOne {
	supuo.mutation.SetUserID(id)
	return supuo
}

// SetUser sets the "user" edge to the SysUser entity.
func (supuo *SysUserPostUpdateOne) SetUser(s *SysUser) *SysUserPostUpdateOne {
	return supuo.SetUserID(s.ID)
}

// SetPostID sets the "post" edge to the SysPost entity by ID.
func (supuo *SysUserPostUpdateOne) SetPostID(id uint64) *SysUserPostUpdateOne {
	supuo.mutation.SetPostID(id)
	return supuo
}

// SetPost sets the "post" edge to the SysPost entity.
func (supuo *SysUserPostUpdateOne) SetPost(s *SysPost) *SysUserPostUpdateOne {
	return supuo.SetPostID(s.ID)
}

// Mutation returns the SysUserPostMutation object of the builder.
func (supuo *SysUserPostUpdateOne) Mutation() *SysUserPostMutation {
	return supuo.mutation
}

// ClearUser clears the "user" edge to the SysUser entity.
func (supuo *SysUserPostUpdateOne) ClearUser() *SysUserPostUpdateOne {
	supuo.mutation.ClearUser()
	return supuo
}

// ClearPost clears the "post" edge to the SysPost entity.
func (supuo *SysUserPostUpdateOne) ClearPost() *SysUserPostUpdateOne {
	supuo.mutation.ClearPost()
	return supuo
}

// Where appends a list predicates to the SysUserPostUpdate builder.
func (supuo *SysUserPostUpdateOne) Where(ps ...predicate.SysUserPost) *SysUserPostUpdateOne {
	supuo.mutation.Where(ps...)
	return supuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (supuo *SysUserPostUpdateOne) Select(field string, fields ...string) *SysUserPostUpdateOne {
	supuo.fields = append([]string{field}, fields...)
	return supuo
}

// Save executes the query and returns the updated SysUserPost entity.
func (supuo *SysUserPostUpdateOne) Save(ctx context.Context) (*SysUserPost, error) {
	return withHooks[*SysUserPost, SysUserPostMutation](ctx, supuo.sqlSave, supuo.mutation, supuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (supuo *SysUserPostUpdateOne) SaveX(ctx context.Context) *SysUserPost {
	node, err := supuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (supuo *SysUserPostUpdateOne) Exec(ctx context.Context) error {
	_, err := supuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (supuo *SysUserPostUpdateOne) ExecX(ctx context.Context) {
	if err := supuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (supuo *SysUserPostUpdateOne) check() error {
	if _, ok := supuo.mutation.UserID(); supuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "SysUserPost.user"`)
	}
	if _, ok := supuo.mutation.PostID(); supuo.mutation.PostCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "SysUserPost.post"`)
	}
	return nil
}

func (supuo *SysUserPostUpdateOne) sqlSave(ctx context.Context) (_node *SysUserPost, err error) {
	if err := supuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(sysuserpost.Table, sysuserpost.Columns, sqlgraph.NewFieldSpec(sysuserpost.FieldSysUserID, field.TypeUint64), sqlgraph.NewFieldSpec(sysuserpost.FieldSysPostID, field.TypeUint64))
	if id, ok := supuo.mutation.SysUserID(); !ok {
		return nil, &ValidationError{Name: "sys_user_id", err: errors.New(`ent: missing "SysUserPost.sys_user_id" for update`)}
	} else {
		_spec.Node.CompositeID[0].Value = id
	}
	if id, ok := supuo.mutation.SysPostID(); !ok {
		return nil, &ValidationError{Name: "sys_post_id", err: errors.New(`ent: missing "SysUserPost.sys_post_id" for update`)}
	} else {
		_spec.Node.CompositeID[1].Value = id
	}
	if fields := supuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, len(fields))
		for i, f := range fields {
			if !sysuserpost.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			_spec.Node.Columns[i] = f
		}
	}
	if ps := supuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := supuo.mutation.CreatedAt(); ok {
		_spec.SetField(sysuserpost.FieldCreatedAt, field.TypeTime, value)
	}
	if supuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sysuserpost.UserTable,
			Columns: []string{sysuserpost.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysuser.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := supuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sysuserpost.UserTable,
			Columns: []string{sysuserpost.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysuser.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if supuo.mutation.PostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sysuserpost.PostTable,
			Columns: []string{sysuserpost.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(syspost.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := supuo.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sysuserpost.PostTable,
			Columns: []string{sysuserpost.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(syspost.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SysUserPost{config: supuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, supuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysuserpost.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	supuo.mutation.done = true
	return _node, nil
}

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/liangtengfei/beer-admin/internal/data/ent/syspost"
	"github.com/liangtengfei/beer-admin/internal/data/ent/sysuser"
)

// SysPostCreate is the builder for creating a SysPost entity.
type SysPostCreate struct {
	config
	mutation *SysPostMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (spc *SysPostCreate) SetCreatedAt(t time.Time) *SysPostCreate {
	spc.mutation.SetCreatedAt(t)
	return spc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (spc *SysPostCreate) SetNillableCreatedAt(t *time.Time) *SysPostCreate {
	if t != nil {
		spc.SetCreatedAt(*t)
	}
	return spc
}

// SetUpdatedAt sets the "updated_at" field.
func (spc *SysPostCreate) SetUpdatedAt(t time.Time) *SysPostCreate {
	spc.mutation.SetUpdatedAt(t)
	return spc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (spc *SysPostCreate) SetNillableUpdatedAt(t *time.Time) *SysPostCreate {
	if t != nil {
		spc.SetUpdatedAt(*t)
	}
	return spc
}

// SetDeletedAt sets the "deleted_at" field.
func (spc *SysPostCreate) SetDeletedAt(t time.Time) *SysPostCreate {
	spc.mutation.SetDeletedAt(t)
	return spc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (spc *SysPostCreate) SetNillableDeletedAt(t *time.Time) *SysPostCreate {
	if t != nil {
		spc.SetDeletedAt(*t)
	}
	return spc
}

// SetCreatedBy sets the "created_by" field.
func (spc *SysPostCreate) SetCreatedBy(s string) *SysPostCreate {
	spc.mutation.SetCreatedBy(s)
	return spc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (spc *SysPostCreate) SetNillableCreatedBy(s *string) *SysPostCreate {
	if s != nil {
		spc.SetCreatedBy(*s)
	}
	return spc
}

// SetUpdatedBy sets the "updated_by" field.
func (spc *SysPostCreate) SetUpdatedBy(s string) *SysPostCreate {
	spc.mutation.SetUpdatedBy(s)
	return spc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (spc *SysPostCreate) SetNillableUpdatedBy(s *string) *SysPostCreate {
	if s != nil {
		spc.SetUpdatedBy(*s)
	}
	return spc
}

// SetPostName sets the "post_name" field.
func (spc *SysPostCreate) SetPostName(s string) *SysPostCreate {
	spc.mutation.SetPostName(s)
	return spc
}

// SetStatus sets the "status" field.
func (spc *SysPostCreate) SetStatus(s syspost.Status) *SysPostCreate {
	spc.mutation.SetStatus(s)
	return spc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (spc *SysPostCreate) SetNillableStatus(s *syspost.Status) *SysPostCreate {
	if s != nil {
		spc.SetStatus(*s)
	}
	return spc
}

// SetRemark sets the "remark" field.
func (spc *SysPostCreate) SetRemark(s string) *SysPostCreate {
	spc.mutation.SetRemark(s)
	return spc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (spc *SysPostCreate) SetNillableRemark(s *string) *SysPostCreate {
	if s != nil {
		spc.SetRemark(*s)
	}
	return spc
}

// SetID sets the "id" field.
func (spc *SysPostCreate) SetID(u uint64) *SysPostCreate {
	spc.mutation.SetID(u)
	return spc
}

// AddPostsUserIDs adds the "posts_user" edge to the SysUser entity by IDs.
func (spc *SysPostCreate) AddPostsUserIDs(ids ...uint64) *SysPostCreate {
	spc.mutation.AddPostsUserIDs(ids...)
	return spc
}

// AddPostsUser adds the "posts_user" edges to the SysUser entity.
func (spc *SysPostCreate) AddPostsUser(s ...*SysUser) *SysPostCreate {
	ids := make([]uint64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return spc.AddPostsUserIDs(ids...)
}

// Mutation returns the SysPostMutation object of the builder.
func (spc *SysPostCreate) Mutation() *SysPostMutation {
	return spc.mutation
}

// Save creates the SysPost in the database.
func (spc *SysPostCreate) Save(ctx context.Context) (*SysPost, error) {
	spc.defaults()
	return withHooks[*SysPost, SysPostMutation](ctx, spc.sqlSave, spc.mutation, spc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (spc *SysPostCreate) SaveX(ctx context.Context) *SysPost {
	v, err := spc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (spc *SysPostCreate) Exec(ctx context.Context) error {
	_, err := spc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spc *SysPostCreate) ExecX(ctx context.Context) {
	if err := spc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (spc *SysPostCreate) defaults() {
	if _, ok := spc.mutation.CreatedAt(); !ok {
		v := syspost.DefaultCreatedAt()
		spc.mutation.SetCreatedAt(v)
	}
	if _, ok := spc.mutation.UpdatedAt(); !ok {
		v := syspost.DefaultUpdatedAt()
		spc.mutation.SetUpdatedAt(v)
	}
	if _, ok := spc.mutation.CreatedBy(); !ok {
		v := syspost.DefaultCreatedBy
		spc.mutation.SetCreatedBy(v)
	}
	if _, ok := spc.mutation.UpdatedBy(); !ok {
		v := syspost.DefaultUpdatedBy
		spc.mutation.SetUpdatedBy(v)
	}
	if _, ok := spc.mutation.Status(); !ok {
		v := syspost.DefaultStatus
		spc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (spc *SysPostCreate) check() error {
	if _, ok := spc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "SysPost.created_at"`)}
	}
	if _, ok := spc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "SysPost.updated_at"`)}
	}
	if _, ok := spc.mutation.PostName(); !ok {
		return &ValidationError{Name: "post_name", err: errors.New(`ent: missing required field "SysPost.post_name"`)}
	}
	if v, ok := spc.mutation.PostName(); ok {
		if err := syspost.PostNameValidator(v); err != nil {
			return &ValidationError{Name: "post_name", err: fmt.Errorf(`ent: validator failed for field "SysPost.post_name": %w`, err)}
		}
	}
	if _, ok := spc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "SysPost.status"`)}
	}
	if v, ok := spc.mutation.Status(); ok {
		if err := syspost.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "SysPost.status": %w`, err)}
		}
	}
	return nil
}

func (spc *SysPostCreate) sqlSave(ctx context.Context) (*SysPost, error) {
	if err := spc.check(); err != nil {
		return nil, err
	}
	_node, _spec := spc.createSpec()
	if err := sqlgraph.CreateNode(ctx, spc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	spc.mutation.id = &_node.ID
	spc.mutation.done = true
	return _node, nil
}

func (spc *SysPostCreate) createSpec() (*SysPost, *sqlgraph.CreateSpec) {
	var (
		_node = &SysPost{config: spc.config}
		_spec = sqlgraph.NewCreateSpec(syspost.Table, sqlgraph.NewFieldSpec(syspost.FieldID, field.TypeUint64))
	)
	if id, ok := spc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := spc.mutation.CreatedAt(); ok {
		_spec.SetField(syspost.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := spc.mutation.UpdatedAt(); ok {
		_spec.SetField(syspost.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := spc.mutation.DeletedAt(); ok {
		_spec.SetField(syspost.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := spc.mutation.CreatedBy(); ok {
		_spec.SetField(syspost.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := spc.mutation.UpdatedBy(); ok {
		_spec.SetField(syspost.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := spc.mutation.PostName(); ok {
		_spec.SetField(syspost.FieldPostName, field.TypeString, value)
		_node.PostName = value
	}
	if value, ok := spc.mutation.Status(); ok {
		_spec.SetField(syspost.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := spc.mutation.Remark(); ok {
		_spec.SetField(syspost.FieldRemark, field.TypeString, value)
		_node.Remark = value
	}
	if nodes := spc.mutation.PostsUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   syspost.PostsUserTable,
			Columns: syspost.PostsUserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysuser.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &SysUserPostCreate{config: spc.config, mutation: newSysUserPostMutation(spc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SysPostCreateBulk is the builder for creating many SysPost entities in bulk.
type SysPostCreateBulk struct {
	config
	builders []*SysPostCreate
}

// Save creates the SysPost entities in the database.
func (spcb *SysPostCreateBulk) Save(ctx context.Context) ([]*SysPost, error) {
	specs := make([]*sqlgraph.CreateSpec, len(spcb.builders))
	nodes := make([]*SysPost, len(spcb.builders))
	mutators := make([]Mutator, len(spcb.builders))
	for i := range spcb.builders {
		func(i int, root context.Context) {
			builder := spcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysPostMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, spcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, spcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, spcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (spcb *SysPostCreateBulk) SaveX(ctx context.Context) []*SysPost {
	v, err := spcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (spcb *SysPostCreateBulk) Exec(ctx context.Context) error {
	_, err := spcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spcb *SysPostCreateBulk) ExecX(ctx context.Context) {
	if err := spcb.Exec(ctx); err != nil {
		panic(err)
	}
}

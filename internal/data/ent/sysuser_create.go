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

// SysUserCreate is the builder for creating a SysUser entity.
type SysUserCreate struct {
	config
	mutation *SysUserMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (suc *SysUserCreate) SetCreatedAt(t time.Time) *SysUserCreate {
	suc.mutation.SetCreatedAt(t)
	return suc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableCreatedAt(t *time.Time) *SysUserCreate {
	if t != nil {
		suc.SetCreatedAt(*t)
	}
	return suc
}

// SetUpdatedAt sets the "updated_at" field.
func (suc *SysUserCreate) SetUpdatedAt(t time.Time) *SysUserCreate {
	suc.mutation.SetUpdatedAt(t)
	return suc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableUpdatedAt(t *time.Time) *SysUserCreate {
	if t != nil {
		suc.SetUpdatedAt(*t)
	}
	return suc
}

// SetDeletedAt sets the "deleted_at" field.
func (suc *SysUserCreate) SetDeletedAt(t time.Time) *SysUserCreate {
	suc.mutation.SetDeletedAt(t)
	return suc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableDeletedAt(t *time.Time) *SysUserCreate {
	if t != nil {
		suc.SetDeletedAt(*t)
	}
	return suc
}

// SetCreatedBy sets the "created_by" field.
func (suc *SysUserCreate) SetCreatedBy(s string) *SysUserCreate {
	suc.mutation.SetCreatedBy(s)
	return suc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableCreatedBy(s *string) *SysUserCreate {
	if s != nil {
		suc.SetCreatedBy(*s)
	}
	return suc
}

// SetUpdatedBy sets the "updated_by" field.
func (suc *SysUserCreate) SetUpdatedBy(s string) *SysUserCreate {
	suc.mutation.SetUpdatedBy(s)
	return suc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableUpdatedBy(s *string) *SysUserCreate {
	if s != nil {
		suc.SetUpdatedBy(*s)
	}
	return suc
}

// SetDeptID sets the "dept_id" field.
func (suc *SysUserCreate) SetDeptID(u uint64) *SysUserCreate {
	suc.mutation.SetDeptID(u)
	return suc
}

// SetNillableDeptID sets the "dept_id" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableDeptID(u *uint64) *SysUserCreate {
	if u != nil {
		suc.SetDeptID(*u)
	}
	return suc
}

// SetUserName sets the "user_name" field.
func (suc *SysUserCreate) SetUserName(s string) *SysUserCreate {
	suc.mutation.SetUserName(s)
	return suc
}

// SetRealName sets the "real_name" field.
func (suc *SysUserCreate) SetRealName(s string) *SysUserCreate {
	suc.mutation.SetRealName(s)
	return suc
}

// SetMobile sets the "mobile" field.
func (suc *SysUserCreate) SetMobile(s string) *SysUserCreate {
	suc.mutation.SetMobile(s)
	return suc
}

// SetEmail sets the "email" field.
func (suc *SysUserCreate) SetEmail(s string) *SysUserCreate {
	suc.mutation.SetEmail(s)
	return suc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableEmail(s *string) *SysUserCreate {
	if s != nil {
		suc.SetEmail(*s)
	}
	return suc
}

// SetAvatar sets the "avatar" field.
func (suc *SysUserCreate) SetAvatar(s string) *SysUserCreate {
	suc.mutation.SetAvatar(s)
	return suc
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableAvatar(s *string) *SysUserCreate {
	if s != nil {
		suc.SetAvatar(*s)
	}
	return suc
}

// SetSex sets the "sex" field.
func (suc *SysUserCreate) SetSex(s sysuser.Sex) *SysUserCreate {
	suc.mutation.SetSex(s)
	return suc
}

// SetNillableSex sets the "sex" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableSex(s *sysuser.Sex) *SysUserCreate {
	if s != nil {
		suc.SetSex(*s)
	}
	return suc
}

// SetStatus sets the "status" field.
func (suc *SysUserCreate) SetStatus(s sysuser.Status) *SysUserCreate {
	suc.mutation.SetStatus(s)
	return suc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableStatus(s *sysuser.Status) *SysUserCreate {
	if s != nil {
		suc.SetStatus(*s)
	}
	return suc
}

// SetPassword sets the "password" field.
func (suc *SysUserCreate) SetPassword(s string) *SysUserCreate {
	suc.mutation.SetPassword(s)
	return suc
}

// SetRemark sets the "remark" field.
func (suc *SysUserCreate) SetRemark(s string) *SysUserCreate {
	suc.mutation.SetRemark(s)
	return suc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableRemark(s *string) *SysUserCreate {
	if s != nil {
		suc.SetRemark(*s)
	}
	return suc
}

// SetID sets the "id" field.
func (suc *SysUserCreate) SetID(u uint64) *SysUserCreate {
	suc.mutation.SetID(u)
	return suc
}

// AddUsersPostIDs adds the "users_post" edge to the SysPost entity by IDs.
func (suc *SysUserCreate) AddUsersPostIDs(ids ...uint64) *SysUserCreate {
	suc.mutation.AddUsersPostIDs(ids...)
	return suc
}

// AddUsersPost adds the "users_post" edges to the SysPost entity.
func (suc *SysUserCreate) AddUsersPost(s ...*SysPost) *SysUserCreate {
	ids := make([]uint64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suc.AddUsersPostIDs(ids...)
}

// Mutation returns the SysUserMutation object of the builder.
func (suc *SysUserCreate) Mutation() *SysUserMutation {
	return suc.mutation
}

// Save creates the SysUser in the database.
func (suc *SysUserCreate) Save(ctx context.Context) (*SysUser, error) {
	suc.defaults()
	return withHooks[*SysUser, SysUserMutation](ctx, suc.sqlSave, suc.mutation, suc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (suc *SysUserCreate) SaveX(ctx context.Context) *SysUser {
	v, err := suc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (suc *SysUserCreate) Exec(ctx context.Context) error {
	_, err := suc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suc *SysUserCreate) ExecX(ctx context.Context) {
	if err := suc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suc *SysUserCreate) defaults() {
	if _, ok := suc.mutation.CreatedAt(); !ok {
		v := sysuser.DefaultCreatedAt()
		suc.mutation.SetCreatedAt(v)
	}
	if _, ok := suc.mutation.UpdatedAt(); !ok {
		v := sysuser.DefaultUpdatedAt()
		suc.mutation.SetUpdatedAt(v)
	}
	if _, ok := suc.mutation.CreatedBy(); !ok {
		v := sysuser.DefaultCreatedBy
		suc.mutation.SetCreatedBy(v)
	}
	if _, ok := suc.mutation.UpdatedBy(); !ok {
		v := sysuser.DefaultUpdatedBy
		suc.mutation.SetUpdatedBy(v)
	}
	if _, ok := suc.mutation.DeptID(); !ok {
		v := sysuser.DefaultDeptID
		suc.mutation.SetDeptID(v)
	}
	if _, ok := suc.mutation.Sex(); !ok {
		v := sysuser.DefaultSex
		suc.mutation.SetSex(v)
	}
	if _, ok := suc.mutation.Status(); !ok {
		v := sysuser.DefaultStatus
		suc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suc *SysUserCreate) check() error {
	if _, ok := suc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "SysUser.created_at"`)}
	}
	if _, ok := suc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "SysUser.updated_at"`)}
	}
	if _, ok := suc.mutation.UserName(); !ok {
		return &ValidationError{Name: "user_name", err: errors.New(`ent: missing required field "SysUser.user_name"`)}
	}
	if v, ok := suc.mutation.UserName(); ok {
		if err := sysuser.UserNameValidator(v); err != nil {
			return &ValidationError{Name: "user_name", err: fmt.Errorf(`ent: validator failed for field "SysUser.user_name": %w`, err)}
		}
	}
	if _, ok := suc.mutation.RealName(); !ok {
		return &ValidationError{Name: "real_name", err: errors.New(`ent: missing required field "SysUser.real_name"`)}
	}
	if v, ok := suc.mutation.RealName(); ok {
		if err := sysuser.RealNameValidator(v); err != nil {
			return &ValidationError{Name: "real_name", err: fmt.Errorf(`ent: validator failed for field "SysUser.real_name": %w`, err)}
		}
	}
	if _, ok := suc.mutation.Mobile(); !ok {
		return &ValidationError{Name: "mobile", err: errors.New(`ent: missing required field "SysUser.mobile"`)}
	}
	if v, ok := suc.mutation.Mobile(); ok {
		if err := sysuser.MobileValidator(v); err != nil {
			return &ValidationError{Name: "mobile", err: fmt.Errorf(`ent: validator failed for field "SysUser.mobile": %w`, err)}
		}
	}
	if _, ok := suc.mutation.Sex(); !ok {
		return &ValidationError{Name: "sex", err: errors.New(`ent: missing required field "SysUser.sex"`)}
	}
	if v, ok := suc.mutation.Sex(); ok {
		if err := sysuser.SexValidator(v); err != nil {
			return &ValidationError{Name: "sex", err: fmt.Errorf(`ent: validator failed for field "SysUser.sex": %w`, err)}
		}
	}
	if _, ok := suc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "SysUser.status"`)}
	}
	if v, ok := suc.mutation.Status(); ok {
		if err := sysuser.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "SysUser.status": %w`, err)}
		}
	}
	if _, ok := suc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "SysUser.password"`)}
	}
	if v, ok := suc.mutation.Password(); ok {
		if err := sysuser.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "SysUser.password": %w`, err)}
		}
	}
	return nil
}

func (suc *SysUserCreate) sqlSave(ctx context.Context) (*SysUser, error) {
	if err := suc.check(); err != nil {
		return nil, err
	}
	_node, _spec := suc.createSpec()
	if err := sqlgraph.CreateNode(ctx, suc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	suc.mutation.id = &_node.ID
	suc.mutation.done = true
	return _node, nil
}

func (suc *SysUserCreate) createSpec() (*SysUser, *sqlgraph.CreateSpec) {
	var (
		_node = &SysUser{config: suc.config}
		_spec = sqlgraph.NewCreateSpec(sysuser.Table, sqlgraph.NewFieldSpec(sysuser.FieldID, field.TypeUint64))
	)
	if id, ok := suc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := suc.mutation.CreatedAt(); ok {
		_spec.SetField(sysuser.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := suc.mutation.UpdatedAt(); ok {
		_spec.SetField(sysuser.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := suc.mutation.DeletedAt(); ok {
		_spec.SetField(sysuser.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := suc.mutation.CreatedBy(); ok {
		_spec.SetField(sysuser.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := suc.mutation.UpdatedBy(); ok {
		_spec.SetField(sysuser.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := suc.mutation.DeptID(); ok {
		_spec.SetField(sysuser.FieldDeptID, field.TypeUint64, value)
		_node.DeptID = value
	}
	if value, ok := suc.mutation.UserName(); ok {
		_spec.SetField(sysuser.FieldUserName, field.TypeString, value)
		_node.UserName = value
	}
	if value, ok := suc.mutation.RealName(); ok {
		_spec.SetField(sysuser.FieldRealName, field.TypeString, value)
		_node.RealName = value
	}
	if value, ok := suc.mutation.Mobile(); ok {
		_spec.SetField(sysuser.FieldMobile, field.TypeString, value)
		_node.Mobile = value
	}
	if value, ok := suc.mutation.Email(); ok {
		_spec.SetField(sysuser.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := suc.mutation.Avatar(); ok {
		_spec.SetField(sysuser.FieldAvatar, field.TypeString, value)
		_node.Avatar = value
	}
	if value, ok := suc.mutation.Sex(); ok {
		_spec.SetField(sysuser.FieldSex, field.TypeEnum, value)
		_node.Sex = value
	}
	if value, ok := suc.mutation.Status(); ok {
		_spec.SetField(sysuser.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := suc.mutation.Password(); ok {
		_spec.SetField(sysuser.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := suc.mutation.Remark(); ok {
		_spec.SetField(sysuser.FieldRemark, field.TypeString, value)
		_node.Remark = value
	}
	if nodes := suc.mutation.UsersPostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sysuser.UsersPostTable,
			Columns: sysuser.UsersPostPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(syspost.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &SysUserPostCreate{config: suc.config, mutation: newSysUserPostMutation(suc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SysUserCreateBulk is the builder for creating many SysUser entities in bulk.
type SysUserCreateBulk struct {
	config
	builders []*SysUserCreate
}

// Save creates the SysUser entities in the database.
func (sucb *SysUserCreateBulk) Save(ctx context.Context) ([]*SysUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sucb.builders))
	nodes := make([]*SysUser, len(sucb.builders))
	mutators := make([]Mutator, len(sucb.builders))
	for i := range sucb.builders {
		func(i int, root context.Context) {
			builder := sucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysUserMutation)
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
					_, err = mutators[i+1].Mutate(root, sucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sucb *SysUserCreateBulk) SaveX(ctx context.Context) []*SysUser {
	v, err := sucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sucb *SysUserCreateBulk) Exec(ctx context.Context) error {
	_, err := sucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sucb *SysUserCreateBulk) ExecX(ctx context.Context) {
	if err := sucb.Exec(ctx); err != nil {
		panic(err)
	}
}

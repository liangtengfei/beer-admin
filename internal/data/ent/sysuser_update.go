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
)

// SysUserUpdate is the builder for updating SysUser entities.
type SysUserUpdate struct {
	config
	hooks    []Hook
	mutation *SysUserMutation
}

// Where appends a list predicates to the SysUserUpdate builder.
func (suu *SysUserUpdate) Where(ps ...predicate.SysUser) *SysUserUpdate {
	suu.mutation.Where(ps...)
	return suu
}

// SetUpdatedAt sets the "updated_at" field.
func (suu *SysUserUpdate) SetUpdatedAt(t time.Time) *SysUserUpdate {
	suu.mutation.SetUpdatedAt(t)
	return suu
}

// SetDeletedAt sets the "deleted_at" field.
func (suu *SysUserUpdate) SetDeletedAt(t time.Time) *SysUserUpdate {
	suu.mutation.SetDeletedAt(t)
	return suu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (suu *SysUserUpdate) SetNillableDeletedAt(t *time.Time) *SysUserUpdate {
	if t != nil {
		suu.SetDeletedAt(*t)
	}
	return suu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (suu *SysUserUpdate) ClearDeletedAt() *SysUserUpdate {
	suu.mutation.ClearDeletedAt()
	return suu
}

// SetCreatedBy sets the "created_by" field.
func (suu *SysUserUpdate) SetCreatedBy(s string) *SysUserUpdate {
	suu.mutation.SetCreatedBy(s)
	return suu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (suu *SysUserUpdate) SetNillableCreatedBy(s *string) *SysUserUpdate {
	if s != nil {
		suu.SetCreatedBy(*s)
	}
	return suu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (suu *SysUserUpdate) ClearCreatedBy() *SysUserUpdate {
	suu.mutation.ClearCreatedBy()
	return suu
}

// SetUpdatedBy sets the "updated_by" field.
func (suu *SysUserUpdate) SetUpdatedBy(s string) *SysUserUpdate {
	suu.mutation.SetUpdatedBy(s)
	return suu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (suu *SysUserUpdate) SetNillableUpdatedBy(s *string) *SysUserUpdate {
	if s != nil {
		suu.SetUpdatedBy(*s)
	}
	return suu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (suu *SysUserUpdate) ClearUpdatedBy() *SysUserUpdate {
	suu.mutation.ClearUpdatedBy()
	return suu
}

// SetDeptID sets the "dept_id" field.
func (suu *SysUserUpdate) SetDeptID(u uint64) *SysUserUpdate {
	suu.mutation.ResetDeptID()
	suu.mutation.SetDeptID(u)
	return suu
}

// SetNillableDeptID sets the "dept_id" field if the given value is not nil.
func (suu *SysUserUpdate) SetNillableDeptID(u *uint64) *SysUserUpdate {
	if u != nil {
		suu.SetDeptID(*u)
	}
	return suu
}

// AddDeptID adds u to the "dept_id" field.
func (suu *SysUserUpdate) AddDeptID(u int64) *SysUserUpdate {
	suu.mutation.AddDeptID(u)
	return suu
}

// ClearDeptID clears the value of the "dept_id" field.
func (suu *SysUserUpdate) ClearDeptID() *SysUserUpdate {
	suu.mutation.ClearDeptID()
	return suu
}

// SetRealName sets the "real_name" field.
func (suu *SysUserUpdate) SetRealName(s string) *SysUserUpdate {
	suu.mutation.SetRealName(s)
	return suu
}

// SetMobile sets the "mobile" field.
func (suu *SysUserUpdate) SetMobile(s string) *SysUserUpdate {
	suu.mutation.SetMobile(s)
	return suu
}

// SetEmail sets the "email" field.
func (suu *SysUserUpdate) SetEmail(s string) *SysUserUpdate {
	suu.mutation.SetEmail(s)
	return suu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (suu *SysUserUpdate) SetNillableEmail(s *string) *SysUserUpdate {
	if s != nil {
		suu.SetEmail(*s)
	}
	return suu
}

// ClearEmail clears the value of the "email" field.
func (suu *SysUserUpdate) ClearEmail() *SysUserUpdate {
	suu.mutation.ClearEmail()
	return suu
}

// SetAvatar sets the "avatar" field.
func (suu *SysUserUpdate) SetAvatar(s string) *SysUserUpdate {
	suu.mutation.SetAvatar(s)
	return suu
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (suu *SysUserUpdate) SetNillableAvatar(s *string) *SysUserUpdate {
	if s != nil {
		suu.SetAvatar(*s)
	}
	return suu
}

// ClearAvatar clears the value of the "avatar" field.
func (suu *SysUserUpdate) ClearAvatar() *SysUserUpdate {
	suu.mutation.ClearAvatar()
	return suu
}

// SetSex sets the "sex" field.
func (suu *SysUserUpdate) SetSex(s sysuser.Sex) *SysUserUpdate {
	suu.mutation.SetSex(s)
	return suu
}

// SetNillableSex sets the "sex" field if the given value is not nil.
func (suu *SysUserUpdate) SetNillableSex(s *sysuser.Sex) *SysUserUpdate {
	if s != nil {
		suu.SetSex(*s)
	}
	return suu
}

// SetStatus sets the "status" field.
func (suu *SysUserUpdate) SetStatus(s sysuser.Status) *SysUserUpdate {
	suu.mutation.SetStatus(s)
	return suu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (suu *SysUserUpdate) SetNillableStatus(s *sysuser.Status) *SysUserUpdate {
	if s != nil {
		suu.SetStatus(*s)
	}
	return suu
}

// SetPassword sets the "password" field.
func (suu *SysUserUpdate) SetPassword(s string) *SysUserUpdate {
	suu.mutation.SetPassword(s)
	return suu
}

// SetRemark sets the "remark" field.
func (suu *SysUserUpdate) SetRemark(s string) *SysUserUpdate {
	suu.mutation.SetRemark(s)
	return suu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (suu *SysUserUpdate) SetNillableRemark(s *string) *SysUserUpdate {
	if s != nil {
		suu.SetRemark(*s)
	}
	return suu
}

// ClearRemark clears the value of the "remark" field.
func (suu *SysUserUpdate) ClearRemark() *SysUserUpdate {
	suu.mutation.ClearRemark()
	return suu
}

// AddUsersPostIDs adds the "users_post" edge to the SysPost entity by IDs.
func (suu *SysUserUpdate) AddUsersPostIDs(ids ...uint64) *SysUserUpdate {
	suu.mutation.AddUsersPostIDs(ids...)
	return suu
}

// AddUsersPost adds the "users_post" edges to the SysPost entity.
func (suu *SysUserUpdate) AddUsersPost(s ...*SysPost) *SysUserUpdate {
	ids := make([]uint64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suu.AddUsersPostIDs(ids...)
}

// Mutation returns the SysUserMutation object of the builder.
func (suu *SysUserUpdate) Mutation() *SysUserMutation {
	return suu.mutation
}

// ClearUsersPost clears all "users_post" edges to the SysPost entity.
func (suu *SysUserUpdate) ClearUsersPost() *SysUserUpdate {
	suu.mutation.ClearUsersPost()
	return suu
}

// RemoveUsersPostIDs removes the "users_post" edge to SysPost entities by IDs.
func (suu *SysUserUpdate) RemoveUsersPostIDs(ids ...uint64) *SysUserUpdate {
	suu.mutation.RemoveUsersPostIDs(ids...)
	return suu
}

// RemoveUsersPost removes "users_post" edges to SysPost entities.
func (suu *SysUserUpdate) RemoveUsersPost(s ...*SysPost) *SysUserUpdate {
	ids := make([]uint64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suu.RemoveUsersPostIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (suu *SysUserUpdate) Save(ctx context.Context) (int, error) {
	suu.defaults()
	return withHooks[int, SysUserMutation](ctx, suu.sqlSave, suu.mutation, suu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suu *SysUserUpdate) SaveX(ctx context.Context) int {
	affected, err := suu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (suu *SysUserUpdate) Exec(ctx context.Context) error {
	_, err := suu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suu *SysUserUpdate) ExecX(ctx context.Context) {
	if err := suu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suu *SysUserUpdate) defaults() {
	if _, ok := suu.mutation.UpdatedAt(); !ok {
		v := sysuser.UpdateDefaultUpdatedAt()
		suu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suu *SysUserUpdate) check() error {
	if v, ok := suu.mutation.RealName(); ok {
		if err := sysuser.RealNameValidator(v); err != nil {
			return &ValidationError{Name: "real_name", err: fmt.Errorf(`ent: validator failed for field "SysUser.real_name": %w`, err)}
		}
	}
	if v, ok := suu.mutation.Mobile(); ok {
		if err := sysuser.MobileValidator(v); err != nil {
			return &ValidationError{Name: "mobile", err: fmt.Errorf(`ent: validator failed for field "SysUser.mobile": %w`, err)}
		}
	}
	if v, ok := suu.mutation.Sex(); ok {
		if err := sysuser.SexValidator(v); err != nil {
			return &ValidationError{Name: "sex", err: fmt.Errorf(`ent: validator failed for field "SysUser.sex": %w`, err)}
		}
	}
	if v, ok := suu.mutation.Status(); ok {
		if err := sysuser.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "SysUser.status": %w`, err)}
		}
	}
	if v, ok := suu.mutation.Password(); ok {
		if err := sysuser.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "SysUser.password": %w`, err)}
		}
	}
	return nil
}

func (suu *SysUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := suu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(sysuser.Table, sysuser.Columns, sqlgraph.NewFieldSpec(sysuser.FieldID, field.TypeUint64))
	if ps := suu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suu.mutation.UpdatedAt(); ok {
		_spec.SetField(sysuser.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := suu.mutation.DeletedAt(); ok {
		_spec.SetField(sysuser.FieldDeletedAt, field.TypeTime, value)
	}
	if suu.mutation.DeletedAtCleared() {
		_spec.ClearField(sysuser.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := suu.mutation.CreatedBy(); ok {
		_spec.SetField(sysuser.FieldCreatedBy, field.TypeString, value)
	}
	if suu.mutation.CreatedByCleared() {
		_spec.ClearField(sysuser.FieldCreatedBy, field.TypeString)
	}
	if value, ok := suu.mutation.UpdatedBy(); ok {
		_spec.SetField(sysuser.FieldUpdatedBy, field.TypeString, value)
	}
	if suu.mutation.UpdatedByCleared() {
		_spec.ClearField(sysuser.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := suu.mutation.DeptID(); ok {
		_spec.SetField(sysuser.FieldDeptID, field.TypeUint64, value)
	}
	if value, ok := suu.mutation.AddedDeptID(); ok {
		_spec.AddField(sysuser.FieldDeptID, field.TypeUint64, value)
	}
	if suu.mutation.DeptIDCleared() {
		_spec.ClearField(sysuser.FieldDeptID, field.TypeUint64)
	}
	if value, ok := suu.mutation.RealName(); ok {
		_spec.SetField(sysuser.FieldRealName, field.TypeString, value)
	}
	if value, ok := suu.mutation.Mobile(); ok {
		_spec.SetField(sysuser.FieldMobile, field.TypeString, value)
	}
	if value, ok := suu.mutation.Email(); ok {
		_spec.SetField(sysuser.FieldEmail, field.TypeString, value)
	}
	if suu.mutation.EmailCleared() {
		_spec.ClearField(sysuser.FieldEmail, field.TypeString)
	}
	if value, ok := suu.mutation.Avatar(); ok {
		_spec.SetField(sysuser.FieldAvatar, field.TypeString, value)
	}
	if suu.mutation.AvatarCleared() {
		_spec.ClearField(sysuser.FieldAvatar, field.TypeString)
	}
	if value, ok := suu.mutation.Sex(); ok {
		_spec.SetField(sysuser.FieldSex, field.TypeEnum, value)
	}
	if value, ok := suu.mutation.Status(); ok {
		_spec.SetField(sysuser.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := suu.mutation.Password(); ok {
		_spec.SetField(sysuser.FieldPassword, field.TypeString, value)
	}
	if value, ok := suu.mutation.Remark(); ok {
		_spec.SetField(sysuser.FieldRemark, field.TypeString, value)
	}
	if suu.mutation.RemarkCleared() {
		_spec.ClearField(sysuser.FieldRemark, field.TypeString)
	}
	if suu.mutation.UsersPostCleared() {
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
		createE := &SysUserPostCreate{config: suu.config, mutation: newSysUserPostMutation(suu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suu.mutation.RemovedUsersPostIDs(); len(nodes) > 0 && !suu.mutation.UsersPostCleared() {
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
		createE := &SysUserPostCreate{config: suu.config, mutation: newSysUserPostMutation(suu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suu.mutation.UsersPostIDs(); len(nodes) > 0 {
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
		createE := &SysUserPostCreate{config: suu.config, mutation: newSysUserPostMutation(suu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, suu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	suu.mutation.done = true
	return n, nil
}

// SysUserUpdateOne is the builder for updating a single SysUser entity.
type SysUserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SysUserMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (suuo *SysUserUpdateOne) SetUpdatedAt(t time.Time) *SysUserUpdateOne {
	suuo.mutation.SetUpdatedAt(t)
	return suuo
}

// SetDeletedAt sets the "deleted_at" field.
func (suuo *SysUserUpdateOne) SetDeletedAt(t time.Time) *SysUserUpdateOne {
	suuo.mutation.SetDeletedAt(t)
	return suuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (suuo *SysUserUpdateOne) SetNillableDeletedAt(t *time.Time) *SysUserUpdateOne {
	if t != nil {
		suuo.SetDeletedAt(*t)
	}
	return suuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (suuo *SysUserUpdateOne) ClearDeletedAt() *SysUserUpdateOne {
	suuo.mutation.ClearDeletedAt()
	return suuo
}

// SetCreatedBy sets the "created_by" field.
func (suuo *SysUserUpdateOne) SetCreatedBy(s string) *SysUserUpdateOne {
	suuo.mutation.SetCreatedBy(s)
	return suuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (suuo *SysUserUpdateOne) SetNillableCreatedBy(s *string) *SysUserUpdateOne {
	if s != nil {
		suuo.SetCreatedBy(*s)
	}
	return suuo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (suuo *SysUserUpdateOne) ClearCreatedBy() *SysUserUpdateOne {
	suuo.mutation.ClearCreatedBy()
	return suuo
}

// SetUpdatedBy sets the "updated_by" field.
func (suuo *SysUserUpdateOne) SetUpdatedBy(s string) *SysUserUpdateOne {
	suuo.mutation.SetUpdatedBy(s)
	return suuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (suuo *SysUserUpdateOne) SetNillableUpdatedBy(s *string) *SysUserUpdateOne {
	if s != nil {
		suuo.SetUpdatedBy(*s)
	}
	return suuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (suuo *SysUserUpdateOne) ClearUpdatedBy() *SysUserUpdateOne {
	suuo.mutation.ClearUpdatedBy()
	return suuo
}

// SetDeptID sets the "dept_id" field.
func (suuo *SysUserUpdateOne) SetDeptID(u uint64) *SysUserUpdateOne {
	suuo.mutation.ResetDeptID()
	suuo.mutation.SetDeptID(u)
	return suuo
}

// SetNillableDeptID sets the "dept_id" field if the given value is not nil.
func (suuo *SysUserUpdateOne) SetNillableDeptID(u *uint64) *SysUserUpdateOne {
	if u != nil {
		suuo.SetDeptID(*u)
	}
	return suuo
}

// AddDeptID adds u to the "dept_id" field.
func (suuo *SysUserUpdateOne) AddDeptID(u int64) *SysUserUpdateOne {
	suuo.mutation.AddDeptID(u)
	return suuo
}

// ClearDeptID clears the value of the "dept_id" field.
func (suuo *SysUserUpdateOne) ClearDeptID() *SysUserUpdateOne {
	suuo.mutation.ClearDeptID()
	return suuo
}

// SetRealName sets the "real_name" field.
func (suuo *SysUserUpdateOne) SetRealName(s string) *SysUserUpdateOne {
	suuo.mutation.SetRealName(s)
	return suuo
}

// SetMobile sets the "mobile" field.
func (suuo *SysUserUpdateOne) SetMobile(s string) *SysUserUpdateOne {
	suuo.mutation.SetMobile(s)
	return suuo
}

// SetEmail sets the "email" field.
func (suuo *SysUserUpdateOne) SetEmail(s string) *SysUserUpdateOne {
	suuo.mutation.SetEmail(s)
	return suuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (suuo *SysUserUpdateOne) SetNillableEmail(s *string) *SysUserUpdateOne {
	if s != nil {
		suuo.SetEmail(*s)
	}
	return suuo
}

// ClearEmail clears the value of the "email" field.
func (suuo *SysUserUpdateOne) ClearEmail() *SysUserUpdateOne {
	suuo.mutation.ClearEmail()
	return suuo
}

// SetAvatar sets the "avatar" field.
func (suuo *SysUserUpdateOne) SetAvatar(s string) *SysUserUpdateOne {
	suuo.mutation.SetAvatar(s)
	return suuo
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (suuo *SysUserUpdateOne) SetNillableAvatar(s *string) *SysUserUpdateOne {
	if s != nil {
		suuo.SetAvatar(*s)
	}
	return suuo
}

// ClearAvatar clears the value of the "avatar" field.
func (suuo *SysUserUpdateOne) ClearAvatar() *SysUserUpdateOne {
	suuo.mutation.ClearAvatar()
	return suuo
}

// SetSex sets the "sex" field.
func (suuo *SysUserUpdateOne) SetSex(s sysuser.Sex) *SysUserUpdateOne {
	suuo.mutation.SetSex(s)
	return suuo
}

// SetNillableSex sets the "sex" field if the given value is not nil.
func (suuo *SysUserUpdateOne) SetNillableSex(s *sysuser.Sex) *SysUserUpdateOne {
	if s != nil {
		suuo.SetSex(*s)
	}
	return suuo
}

// SetStatus sets the "status" field.
func (suuo *SysUserUpdateOne) SetStatus(s sysuser.Status) *SysUserUpdateOne {
	suuo.mutation.SetStatus(s)
	return suuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (suuo *SysUserUpdateOne) SetNillableStatus(s *sysuser.Status) *SysUserUpdateOne {
	if s != nil {
		suuo.SetStatus(*s)
	}
	return suuo
}

// SetPassword sets the "password" field.
func (suuo *SysUserUpdateOne) SetPassword(s string) *SysUserUpdateOne {
	suuo.mutation.SetPassword(s)
	return suuo
}

// SetRemark sets the "remark" field.
func (suuo *SysUserUpdateOne) SetRemark(s string) *SysUserUpdateOne {
	suuo.mutation.SetRemark(s)
	return suuo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (suuo *SysUserUpdateOne) SetNillableRemark(s *string) *SysUserUpdateOne {
	if s != nil {
		suuo.SetRemark(*s)
	}
	return suuo
}

// ClearRemark clears the value of the "remark" field.
func (suuo *SysUserUpdateOne) ClearRemark() *SysUserUpdateOne {
	suuo.mutation.ClearRemark()
	return suuo
}

// AddUsersPostIDs adds the "users_post" edge to the SysPost entity by IDs.
func (suuo *SysUserUpdateOne) AddUsersPostIDs(ids ...uint64) *SysUserUpdateOne {
	suuo.mutation.AddUsersPostIDs(ids...)
	return suuo
}

// AddUsersPost adds the "users_post" edges to the SysPost entity.
func (suuo *SysUserUpdateOne) AddUsersPost(s ...*SysPost) *SysUserUpdateOne {
	ids := make([]uint64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suuo.AddUsersPostIDs(ids...)
}

// Mutation returns the SysUserMutation object of the builder.
func (suuo *SysUserUpdateOne) Mutation() *SysUserMutation {
	return suuo.mutation
}

// ClearUsersPost clears all "users_post" edges to the SysPost entity.
func (suuo *SysUserUpdateOne) ClearUsersPost() *SysUserUpdateOne {
	suuo.mutation.ClearUsersPost()
	return suuo
}

// RemoveUsersPostIDs removes the "users_post" edge to SysPost entities by IDs.
func (suuo *SysUserUpdateOne) RemoveUsersPostIDs(ids ...uint64) *SysUserUpdateOne {
	suuo.mutation.RemoveUsersPostIDs(ids...)
	return suuo
}

// RemoveUsersPost removes "users_post" edges to SysPost entities.
func (suuo *SysUserUpdateOne) RemoveUsersPost(s ...*SysPost) *SysUserUpdateOne {
	ids := make([]uint64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suuo.RemoveUsersPostIDs(ids...)
}

// Where appends a list predicates to the SysUserUpdate builder.
func (suuo *SysUserUpdateOne) Where(ps ...predicate.SysUser) *SysUserUpdateOne {
	suuo.mutation.Where(ps...)
	return suuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suuo *SysUserUpdateOne) Select(field string, fields ...string) *SysUserUpdateOne {
	suuo.fields = append([]string{field}, fields...)
	return suuo
}

// Save executes the query and returns the updated SysUser entity.
func (suuo *SysUserUpdateOne) Save(ctx context.Context) (*SysUser, error) {
	suuo.defaults()
	return withHooks[*SysUser, SysUserMutation](ctx, suuo.sqlSave, suuo.mutation, suuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suuo *SysUserUpdateOne) SaveX(ctx context.Context) *SysUser {
	node, err := suuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suuo *SysUserUpdateOne) Exec(ctx context.Context) error {
	_, err := suuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suuo *SysUserUpdateOne) ExecX(ctx context.Context) {
	if err := suuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suuo *SysUserUpdateOne) defaults() {
	if _, ok := suuo.mutation.UpdatedAt(); !ok {
		v := sysuser.UpdateDefaultUpdatedAt()
		suuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suuo *SysUserUpdateOne) check() error {
	if v, ok := suuo.mutation.RealName(); ok {
		if err := sysuser.RealNameValidator(v); err != nil {
			return &ValidationError{Name: "real_name", err: fmt.Errorf(`ent: validator failed for field "SysUser.real_name": %w`, err)}
		}
	}
	if v, ok := suuo.mutation.Mobile(); ok {
		if err := sysuser.MobileValidator(v); err != nil {
			return &ValidationError{Name: "mobile", err: fmt.Errorf(`ent: validator failed for field "SysUser.mobile": %w`, err)}
		}
	}
	if v, ok := suuo.mutation.Sex(); ok {
		if err := sysuser.SexValidator(v); err != nil {
			return &ValidationError{Name: "sex", err: fmt.Errorf(`ent: validator failed for field "SysUser.sex": %w`, err)}
		}
	}
	if v, ok := suuo.mutation.Status(); ok {
		if err := sysuser.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "SysUser.status": %w`, err)}
		}
	}
	if v, ok := suuo.mutation.Password(); ok {
		if err := sysuser.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "SysUser.password": %w`, err)}
		}
	}
	return nil
}

func (suuo *SysUserUpdateOne) sqlSave(ctx context.Context) (_node *SysUser, err error) {
	if err := suuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(sysuser.Table, sysuser.Columns, sqlgraph.NewFieldSpec(sysuser.FieldID, field.TypeUint64))
	id, ok := suuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SysUser.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysuser.FieldID)
		for _, f := range fields {
			if !sysuser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sysuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suuo.mutation.UpdatedAt(); ok {
		_spec.SetField(sysuser.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := suuo.mutation.DeletedAt(); ok {
		_spec.SetField(sysuser.FieldDeletedAt, field.TypeTime, value)
	}
	if suuo.mutation.DeletedAtCleared() {
		_spec.ClearField(sysuser.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := suuo.mutation.CreatedBy(); ok {
		_spec.SetField(sysuser.FieldCreatedBy, field.TypeString, value)
	}
	if suuo.mutation.CreatedByCleared() {
		_spec.ClearField(sysuser.FieldCreatedBy, field.TypeString)
	}
	if value, ok := suuo.mutation.UpdatedBy(); ok {
		_spec.SetField(sysuser.FieldUpdatedBy, field.TypeString, value)
	}
	if suuo.mutation.UpdatedByCleared() {
		_spec.ClearField(sysuser.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := suuo.mutation.DeptID(); ok {
		_spec.SetField(sysuser.FieldDeptID, field.TypeUint64, value)
	}
	if value, ok := suuo.mutation.AddedDeptID(); ok {
		_spec.AddField(sysuser.FieldDeptID, field.TypeUint64, value)
	}
	if suuo.mutation.DeptIDCleared() {
		_spec.ClearField(sysuser.FieldDeptID, field.TypeUint64)
	}
	if value, ok := suuo.mutation.RealName(); ok {
		_spec.SetField(sysuser.FieldRealName, field.TypeString, value)
	}
	if value, ok := suuo.mutation.Mobile(); ok {
		_spec.SetField(sysuser.FieldMobile, field.TypeString, value)
	}
	if value, ok := suuo.mutation.Email(); ok {
		_spec.SetField(sysuser.FieldEmail, field.TypeString, value)
	}
	if suuo.mutation.EmailCleared() {
		_spec.ClearField(sysuser.FieldEmail, field.TypeString)
	}
	if value, ok := suuo.mutation.Avatar(); ok {
		_spec.SetField(sysuser.FieldAvatar, field.TypeString, value)
	}
	if suuo.mutation.AvatarCleared() {
		_spec.ClearField(sysuser.FieldAvatar, field.TypeString)
	}
	if value, ok := suuo.mutation.Sex(); ok {
		_spec.SetField(sysuser.FieldSex, field.TypeEnum, value)
	}
	if value, ok := suuo.mutation.Status(); ok {
		_spec.SetField(sysuser.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := suuo.mutation.Password(); ok {
		_spec.SetField(sysuser.FieldPassword, field.TypeString, value)
	}
	if value, ok := suuo.mutation.Remark(); ok {
		_spec.SetField(sysuser.FieldRemark, field.TypeString, value)
	}
	if suuo.mutation.RemarkCleared() {
		_spec.ClearField(sysuser.FieldRemark, field.TypeString)
	}
	if suuo.mutation.UsersPostCleared() {
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
		createE := &SysUserPostCreate{config: suuo.config, mutation: newSysUserPostMutation(suuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suuo.mutation.RemovedUsersPostIDs(); len(nodes) > 0 && !suuo.mutation.UsersPostCleared() {
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
		createE := &SysUserPostCreate{config: suuo.config, mutation: newSysUserPostMutation(suuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suuo.mutation.UsersPostIDs(); len(nodes) > 0 {
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
		createE := &SysUserPostCreate{config: suuo.config, mutation: newSysUserPostMutation(suuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SysUser{config: suuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suuo.mutation.done = true
	return _node, nil
}

// Code generated by ent, DO NOT EDIT.

package syspost

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/liangtengfei/beer-admin/internal/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.SysPost {
	return predicate.SysPost(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.SysPost {
	return predicate.SysPost(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.SysPost {
	return predicate.SysPost(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.SysPost {
	return predicate.SysPost(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.SysPost {
	return predicate.SysPost(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.SysPost {
	return predicate.SysPost(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.SysPost {
	return predicate.SysPost(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldDeletedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldUpdatedBy, v))
}

// PostName applies equality check predicate on the "post_name" field. It's identical to PostNameEQ.
func PostName(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldPostName, v))
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldRemark, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.SysPost {
	return predicate.SysPost(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.SysPost {
	return predicate.SysPost(sql.FieldNotNull(FieldDeletedAt))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.SysPost {
	return predicate.SysPost(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.SysPost {
	return predicate.SysPost(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.SysPost {
	return predicate.SysPost(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.SysPost {
	return predicate.SysPost(sql.FieldNotNull(FieldCreatedBy))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.SysPost {
	return predicate.SysPost(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.SysPost {
	return predicate.SysPost(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.SysPost {
	return predicate.SysPost(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.SysPost {
	return predicate.SysPost(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// PostNameEQ applies the EQ predicate on the "post_name" field.
func PostNameEQ(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldPostName, v))
}

// PostNameNEQ applies the NEQ predicate on the "post_name" field.
func PostNameNEQ(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldNEQ(FieldPostName, v))
}

// PostNameIn applies the In predicate on the "post_name" field.
func PostNameIn(vs ...string) predicate.SysPost {
	return predicate.SysPost(sql.FieldIn(FieldPostName, vs...))
}

// PostNameNotIn applies the NotIn predicate on the "post_name" field.
func PostNameNotIn(vs ...string) predicate.SysPost {
	return predicate.SysPost(sql.FieldNotIn(FieldPostName, vs...))
}

// PostNameGT applies the GT predicate on the "post_name" field.
func PostNameGT(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldGT(FieldPostName, v))
}

// PostNameGTE applies the GTE predicate on the "post_name" field.
func PostNameGTE(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldGTE(FieldPostName, v))
}

// PostNameLT applies the LT predicate on the "post_name" field.
func PostNameLT(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldLT(FieldPostName, v))
}

// PostNameLTE applies the LTE predicate on the "post_name" field.
func PostNameLTE(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldLTE(FieldPostName, v))
}

// PostNameContains applies the Contains predicate on the "post_name" field.
func PostNameContains(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldContains(FieldPostName, v))
}

// PostNameHasPrefix applies the HasPrefix predicate on the "post_name" field.
func PostNameHasPrefix(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldHasPrefix(FieldPostName, v))
}

// PostNameHasSuffix applies the HasSuffix predicate on the "post_name" field.
func PostNameHasSuffix(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldHasSuffix(FieldPostName, v))
}

// PostNameEqualFold applies the EqualFold predicate on the "post_name" field.
func PostNameEqualFold(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldEqualFold(FieldPostName, v))
}

// PostNameContainsFold applies the ContainsFold predicate on the "post_name" field.
func PostNameContainsFold(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldContainsFold(FieldPostName, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.SysPost {
	return predicate.SysPost(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.SysPost {
	return predicate.SysPost(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.SysPost {
	return predicate.SysPost(sql.FieldNotIn(FieldStatus, vs...))
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldRemark, v))
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldNEQ(FieldRemark, v))
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.SysPost {
	return predicate.SysPost(sql.FieldIn(FieldRemark, vs...))
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.SysPost {
	return predicate.SysPost(sql.FieldNotIn(FieldRemark, vs...))
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldGT(FieldRemark, v))
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldGTE(FieldRemark, v))
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldLT(FieldRemark, v))
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldLTE(FieldRemark, v))
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldContains(FieldRemark, v))
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldHasPrefix(FieldRemark, v))
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldHasSuffix(FieldRemark, v))
}

// RemarkIsNil applies the IsNil predicate on the "remark" field.
func RemarkIsNil() predicate.SysPost {
	return predicate.SysPost(sql.FieldIsNull(FieldRemark))
}

// RemarkNotNil applies the NotNil predicate on the "remark" field.
func RemarkNotNil() predicate.SysPost {
	return predicate.SysPost(sql.FieldNotNull(FieldRemark))
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldEqualFold(FieldRemark, v))
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldContainsFold(FieldRemark, v))
}

// HasPostsUser applies the HasEdge predicate on the "posts_user" edge.
func HasPostsUser() predicate.SysPost {
	return predicate.SysPost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, PostsUserTable, PostsUserPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPostsUserWith applies the HasEdge predicate on the "posts_user" edge with a given conditions (other predicates).
func HasPostsUserWith(preds ...predicate.SysUser) predicate.SysPost {
	return predicate.SysPost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PostsUserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, PostsUserTable, PostsUserPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSysUserPost applies the HasEdge predicate on the "sys_user_post" edge.
func HasSysUserPost() predicate.SysPost {
	return predicate.SysPost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, SysUserPostTable, SysUserPostColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSysUserPostWith applies the HasEdge predicate on the "sys_user_post" edge with a given conditions (other predicates).
func HasSysUserPostWith(preds ...predicate.SysUserPost) predicate.SysPost {
	return predicate.SysPost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SysUserPostInverseTable, SysUserPostColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, SysUserPostTable, SysUserPostColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SysPost) predicate.SysPost {
	return predicate.SysPost(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SysPost) predicate.SysPost {
	return predicate.SysPost(func(s *sql.Selector) {
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
func Not(p predicate.SysPost) predicate.SysPost {
	return predicate.SysPost(func(s *sql.Selector) {
		p(s.Not())
	})
}

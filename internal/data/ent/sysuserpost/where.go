// Code generated by ent, DO NOT EDIT.

package sysuserpost

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/liangtengfei/beer-admin/internal/data/ent/predicate"
)

// SysUserID applies equality check predicate on the "sys_user_id" field. It's identical to SysUserIDEQ.
func SysUserID(v uint64) predicate.SysUserPost {
	return predicate.SysUserPost(sql.FieldEQ(FieldSysUserID, v))
}

// SysPostID applies equality check predicate on the "sys_post_id" field. It's identical to SysPostIDEQ.
func SysPostID(v uint64) predicate.SysUserPost {
	return predicate.SysUserPost(sql.FieldEQ(FieldSysPostID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SysUserPost {
	return predicate.SysUserPost(sql.FieldEQ(FieldCreatedAt, v))
}

// SysUserIDEQ applies the EQ predicate on the "sys_user_id" field.
func SysUserIDEQ(v uint64) predicate.SysUserPost {
	return predicate.SysUserPost(sql.FieldEQ(FieldSysUserID, v))
}

// SysUserIDNEQ applies the NEQ predicate on the "sys_user_id" field.
func SysUserIDNEQ(v uint64) predicate.SysUserPost {
	return predicate.SysUserPost(sql.FieldNEQ(FieldSysUserID, v))
}

// SysUserIDIn applies the In predicate on the "sys_user_id" field.
func SysUserIDIn(vs ...uint64) predicate.SysUserPost {
	return predicate.SysUserPost(sql.FieldIn(FieldSysUserID, vs...))
}

// SysUserIDNotIn applies the NotIn predicate on the "sys_user_id" field.
func SysUserIDNotIn(vs ...uint64) predicate.SysUserPost {
	return predicate.SysUserPost(sql.FieldNotIn(FieldSysUserID, vs...))
}

// SysPostIDEQ applies the EQ predicate on the "sys_post_id" field.
func SysPostIDEQ(v uint64) predicate.SysUserPost {
	return predicate.SysUserPost(sql.FieldEQ(FieldSysPostID, v))
}

// SysPostIDNEQ applies the NEQ predicate on the "sys_post_id" field.
func SysPostIDNEQ(v uint64) predicate.SysUserPost {
	return predicate.SysUserPost(sql.FieldNEQ(FieldSysPostID, v))
}

// SysPostIDIn applies the In predicate on the "sys_post_id" field.
func SysPostIDIn(vs ...uint64) predicate.SysUserPost {
	return predicate.SysUserPost(sql.FieldIn(FieldSysPostID, vs...))
}

// SysPostIDNotIn applies the NotIn predicate on the "sys_post_id" field.
func SysPostIDNotIn(vs ...uint64) predicate.SysUserPost {
	return predicate.SysUserPost(sql.FieldNotIn(FieldSysPostID, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SysUserPost {
	return predicate.SysUserPost(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SysUserPost {
	return predicate.SysUserPost(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SysUserPost {
	return predicate.SysUserPost(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SysUserPost {
	return predicate.SysUserPost(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SysUserPost {
	return predicate.SysUserPost(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SysUserPost {
	return predicate.SysUserPost(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SysUserPost {
	return predicate.SysUserPost(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SysUserPost {
	return predicate.SysUserPost(sql.FieldLTE(FieldCreatedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.SysUserPost {
	return predicate.SysUserPost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, UserColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.SysUser) predicate.SysUserPost {
	return predicate.SysUserPost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, UserColumn),
			sqlgraph.To(UserInverseTable, SysUserFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPost applies the HasEdge predicate on the "post" edge.
func HasPost() predicate.SysUserPost {
	return predicate.SysUserPost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, PostColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, PostTable, PostColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPostWith applies the HasEdge predicate on the "post" edge with a given conditions (other predicates).
func HasPostWith(preds ...predicate.SysPost) predicate.SysUserPost {
	return predicate.SysUserPost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, PostColumn),
			sqlgraph.To(PostInverseTable, SysPostFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, PostTable, PostColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SysUserPost) predicate.SysUserPost {
	return predicate.SysUserPost(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SysUserPost) predicate.SysUserPost {
	return predicate.SysUserPost(func(s *sql.Selector) {
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
func Not(p predicate.SysUserPost) predicate.SysUserPost {
	return predicate.SysUserPost(func(s *sql.Selector) {
		p(s.Not())
	})
}

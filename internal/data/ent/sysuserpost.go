// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/liangtengfei/beer-admin/internal/data/ent/syspost"
	"github.com/liangtengfei/beer-admin/internal/data/ent/sysuser"
	"github.com/liangtengfei/beer-admin/internal/data/ent/sysuserpost"
)

// SysUserPost is the model entity for the SysUserPost schema.
type SysUserPost struct {
	config `json:"-"`
	// 用户标识
	SysUserID uint64 `json:"sys_user_id,omitempty"`
	// 岗位标识
	SysPostID uint64 `json:"sys_post_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SysUserPostQuery when eager-loading is set.
	Edges SysUserPostEdges `json:"edges"`
}

// SysUserPostEdges holds the relations/edges for other nodes in the graph.
type SysUserPostEdges struct {
	// User holds the value of the user edge.
	User *SysUser `json:"user,omitempty"`
	// Post holds the value of the post edge.
	Post *SysPost `json:"post,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SysUserPostEdges) UserOrErr() (*SysUser, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: sysuser.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// PostOrErr returns the Post value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SysUserPostEdges) PostOrErr() (*SysPost, error) {
	if e.loadedTypes[1] {
		if e.Post == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: syspost.Label}
		}
		return e.Post, nil
	}
	return nil, &NotLoadedError{edge: "post"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysUserPost) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysuserpost.FieldSysUserID, sysuserpost.FieldSysPostID:
			values[i] = new(sql.NullInt64)
		case sysuserpost.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SysUserPost", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysUserPost fields.
func (sup *SysUserPost) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sysuserpost.FieldSysUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sys_user_id", values[i])
			} else if value.Valid {
				sup.SysUserID = uint64(value.Int64)
			}
		case sysuserpost.FieldSysPostID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sys_post_id", values[i])
			} else if value.Valid {
				sup.SysPostID = uint64(value.Int64)
			}
		case sysuserpost.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sup.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the SysUserPost entity.
func (sup *SysUserPost) QueryUser() *SysUserQuery {
	return NewSysUserPostClient(sup.config).QueryUser(sup)
}

// QueryPost queries the "post" edge of the SysUserPost entity.
func (sup *SysUserPost) QueryPost() *SysPostQuery {
	return NewSysUserPostClient(sup.config).QueryPost(sup)
}

// Update returns a builder for updating this SysUserPost.
// Note that you need to call SysUserPost.Unwrap() before calling this method if this SysUserPost
// was returned from a transaction, and the transaction was committed or rolled back.
func (sup *SysUserPost) Update() *SysUserPostUpdateOne {
	return NewSysUserPostClient(sup.config).UpdateOne(sup)
}

// Unwrap unwraps the SysUserPost entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sup *SysUserPost) Unwrap() *SysUserPost {
	_tx, ok := sup.config.driver.(*txDriver)
	if !ok {
		panic("ent: SysUserPost is not a transactional entity")
	}
	sup.config.driver = _tx.drv
	return sup
}

// String implements the fmt.Stringer.
func (sup *SysUserPost) String() string {
	var builder strings.Builder
	builder.WriteString("SysUserPost(")
	builder.WriteString("sys_user_id=")
	builder.WriteString(fmt.Sprintf("%v", sup.SysUserID))
	builder.WriteString(", ")
	builder.WriteString("sys_post_id=")
	builder.WriteString(fmt.Sprintf("%v", sup.SysPostID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(sup.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// SysUserPosts is a parsable slice of SysUserPost.
type SysUserPosts []*SysUserPost

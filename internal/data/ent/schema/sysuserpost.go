package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SysUserPost holds the schema definition for the SysUserPost entity.
type SysUserPost struct {
	ent.Schema
}

func (SysUserPost) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("sys_user_id", "sys_post_id"),
		entsql.Annotation{
			Table: "sys_user_post",
		},
		entsql.WithComments(true),
	}
}

// Fields of the SysUserPost.
func (SysUserPost) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("sys_user_id").Comment("用户标识"),
		field.Uint64("sys_post_id").Comment("岗位标识"),
		field.Time("created_at").Default(time.Now()),
	}
}

// Edges of the SysUserPost.
func (SysUserPost) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", SysUser.Type).Unique().Required().Field("sys_user_id"),
		edge.To("post", SysPost.Type).Unique().Required().Field("sys_post_id"),
	}
}

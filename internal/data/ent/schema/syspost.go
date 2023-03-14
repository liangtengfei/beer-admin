package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SysPost holds the schema definition for the SysPost entity.
type SysPost struct {
	ent.Schema
}

func (SysPost) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_post"},
		entsql.WithComments(true),
	}
}

// Fields of the SysPost.
func (SysPost) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").StructTag(`json:"id"`).Comment("系统岗位"),
		field.String("post_name").NotEmpty().MinLen(3).MaxLen(100).Comment("岗位名称"),
		field.Enum("status").Values("0", "1").Default("0").Comment("状态: 0=ON 1=OFF"),
		field.String("remark").Optional().Comment("备注"),
	}
}

func (SysPost) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		OperatorMixin{},
	}
}

// Edges of the SysPost.
func (SysPost) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("posts_user", SysUser.Type).
			Ref("users_post").
			Through("sys_user_post", SysUserPost.Type),
	}
}

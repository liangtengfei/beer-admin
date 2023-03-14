package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SysUser holds the schema definition for the SysUser entity.
type SysUser struct {
	ent.Schema
}

// 一些扩展选项，可以自定义表名、以及表注释
func (SysUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:   "sys_user",
			Options: "COMMENT = '系统用户表'",
		},
		entsql.WithComments(true),
	}
}

// Fields of the SysUser.
func (SysUser) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").StructTag(`json:"id"`).Comment("系统用户"),
		field.Uint64("dept_id").Optional().Default(uint64(0)).Comment("部门编号"),
		field.String("user_name").Unique().Immutable().NotEmpty().MinLen(3).MaxLen(100).Comment("登陆用户名称"),
		field.String("real_name").NotEmpty().Comment("真实姓名"),
		field.String("mobile").NotEmpty().Comment("手机号码"),
		field.String("email").Optional().Comment("邮箱地址"),
		field.String("avatar").Optional().Comment("头像地址"),
		field.Enum("sex").Values("male", "female").Default("male").Comment("性别:male男性，female女性"),
		field.Enum("status").Values("0", "1").Default("0").Comment("状态: 0=ON 1=OFF"),
		field.String("password").Sensitive().MinLen(6).MaxLen(100).Comment("登陆密码"),
		field.String("remark").Optional().Comment("备注"),
	}
}

// 一些公用字段
func (SysUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		OperatorMixin{},
	}
}

// Edges of the SysUser.
func (SysUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users_post", SysPost.Type).Through("sys_user_post", SysUserPost.Type),
	}
}

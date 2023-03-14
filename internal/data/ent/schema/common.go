package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type TimeMixin struct {
	mixin.Schema
}

// Fields Immutable: 默认填充 不允许修改和手动填充
func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
		field.Time("deleted_at").Optional().Nillable().Comment("删除时间"),
	}
}

type OperatorMixin struct {
	mixin.Schema
}

func (OperatorMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("created_by").Optional().Default("").Comment("创建人员"),
		field.String("updated_by").Optional().Default("").Comment("更新人员"),
	}
}

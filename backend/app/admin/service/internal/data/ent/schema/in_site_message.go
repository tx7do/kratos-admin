package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/go-utils/entgo/mixin"
)

// InSiteMessage holds the schema definition for the InSiteMessage entity.
type InSiteMessage struct {
	ent.Schema
}

func (InSiteMessage) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "in_site_messages",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
	}
}

// Fields of the InSiteMessage.
func (InSiteMessage) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			Comment("标题").
			Optional().
			Nillable(),

		field.String("content").
			Comment("内容").
			Optional().
			Nillable(),

		field.Uint32("category_id").
			Comment("分类ID").
			Optional().
			Nillable(),

		field.Enum("status").
			Comment("消息状态").
			Values("Draft", "Published", "Scheduled", "Revoked", "Archived", "Unknown").
			Optional().
			Nillable(),
	}
}

// Mixin of the InSiteMessage.
func (InSiteMessage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Time{},
		mixin.CreateBy{},
		mixin.UpdateBy{},
		mixin.Remark{},
	}
}

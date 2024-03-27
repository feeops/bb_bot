package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Default("").Comment(""),
		field.String("emailPassword").Default("").Comment(""),
		field.String("bindEmail").Default("").Comment(""),
		field.String("postCode").Default("").Comment(""),
		field.String("password").Default("").Comment(""),
		field.String("firstName").Default("").Comment(""),
		field.String("lastName").Default("").Comment(""),
		field.String("region").Default("").Comment(""),
		field.String("IP").Default("").Comment(""),
		field.String("windowName").Default("").Comment("窗口名称"),
		field.Float("balance").Default(0).Comment("账号余额").
			SchemaType(map[string]string{
				dialect.MySQL:    "decimal(10,2)", // Override MySQL.
				dialect.Postgres: "numeric",       // Override Postgres.,
			}),
		field.Bool("IPUsed").Default(false).Comment(""),
		field.String("refURL").Default("").Comment("推广链接"),
		field.String("remark").Default("").Comment("备注"),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return nil
}

func (Account) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

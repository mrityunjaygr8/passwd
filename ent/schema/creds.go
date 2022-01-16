package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Creds holds the schema definition for the Creds entity.
type Creds struct {
	ent.Schema
}

func (Creds) Mixin() []ent.Mixin {
	return []ent.Mixin{mixin.Time{}}
}

func (Creds) Indexes() []ent.Index {
	return []ent.Index{index.Fields("name").Edges("user").Unique()}
}

// Fields of the Creds.
func (Creds) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("username"),
		field.String("url"),
	}
}

// Edges of the Creds.
func (Creds) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("creds").Unique(),
		edge.To("passwords", Passwords.Type),
	}
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Passwords holds the schema definition for the Passwords entity.
type Passwords struct {
	ent.Schema
}

func (Passwords) Mixin() []ent.Mixin {
	return []ent.Mixin{mixin.Time{}}
}

// Fields of the Passwords.
func (Passwords) Fields() []ent.Field {
	return []ent.Field{field.String("password").Sensitive()}
}

// Edges of the Passwords.
func (Passwords) Edges() []ent.Edge {
	return []ent.Edge{edge.From("cred", Creds.Type).Ref("passwords").Unique()}
}

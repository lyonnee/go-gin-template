package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique().Nillable(),
		field.String("name").Nillable(),
		field.Uint8("age").Nillable().Positive(),
		field.String("phone_number").Unique().Nillable(),
		field.String("auth_hash").Nillable(),
		field.Time("create_at").Default(time.Now),
		field.Time("update_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("phone_number").Unique(),
	}
}

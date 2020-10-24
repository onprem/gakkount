package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// OClient holds the schema definition for the OClient entity.
type OClient struct {
	ent.Schema
}

// Fields of the OClient.
func (OClient) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("clientID").Unique(),
		field.String("secret"),
	}
}

// Edges of the OClient.
func (OClient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique(),
	}
}

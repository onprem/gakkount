package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("email").Unique(),
		field.String("hash"),
		field.Enum("role").Values("student", "faculty", "staff", "admin").Default("student"),
		field.String("photo").Optional(),
		field.String("altEmail").Optional(),
		field.String("phone").Optional(),
		field.String("salutation").Optional(),
		// Social handles
		field.String("linkedin").Optional(),
		field.String("twitter").Optional(),
		field.String("facebook").Optional(),
		field.String("github").Optional(),
		// Student related fields
		field.String("rollNo").Optional(),
		field.Time("admissionTime").Optional(),
		field.Time("courseEndTime").Optional(),
		// Staff related fields
		field.String("Designation").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("course", Course.Type).Ref("users").Unique(),
		edge.From("department", Department.Type).Ref("users").Unique(),
	}
}

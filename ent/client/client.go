// Code generated by entc, DO NOT EDIT.

package client

const (
	// Label holds the string label denoting the client type in the database.
	Label = "client"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldClientID holds the string denoting the clientid field in the database.
	FieldClientID = "client_id"
	// FieldSecret holds the string denoting the secret field in the database.
	FieldSecret = "secret"

	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"

	// Table holds the table name of the client in the database.
	Table = "clients"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "clients"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "client_user"
)

// Columns holds all SQL columns for client fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldClientID,
	FieldSecret,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Client type.
var ForeignKeys = []string{
	"client_user",
}
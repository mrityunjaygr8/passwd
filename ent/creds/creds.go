// Code generated by entc, DO NOT EDIT.

package creds

import (
	"time"
)

const (
	// Label holds the string label denoting the creds type in the database.
	Label = "creds"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgePasswords holds the string denoting the passwords edge name in mutations.
	EdgePasswords = "passwords"
	// Table holds the table name of the creds in the database.
	Table = "creds"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "creds"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_creds"
	// PasswordsTable is the table that holds the passwords relation/edge.
	PasswordsTable = "passwords"
	// PasswordsInverseTable is the table name for the Passwords entity.
	// It exists in this package in order to avoid circular dependency with the "passwords" package.
	PasswordsInverseTable = "passwords"
	// PasswordsColumn is the table column denoting the passwords relation/edge.
	PasswordsColumn = "creds_passwords"
)

// Columns holds all SQL columns for creds fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
	FieldUsername,
	FieldURL,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "creds"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_creds",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)

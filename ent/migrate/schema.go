// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CredsColumns holds the columns for the "creds" table.
	CredsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "username", Type: field.TypeString},
		{Name: "url", Type: field.TypeString},
		{Name: "user_creds", Type: field.TypeInt, Nullable: true},
	}
	// CredsTable holds the schema information for the "creds" table.
	CredsTable = &schema.Table{
		Name:       "creds",
		Columns:    CredsColumns,
		PrimaryKey: []*schema.Column{CredsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "creds_users_creds",
				Columns:    []*schema.Column{CredsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "creds_name_user_creds",
				Unique:  true,
				Columns: []*schema.Column{CredsColumns[3], CredsColumns[6]},
			},
		},
	}
	// PasswordsColumns holds the columns for the "passwords" table.
	PasswordsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "password", Type: field.TypeString},
		{Name: "creds_passwords", Type: field.TypeInt, Nullable: true},
	}
	// PasswordsTable holds the schema information for the "passwords" table.
	PasswordsTable = &schema.Table{
		Name:       "passwords",
		Columns:    PasswordsColumns,
		PrimaryKey: []*schema.Column{PasswordsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "passwords_creds_passwords",
				Columns:    []*schema.Column{PasswordsColumns[4]},
				RefColumns: []*schema.Column{CredsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CredsTable,
		PasswordsTable,
		UsersTable,
	}
)

func init() {
	CredsTable.ForeignKeys[0].RefTable = UsersTable
	PasswordsTable.ForeignKeys[0].RefTable = CredsTable
}

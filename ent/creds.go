// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/mrtyunjaygr8/passwd/ent/creds"
	"github.com/mrtyunjaygr8/passwd/ent/user"
)

// Creds is the model entity for the Creds schema.
type Creds struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CredsQuery when eager-loading is set.
	Edges      CredsEdges `json:"edges"`
	user_creds *int
}

// CredsEdges holds the relations/edges for other nodes in the graph.
type CredsEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Passwords holds the value of the passwords edge.
	Passwords []*Passwords `json:"passwords,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CredsEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// PasswordsOrErr returns the Passwords value or an error if the edge
// was not loaded in eager-loading.
func (e CredsEdges) PasswordsOrErr() ([]*Passwords, error) {
	if e.loadedTypes[1] {
		return e.Passwords, nil
	}
	return nil, &NotLoadedError{edge: "passwords"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Creds) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case creds.FieldID:
			values[i] = new(sql.NullInt64)
		case creds.FieldName, creds.FieldUsername, creds.FieldURL:
			values[i] = new(sql.NullString)
		case creds.FieldCreateTime, creds.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case creds.ForeignKeys[0]: // user_creds
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Creds", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Creds fields.
func (c *Creds) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case creds.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case creds.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				c.CreateTime = value.Time
			}
		case creds.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				c.UpdateTime = value.Time
			}
		case creds.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case creds.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				c.Username = value.String
			}
		case creds.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				c.URL = value.String
			}
		case creds.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_creds", value)
			} else if value.Valid {
				c.user_creds = new(int)
				*c.user_creds = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Creds entity.
func (c *Creds) QueryUser() *UserQuery {
	return (&CredsClient{config: c.config}).QueryUser(c)
}

// QueryPasswords queries the "passwords" edge of the Creds entity.
func (c *Creds) QueryPasswords() *PasswordsQuery {
	return (&CredsClient{config: c.config}).QueryPasswords(c)
}

// Update returns a builder for updating this Creds.
// Note that you need to call Creds.Unwrap() before calling this method if this Creds
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Creds) Update() *CredsUpdateOne {
	return (&CredsClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Creds entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Creds) Unwrap() *Creds {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Creds is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Creds) String() string {
	var builder strings.Builder
	builder.WriteString("Creds(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(c.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(c.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(c.Name)
	builder.WriteString(", username=")
	builder.WriteString(c.Username)
	builder.WriteString(", url=")
	builder.WriteString(c.URL)
	builder.WriteByte(')')
	return builder.String()
}

// CredsSlice is a parsable slice of Creds.
type CredsSlice []*Creds

func (c CredsSlice) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
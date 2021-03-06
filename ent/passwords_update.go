// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mrtyunjaygr8/passwd/ent/creds"
	"github.com/mrtyunjaygr8/passwd/ent/passwords"
	"github.com/mrtyunjaygr8/passwd/ent/predicate"
)

// PasswordsUpdate is the builder for updating Passwords entities.
type PasswordsUpdate struct {
	config
	hooks    []Hook
	mutation *PasswordsMutation
}

// Where appends a list predicates to the PasswordsUpdate builder.
func (pu *PasswordsUpdate) Where(ps ...predicate.Passwords) *PasswordsUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetPassword sets the "password" field.
func (pu *PasswordsUpdate) SetPassword(s string) *PasswordsUpdate {
	pu.mutation.SetPassword(s)
	return pu
}

// SetCredID sets the "cred" edge to the Creds entity by ID.
func (pu *PasswordsUpdate) SetCredID(id int) *PasswordsUpdate {
	pu.mutation.SetCredID(id)
	return pu
}

// SetNillableCredID sets the "cred" edge to the Creds entity by ID if the given value is not nil.
func (pu *PasswordsUpdate) SetNillableCredID(id *int) *PasswordsUpdate {
	if id != nil {
		pu = pu.SetCredID(*id)
	}
	return pu
}

// SetCred sets the "cred" edge to the Creds entity.
func (pu *PasswordsUpdate) SetCred(c *Creds) *PasswordsUpdate {
	return pu.SetCredID(c.ID)
}

// Mutation returns the PasswordsMutation object of the builder.
func (pu *PasswordsUpdate) Mutation() *PasswordsMutation {
	return pu.mutation
}

// ClearCred clears the "cred" edge to the Creds entity.
func (pu *PasswordsUpdate) ClearCred() *PasswordsUpdate {
	pu.mutation.ClearCred()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PasswordsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	pu.defaults()
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PasswordsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PasswordsUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PasswordsUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PasswordsUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PasswordsUpdate) defaults() {
	if _, ok := pu.mutation.UpdateTime(); !ok {
		v := passwords.UpdateDefaultUpdateTime()
		pu.mutation.SetUpdateTime(v)
	}
}

func (pu *PasswordsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   passwords.Table,
			Columns: passwords.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: passwords.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: passwords.FieldUpdateTime,
		})
	}
	if value, ok := pu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: passwords.FieldPassword,
		})
	}
	if pu.mutation.CredCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   passwords.CredTable,
			Columns: []string{passwords.CredColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: creds.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.CredIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   passwords.CredTable,
			Columns: []string{passwords.CredColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: creds.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{passwords.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// PasswordsUpdateOne is the builder for updating a single Passwords entity.
type PasswordsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PasswordsMutation
}

// SetPassword sets the "password" field.
func (puo *PasswordsUpdateOne) SetPassword(s string) *PasswordsUpdateOne {
	puo.mutation.SetPassword(s)
	return puo
}

// SetCredID sets the "cred" edge to the Creds entity by ID.
func (puo *PasswordsUpdateOne) SetCredID(id int) *PasswordsUpdateOne {
	puo.mutation.SetCredID(id)
	return puo
}

// SetNillableCredID sets the "cred" edge to the Creds entity by ID if the given value is not nil.
func (puo *PasswordsUpdateOne) SetNillableCredID(id *int) *PasswordsUpdateOne {
	if id != nil {
		puo = puo.SetCredID(*id)
	}
	return puo
}

// SetCred sets the "cred" edge to the Creds entity.
func (puo *PasswordsUpdateOne) SetCred(c *Creds) *PasswordsUpdateOne {
	return puo.SetCredID(c.ID)
}

// Mutation returns the PasswordsMutation object of the builder.
func (puo *PasswordsUpdateOne) Mutation() *PasswordsMutation {
	return puo.mutation
}

// ClearCred clears the "cred" edge to the Creds entity.
func (puo *PasswordsUpdateOne) ClearCred() *PasswordsUpdateOne {
	puo.mutation.ClearCred()
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PasswordsUpdateOne) Select(field string, fields ...string) *PasswordsUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Passwords entity.
func (puo *PasswordsUpdateOne) Save(ctx context.Context) (*Passwords, error) {
	var (
		err  error
		node *Passwords
	)
	puo.defaults()
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PasswordsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PasswordsUpdateOne) SaveX(ctx context.Context) *Passwords {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PasswordsUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PasswordsUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PasswordsUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdateTime(); !ok {
		v := passwords.UpdateDefaultUpdateTime()
		puo.mutation.SetUpdateTime(v)
	}
}

func (puo *PasswordsUpdateOne) sqlSave(ctx context.Context) (_node *Passwords, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   passwords.Table,
			Columns: passwords.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: passwords.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Passwords.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, passwords.FieldID)
		for _, f := range fields {
			if !passwords.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != passwords.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: passwords.FieldUpdateTime,
		})
	}
	if value, ok := puo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: passwords.FieldPassword,
		})
	}
	if puo.mutation.CredCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   passwords.CredTable,
			Columns: []string{passwords.CredColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: creds.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.CredIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   passwords.CredTable,
			Columns: []string{passwords.CredColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: creds.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Passwords{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{passwords.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

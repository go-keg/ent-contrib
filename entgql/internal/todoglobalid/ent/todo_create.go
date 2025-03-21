// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/contrib/entgql/internal/todoglobalid/ent/category"
	"entgo.io/contrib/entgql/internal/todoglobalid/ent/schema/customstruct"
	"entgo.io/contrib/entgql/internal/todoglobalid/ent/todo"
	"entgo.io/contrib/entgql/internal/todoglobalid/ent/verysecret"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TodoCreate is the builder for creating a Todo entity.
type TodoCreate struct {
	config
	mutation *TodoMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (tc *TodoCreate) SetCreatedAt(t time.Time) *TodoCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TodoCreate) SetNillableCreatedAt(t *time.Time) *TodoCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetStatus sets the "status" field.
func (tc *TodoCreate) SetStatus(t todo.Status) *TodoCreate {
	tc.mutation.SetStatus(t)
	return tc
}

// SetPriority sets the "priority" field.
func (tc *TodoCreate) SetPriority(i int) *TodoCreate {
	tc.mutation.SetPriority(i)
	return tc
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (tc *TodoCreate) SetNillablePriority(i *int) *TodoCreate {
	if i != nil {
		tc.SetPriority(*i)
	}
	return tc
}

// SetText sets the "text" field.
func (tc *TodoCreate) SetText(s string) *TodoCreate {
	tc.mutation.SetText(s)
	return tc
}

// SetBlob sets the "blob" field.
func (tc *TodoCreate) SetBlob(b []byte) *TodoCreate {
	tc.mutation.SetBlob(b)
	return tc
}

// SetCategoryID sets the "category_id" field.
func (tc *TodoCreate) SetCategoryID(i int) *TodoCreate {
	tc.mutation.SetCategoryID(i)
	return tc
}

// SetNillableCategoryID sets the "category_id" field if the given value is not nil.
func (tc *TodoCreate) SetNillableCategoryID(i *int) *TodoCreate {
	if i != nil {
		tc.SetCategoryID(*i)
	}
	return tc
}

// SetInit sets the "init" field.
func (tc *TodoCreate) SetInit(m map[string]interface{}) *TodoCreate {
	tc.mutation.SetInit(m)
	return tc
}

// SetCustom sets the "custom" field.
func (tc *TodoCreate) SetCustom(c []customstruct.Custom) *TodoCreate {
	tc.mutation.SetCustom(c)
	return tc
}

// SetCustomp sets the "customp" field.
func (tc *TodoCreate) SetCustomp(c []*customstruct.Custom) *TodoCreate {
	tc.mutation.SetCustomp(c)
	return tc
}

// SetValue sets the "value" field.
func (tc *TodoCreate) SetValue(i int) *TodoCreate {
	tc.mutation.SetValue(i)
	return tc
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (tc *TodoCreate) SetNillableValue(i *int) *TodoCreate {
	if i != nil {
		tc.SetValue(*i)
	}
	return tc
}

// SetParentID sets the "parent" edge to the Todo entity by ID.
func (tc *TodoCreate) SetParentID(id int) *TodoCreate {
	tc.mutation.SetParentID(id)
	return tc
}

// SetNillableParentID sets the "parent" edge to the Todo entity by ID if the given value is not nil.
func (tc *TodoCreate) SetNillableParentID(id *int) *TodoCreate {
	if id != nil {
		tc = tc.SetParentID(*id)
	}
	return tc
}

// SetParent sets the "parent" edge to the Todo entity.
func (tc *TodoCreate) SetParent(t *Todo) *TodoCreate {
	return tc.SetParentID(t.ID)
}

// AddChildIDs adds the "children" edge to the Todo entity by IDs.
func (tc *TodoCreate) AddChildIDs(ids ...int) *TodoCreate {
	tc.mutation.AddChildIDs(ids...)
	return tc
}

// AddChildren adds the "children" edges to the Todo entity.
func (tc *TodoCreate) AddChildren(t ...*Todo) *TodoCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddChildIDs(ids...)
}

// SetCategory sets the "category" edge to the Category entity.
func (tc *TodoCreate) SetCategory(c *Category) *TodoCreate {
	return tc.SetCategoryID(c.ID)
}

// SetSecretID sets the "secret" edge to the VerySecret entity by ID.
func (tc *TodoCreate) SetSecretID(id int) *TodoCreate {
	tc.mutation.SetSecretID(id)
	return tc
}

// SetNillableSecretID sets the "secret" edge to the VerySecret entity by ID if the given value is not nil.
func (tc *TodoCreate) SetNillableSecretID(id *int) *TodoCreate {
	if id != nil {
		tc = tc.SetSecretID(*id)
	}
	return tc
}

// SetSecret sets the "secret" edge to the VerySecret entity.
func (tc *TodoCreate) SetSecret(v *VerySecret) *TodoCreate {
	return tc.SetSecretID(v.ID)
}

// Mutation returns the TodoMutation object of the builder.
func (tc *TodoCreate) Mutation() *TodoMutation {
	return tc.mutation
}

// Save creates the Todo in the database.
func (tc *TodoCreate) Save(ctx context.Context) (*Todo, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TodoCreate) SaveX(ctx context.Context) *Todo {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TodoCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TodoCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TodoCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := todo.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.Priority(); !ok {
		v := todo.DefaultPriority
		tc.mutation.SetPriority(v)
	}
	if _, ok := tc.mutation.Value(); !ok {
		v := todo.DefaultValue
		tc.mutation.SetValue(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TodoCreate) check() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Todo.created_at"`)}
	}
	if _, ok := tc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Todo.status"`)}
	}
	if v, ok := tc.mutation.Status(); ok {
		if err := todo.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Todo.status": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Priority(); !ok {
		return &ValidationError{Name: "priority", err: errors.New(`ent: missing required field "Todo.priority"`)}
	}
	if _, ok := tc.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New(`ent: missing required field "Todo.text"`)}
	}
	if v, ok := tc.mutation.Text(); ok {
		if err := todo.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Todo.text": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "Todo.value"`)}
	}
	return nil
}

func (tc *TodoCreate) sqlSave(ctx context.Context) (*Todo, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TodoCreate) createSpec() (*Todo, *sqlgraph.CreateSpec) {
	var (
		_node = &Todo{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(todo.Table, sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt))
	)
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(todo.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.Status(); ok {
		_spec.SetField(todo.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := tc.mutation.Priority(); ok {
		_spec.SetField(todo.FieldPriority, field.TypeInt, value)
		_node.Priority = value
	}
	if value, ok := tc.mutation.Text(); ok {
		_spec.SetField(todo.FieldText, field.TypeString, value)
		_node.Text = value
	}
	if value, ok := tc.mutation.Blob(); ok {
		_spec.SetField(todo.FieldBlob, field.TypeBytes, value)
		_node.Blob = value
	}
	if value, ok := tc.mutation.Init(); ok {
		_spec.SetField(todo.FieldInit, field.TypeJSON, value)
		_node.Init = value
	}
	if value, ok := tc.mutation.Custom(); ok {
		_spec.SetField(todo.FieldCustom, field.TypeJSON, value)
		_node.Custom = value
	}
	if value, ok := tc.mutation.Customp(); ok {
		_spec.SetField(todo.FieldCustomp, field.TypeJSON, value)
		_node.Customp = value
	}
	if value, ok := tc.mutation.Value(); ok {
		_spec.SetField(todo.FieldValue, field.TypeInt, value)
		_node.Value = value
	}
	if nodes := tc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.ParentTable,
			Columns: []string{todo.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.todo_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   todo.ChildrenTable,
			Columns: []string{todo.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.CategoryTable,
			Columns: []string{todo.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CategoryID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.SecretIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   todo.SecretTable,
			Columns: []string{todo.SecretColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(verysecret.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.todo_secret = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TodoCreateBulk is the builder for creating many Todo entities in bulk.
type TodoCreateBulk struct {
	config
	err      error
	builders []*TodoCreate
}

// Save creates the Todo entities in the database.
func (tcb *TodoCreateBulk) Save(ctx context.Context) ([]*Todo, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Todo, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TodoMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TodoCreateBulk) SaveX(ctx context.Context) []*Todo {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TodoCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TodoCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

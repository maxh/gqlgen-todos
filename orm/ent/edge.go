// Code generated by entc, DO NOT EDIT.

package ent

import "context"

func (o *Organization) Users(ctx context.Context) ([]*User, error) {
	result, err := o.Edges.UsersOrErr()
	if IsNotLoaded(err) {
		result, err = o.QueryUsers().All(ctx)
	}
	return result, err
}

func (t *Todo) User(ctx context.Context) (*User, error) {
	result, err := t.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryUser().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) Todos(ctx context.Context) ([]*Todo, error) {
	result, err := u.Edges.TodosOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryTodos().All(ctx)
	}
	return result, err
}

func (u *User) Organization(ctx context.Context) (*Organization, error) {
	result, err := u.Edges.OrganizationOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryOrganization().Only(ctx)
	}
	return result, MaskNotFound(err)
}

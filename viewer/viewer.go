package viewer

import (
	"context"
	"errors"
	"github.com/maxh/gqlgen-todos/orm/ent"
)

// Role for viewer actions.
type Role int

// List of roles.
const (
	_ Role = 1 << iota
	Admin
	View
)

// Viewer describes the query/mutation viewer-context.
type Viewer interface {
	IsAdmin() bool   // If viewer is admin.
	Tenant() string  // Tenant name.
	User() *ent.User // User name.
}

// UserViewer describes a user-viewer.
type UserViewer struct {
	U    *ent.User
	T    *ent.Tenant
	Role Role // Attached roles.
}

func (v UserViewer) IsAdmin() bool {
	return v.Role&Admin != 0
}

func (v UserViewer) Tenant() string {
	if v.T != nil {
		return v.T.Name
	}
	return ""
}

func (v UserViewer) User() *ent.User {
	if v.U != nil {
		return v.U
	}
	return nil
}

type ctxKey struct{}

// UserFromContext returns the Viewer stored in a context.
func UserFromContext(ctx context.Context) (*ent.User, error) {
	v, _ := ctx.Value(ctxKey{}).(Viewer)
	if v.User() == nil {
		return nil, errors.New("no user in context")
	}
	return v.User(), nil
}

// FromContext returns the Viewer stored in a context.
func FromContext(ctx context.Context) Viewer {
	v, _ := ctx.Value(ctxKey{}).(Viewer)
	return v
}

// NewContext returns a copy of parent context with the given Viewer attached with it.
func NewContext(parent context.Context, v Viewer) context.Context {
	return context.WithValue(parent, ctxKey{}, v)
}

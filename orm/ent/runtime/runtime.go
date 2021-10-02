// Code generated by entc, DO NOT EDIT.

package runtime

import (
	"context"

	"github.com/maxh/gqlgen-todos/orm/ent/organization"
	"github.com/maxh/gqlgen-todos/orm/ent/tenant"
	"github.com/maxh/gqlgen-todos/orm/ent/todo"
	"github.com/maxh/gqlgen-todos/orm/ent/user"
	"github.com/maxh/gqlgen-todos/orm/schema"
	"github.com/maxh/gqlgen-todos/qid"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	organizationMixin := schema.Organization{}.Mixin()
	organizationMixinFields0 := organizationMixin[0].Fields()
	_ = organizationMixinFields0
	organizationFields := schema.Organization{}.Fields()
	_ = organizationFields
	// organizationDescName is the schema descriptor for name field.
	organizationDescName := organizationFields[0].Descriptor()
	// organization.DefaultName holds the default value on creation for the name field.
	organization.DefaultName = organizationDescName.Default.(string)
	// organizationDescID is the schema descriptor for id field.
	organizationDescID := organizationMixinFields0[0].Descriptor()
	// organization.DefaultID holds the default value on creation for the id field.
	organization.DefaultID = organizationDescID.Default.(func() qid.ID)
	tenantMixin := schema.Tenant{}.Mixin()
	tenant.Policy = privacy.NewPolicies(tenantMixin[1], schema.Tenant{})
	tenant.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := tenant.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	tenantMixinFields0 := tenantMixin[0].Fields()
	_ = tenantMixinFields0
	tenantFields := schema.Tenant{}.Fields()
	_ = tenantFields
	// tenantDescName is the schema descriptor for name field.
	tenantDescName := tenantFields[0].Descriptor()
	// tenant.NameValidator is a validator for the "name" field. It is called by the builders before save.
	tenant.NameValidator = tenantDescName.Validators[0].(func(string) error)
	// tenantDescID is the schema descriptor for id field.
	tenantDescID := tenantMixinFields0[0].Descriptor()
	// tenant.DefaultID holds the default value on creation for the id field.
	tenant.DefaultID = tenantDescID.Default.(func() qid.ID)
	todoMixin := schema.Todo{}.Mixin()
	todoMixinFields0 := todoMixin[0].Fields()
	_ = todoMixinFields0
	todoFields := schema.Todo{}.Fields()
	_ = todoFields
	// todoDescText is the schema descriptor for text field.
	todoDescText := todoFields[0].Descriptor()
	// todo.DefaultText holds the default value on creation for the text field.
	todo.DefaultText = todoDescText.Default.(string)
	// todoDescDone is the schema descriptor for done field.
	todoDescDone := todoFields[1].Descriptor()
	// todo.DefaultDone holds the default value on creation for the done field.
	todo.DefaultDone = todoDescDone.Default.(bool)
	// todoDescID is the schema descriptor for id field.
	todoDescID := todoMixinFields0[0].Descriptor()
	// todo.DefaultID holds the default value on creation for the id field.
	todo.DefaultID = todoDescID.Default.(func() qid.ID)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
	// userDescID is the schema descriptor for id field.
	userDescID := userMixinFields0[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() qid.ID)
}

const (
	Version = "(devel)" // Version of ent codegen.
)

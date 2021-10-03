package auth

import (
	"context"
	"fmt"
	"github.com/maxh/gqlgen-todos/orm/ent"
	"github.com/maxh/gqlgen-todos/orm/ent/privacy"
	"github.com/maxh/gqlgen-todos/viewer"
	"log"
	"net/http"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// A stand-in for our database backed user object
type User struct {
	Name    string
	IsAdmin bool
}

// Middleware decodes the share session cookie and packs the session into context
func Middleware(client *ent.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("auth-cookie")

			//// Allow unauthenticated users in
			//if err != nil || c == nil {
			//	next.ServeHTTP(w, r)
			//	return
			//}

			_, err = validateAndGetUserID(c)
			if err != nil {
				http.Error(w, "Invalid cookie", http.StatusForbidden)
				return
			}

			ctx := r.Context()
			allow := privacy.DecisionContext(ctx, privacy.Allow)
			user, err := client.User.Query().First(allow)
			if user == nil || err != nil {
				log.Fatal("unable to get user", err)
			}

			// put it in context
			// ctx := context.WithValue(r.Context(), userCtxKey, user)
			t, err := client.Tenant.Query().First(allow)
			if err != nil {
				log.Fatal("unable to get Tenant", err)
			}

			ctx = viewer.NewContext(ctx, viewer.UserViewer{
				U:    user,
				T:    t,
				Role: viewer.Admin,
			})
			fmt.Println("got t")
			fmt.Println(t.String())
			// and call the next with our new context
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

func validateAndGetUserID(c *http.Cookie) (string, error) {
	return "", nil
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *User {
	raw, _ := ctx.Value(userCtxKey).(*User)
	return raw
}

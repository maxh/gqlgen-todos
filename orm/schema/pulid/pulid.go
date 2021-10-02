package pulid

import (
	"crypto/rand"
	"database/sql/driver"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
)

// ID implements a PULID - a prefixed ULID.
type ID string

// The default entropy source.
var defaultEntropySource *ulid.MonotonicEntropy

var qidPrefix = "qrn"
var serviceName = "qgo"

func init() {
	// Seed the default entropy source.
	// TODO: To improve testability, this package should allow control of entropy sources and the time.Now implementation.
	defaultEntropySource = ulid.Monotonic(rand.Reader, 0)
}

// newULID returns a new ULID for time.Now() using the default entropy source.
func newULID() ulid.ULID {
	return ulid.MustNew(ulid.Timestamp(time.Now()), defaultEntropySource)
}

// newULIDLower returns a lower-cased ULID.
func newULIDLower() string {
	return strings.ToLower(newULID().String())
}

// MustNew returns a new PULID for time.Now() given a resourceType. This uses the default entropy source.
func MustNew(resourceType string) ID {
	return ID(fmt.Sprintf("%s::%s:%s:%s", qidPrefix, serviceName, resourceType, newULIDLower()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (u *ID) UnmarshalGQL(v interface{}) error {
	return u.Scan(v)
}

// MarshalGQL implements the graphql.Marshaler interface
func (u ID) MarshalGQL(w io.Writer) {
	_, _ = io.WriteString(w, strconv.Quote(string(u)))
}

// Scan implements the Scanner interface.
func (u *ID) Scan(src interface{}) error {
	if src == nil {
		return fmt.Errorf("pulid: expected a value")
	}
	s, ok := src.(string)
	if !ok {
		return fmt.Errorf("pulid: expected a string")
	}
	*u = ID(s)
	return nil
}

// Value implements the driver Valuer interface.
func (u ID) Value() (driver.Value, error) {
	return string(u), nil
}

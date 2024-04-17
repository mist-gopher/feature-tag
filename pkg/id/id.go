// A value object abstraction of unique identifier
package id

import "github.com/google/uuid"

type Id struct {
	value string
}

func (i Id) Value() string {
	return i.value
}

func (i Id) IsEqual(other Id) bool {
	return i.value == other.value
}

func New() Id {
	return Id{uuid.NewString()}
}

func FromString(v string) (Id, error) {
	return Id{v}, uuid.Validate(v)
}

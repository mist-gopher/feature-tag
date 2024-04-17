package id_test

import (
	"testing"

	"github.com/mist-gopher/feature-tag/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestFromString(t *testing.T) {
	t.Run("should return error when receive a invalid uuid", func(t *testing.T) {
		input := "invalid"
		v, err := id.FromString(input)
		assert.NotNil(t, err)
		assert.Equal(t, input, v.Value())
	})

	t.Run("should return nil error when receive a valid uuid", func(t *testing.T) {
		input := "b595057d-b343-48dd-9350-ffe6e4ad68d8"
		v, err := id.FromString(input)
		assert.Nil(t, err)
		assert.Equal(t, input, v.Value())
	})
}

func TestIsEqual(t *testing.T) {
	id1 := id.New()
	id2, _ := id.FromString(id1.Value())
	id3 := id.New()

	assert.True(t, id1.IsEqual(id2))
	assert.True(t, id2.IsEqual(id1))
	assert.False(t, id3.IsEqual(id1))
	assert.False(t, id2.IsEqual(id3))
}

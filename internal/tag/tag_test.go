package tag_test

import (
	"testing"

	"github.com/mist-gopher/feature-tag/internal/tag"
	"github.com/stretchr/testify/assert"
)

func TestNewTag(t *testing.T){
  ntag := tag.New("clientid", "tagname", false)
  assert.Equal(t, "tag:clientid:tagname", ntag.Id)
  assert.Equal(t, "tagname", ntag.Name)
  assert.False(t, ntag.Value)
}


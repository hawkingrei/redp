package version

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestVersion(t *testing.T) {
	assert.NotEqual(t, String(), "")
}

package textwrap

import (
	"gotest.tools/assert"
	"testing"
)

func TestTrimLeft(t *testing.T) {
	assert.Equal(t, NewTextWrap().TrimLeft("    trim me"), "trim me")
}

func TestTrimLeftWithTrail(t *testing.T) {
	assert.Equal(t, NewTextWrap().TrimLeft("    trim me  "), "trim me  ")
}

func TestTrimLeftLeadingTabs(t *testing.T) {
	assert.Equal(t, NewTextWrap().TrimLeft("   \t\t\ttrim me"), "trim me")
}

func TestTrimeLeftTrailingTabs(t *testing.T) {
	assert.Equal(t, NewTextWrap().TrimLeft("   \t\t\ttrim me\t  \t"), "trim me\t  \t")
}

// ----
func TestTrimRight(t *testing.T) {
	assert.Equal(t, NewTextWrap().TrimRight("trim me    "), "trim me")
}

func TestTrimRightWithTrail(t *testing.T) {
	assert.Equal(t, NewTextWrap().TrimRight("   trim me  "), "   trim me")
}

func TestTrimRightLeadingTabs(t *testing.T) {
	assert.Equal(t, NewTextWrap().TrimRight("trim me\t  \t  \t   "), "trim me")
}

func TestTrimeRightTrailingTabs(t *testing.T) {
	assert.Equal(t, NewTextWrap().TrimRight("   \t\t\ttrim me\t  \t"), "   \t\t\ttrim me")
}

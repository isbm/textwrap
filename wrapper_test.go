package textwrap

import (
	"gotest.tools/assert"
	"testing"
)

/*
  Test TrimLeft on trimming leading whitespace from the string
*/
func TestTrimLeft(t *testing.T) {
	assert.Equal(t, NewTextWrap().TrimLeft("    trim me"), "trim me")
}

/*
  Test TrimLeft on trimming leading whitespace from the string, containing trailing spaces
*/
func TestTrimLeftWithTrail(t *testing.T) {
	assert.Equal(t, NewTextWrap().TrimLeft("    trim me  "), "trim me  ")
}

/*
  Test TrimLeft on trimming mixed leading whitespace from the string
*/
func TestTrimLeftLeadingTabs(t *testing.T) {
	assert.Equal(t, NewTextWrap().TrimLeft("   \t\t\ttrim me"), "trim me")
}

/*
  Test TrimLeft on trimming leading newline.
*/
func TestTrimLeftNewline(t *testing.T) {
	assert.Equal(t, NewTextWrap().TrimLeft("\nline\n"), "line\n")
}

/*
  Test TrimLeft on trimming leading whitespace from the string,
  containing trailing mixed whitespace.
*/
func TestTrimeLeftTrailingTabs(t *testing.T) {
	assert.Equal(t, NewTextWrap().TrimLeft("   \t\t\ttrim me\t  \t"), "trim me\t  \t")
}

/*
  Test TrimLeft on trimming trailing whitespace from the string.
*/
func TestTrimRight(t *testing.T) {
	assert.Equal(t, NewTextWrap().TrimRight("trim me    "), "trim me")
}

/*
  Test TrimLeft on trimming trailing whitespace from the string,
  containing leading whitespace.
*/
func TestTrimRightWithLead(t *testing.T) {
	assert.Equal(t, NewTextWrap().TrimRight("   trim me  "), "   trim me")
}

/*
  Test TrimLeft on trimming trailing mixed whitespace from the string.
*/
func TestTrimRightTrailingTabs(t *testing.T) {
	assert.Equal(t, NewTextWrap().TrimRight("trim me\t  \t  \n\t   "), "trim me")
}

/*
  Test TrimLeft on trimming trailing mixed whitespace from the string,
  containing leading mixed whitespace
*/
func TestTrimRightLeadingTabs(t *testing.T) {
	assert.Equal(t, NewTextWrap().TrimRight("   \t\t\ttrim me\t  \t\n"), "   \t\t\ttrim me")
}

/*
  Test TrimLeft on trimming trailing newline.
*/
func TestTrimRightNewline(t *testing.T) {
	assert.Equal(t, NewTextWrap().TrimRight("\nline\n"), "\nline")
}

package textwrap

import (
	"gotest.tools/assert"
	"testing"
)

/*
  Dedent one line.
*/
func TestDedentOneLine(t *testing.T) {
	text := "  one line"
	assert.Equal(t, NewTextWrap().Dedent(text), "one line")
}

/*
  Test text dedent with multiple lines,
  containing no empty lines in it.
*/
func TestDedentMultilineNoEmpty(t *testing.T) {
	multiline := "  two spaces\n    four spaces\n      six spaces\n   three spaces"
	expected := "two spaces\n  four spaces\n    six spaces\n three spaces"
	assert.Equal(t, NewTextWrap().Dedent(multiline), expected)
}

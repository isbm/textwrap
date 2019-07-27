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

/*
  Test text dedent with multiple lines,
  containing tabs.
*/
func TestDedentMultilineMixedSpace(t *testing.T) {
	multiline := "\ttab\n onespace"
	expected := "   tab\nonespace"
	assert.Equal(t, NewTextWrap().Dedent(multiline), expected)
}

/*
  Test text dedent with multiple lines,
  containing tabs, different tab space
*/
func TestDedentMultilineTabSpace(t *testing.T) {
	multiline := "\ttab\n onespace"
	expected := "       tab\nonespace"
	assert.Equal(t, NewTextWrap().SetTabSpacesWidth(8).Dedent(multiline), expected)
}

/*
  Test text dedent without dropping tabs, where whitespace contains only tabs.
*/
func TestDedentMultilineNoTabsDropNoSpace(t *testing.T) {
	multiline := "\ttab\n onespace"
	expected := "\ttab\n onespace"
	assert.Equal(t, NewTextWrap().SetExpandTabs(false).Dedent(multiline), expected)
}

/*
  Test text dedent without dropping tabs, where whitespace is mixed.
*/
func TestDedentMultilineNoTabsDrop(t *testing.T) {
	multiline := " \ttab\n  onespace"
	expected := "\ttab\n onespace"
	assert.Equal(t, NewTextWrap().SetExpandTabs(false).Dedent(multiline), expected)
}

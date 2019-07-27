package textwrap

import (
	"gotest.tools/assert"
	"testing"
)

const LOREM_IPSUM = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, " +
	"sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim " +
	"ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip " +
	"ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate " +
	"velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat " +
	"cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

/*
   Test filler (joins wrapped array) by default settings.
*/
func TestFillWrappingDefault(t *testing.T) {
	expected := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do\n" +
		"eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim\n" +
		"ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut\n" +
		"aliquip ex ea commodo consequat. Duis aute irure dolor in\n" +
		"reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla\n" +
		"pariatur. Excepteur sint occaecat cupidatat non proident, sunt in\n" +
		"culpa qui officia deserunt mollit anim id est laborum."
	assert.Equal(t, NewTextWrap().Fill(LOREM_IPSUM), expected)
}

/*
  Test filler (joins wrapped array) by 50 character width.
*/
func TestFillWrapping50(t *testing.T) {
	expected := "Lorem ipsum dolor sit amet, consectetur\n" +
		"adipiscing elit, sed do eiusmod tempor incididunt\n" +
		"ut labore et dolore magna aliqua. Ut enim ad\n" +
		"minim veniam, quis nostrud exercitation ullamco\n" +
		"laboris nisi ut aliquip ex ea commodo consequat.\n" +
		"Duis aute irure dolor in reprehenderit in\n" +
		"voluptate velit esse cillum dolore eu fugiat\n" +
		"nulla pariatur. Excepteur sint occaecat cupidatat\n" +
		"non proident, sunt in culpa qui officia deserunt\n" +
		"mollit anim id est laborum."
	assert.Equal(t, NewTextWrap().SetWidth(50).Fill(LOREM_IPSUM), expected)
}

/*
  Test raw wrapper by default settings.
*/
func TestWrapDefault(t *testing.T) {
	expected := []string{
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do",
		"eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim",
		"ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut",
		"aliquip ex ea commodo consequat. Duis aute irure dolor in",
		"reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla",
		"pariatur. Excepteur sint occaecat cupidatat non proident, sunt in",
		"culpa qui officia deserunt mollit anim id est laborum.",
	}
	result := NewTextWrap().Wrap(LOREM_IPSUM)
	assert.Equal(t, len(result), len(expected))

	for idx, line := range result {
		assert.Equal(t, line, expected[idx])
		assert.Assert(t, len(line) <= 70)
	}
}

/*
  Test raw wrapper by width of 150 characters.
*/
func TestWrapDefault150(t *testing.T) {
	expected := []string{
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do " +
			"eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,",
		"quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea " +
			"commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse",
		"cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat " +
			"cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est",
		"laborum.",
	}
	result := NewTextWrap().SetWidth(150).Wrap(LOREM_IPSUM)
	assert.Equal(t, len(result), len(expected))

	for idx, line := range result {
		assert.Equal(t, line, expected[idx])
		assert.Assert(t, len(line) <= 150)
	}
}

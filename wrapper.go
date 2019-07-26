package textwrap

const (
	WHITESPACE = " \t\n\r\x0b\x0c"
)

type textWrap struct {
	width             int
	replaceWhitespace bool
	dropwWhitespace   bool
	initialIndent     string
	tabSpacesWidth    int
}

// Constructor
func NewTextWrap() *textWrap {
	wrap := new(textWrap)
	wrap.width = 70
	wrap.replaceWhitespace = true
	wrap.initialIndent = ""
	wrap.dropwWhitespace = true
	wrap.tabSpacesWidth = 4

	return wrap
}

/*
  Sets replace whitespace property
*/
func (wrap *textWrap) SetTabSpacesWidth(width int) *textWrap {
	wrap.tabSpacesWidth = width
	return wrap
}

/*
  Sets drop whitespace property
*/
func (wrap *textWrap) SetDropWhitespace(drop bool) *textWrap {
	wrap.dropwWhitespace = drop
	return wrap
}

/*
  Sets initial indent property
*/
func (wrap *textWrap) SetInitialIndent(indent string) *textWrap {
	wrap.initialIndent = indent
	return wrap
}

/*
  Sets replace whitespace property
*/
func (wrap *textWrap) SetReplaceWhitespace(replace bool) *textWrap {
	wrap.replaceWhitespace = replace
	return wrap
}

/*
  Wraps the single paragraph in text (a string)
  so every line is at most width characters long.
  Returns a list of output lines, without final newlines.
*/
func (wrap *textWrap) Wrap(text string) []string {
	return nil
}

/*
  Wraps the single paragraph in text, and returns
  a single string containing the wrapped paragraph.
*/
func (wrap *textWrap) Fill(text string) string {
	return ""
}

/*
  Remove any common leading whitespace from every line in text.
*/
func (wrap *textWrap) Dedent(text string) string {
	return ""
}

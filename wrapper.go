// Package textwrap is a semi-port of Python's equivalent module "textwrap".
// Most of the functionality is similar (or in a progress to get there).
// The textwrap module provides two convenience methods: Wrap() and Fill().
// All of them are methods of textWrap class that does all the work.
// Package also provides an utility function Dedent() as well as
// TrimLeft() and TrimRight() to strip a whitespace respectively.
package textwrap

import (
	"regexp"
	"sort"
	"strings"
)

const (
	// WHITESPACE is a string of symbols that are subject to be stripped
	// in the incoming data
	WHITESPACE = " \t\n\r\x0b\x0c"
)

// The textWrap is an object for keeping configuration of an instance
// that performs the wrapping.
type textWrap struct {
	width             int
	replaceWhitespace bool
	dropwWhitespace   bool
	expandTabs        bool
	initialIndent     string
	tabSpacesWidth    int
	newline           string
}

// NewTextWrap function returns text wrapper instance object.
// Constructor does not accept any params, however can be configured
// in chain methods setting configuration settings.
func NewTextWrap() *textWrap {
	wrap := new(textWrap)
	wrap.width = 70
	wrap.replaceWhitespace = true
	wrap.dropwWhitespace = true
	wrap.expandTabs = true
	wrap.initialIndent = ""
	wrap.tabSpacesWidth = 4
	wrap.newline = "\n"

	return wrap
}

// SetExpandTabls sets tab expansion to spaces. The width of the spaces is set
// by SetTabSpacesWidth method.
func (wrap *textWrap) SetExpandTabs(expand bool) *textWrap {
	wrap.expandTabs = expand
	return wrap
}

// SetNewline sets newline character. Default is "\n".
func (wrap *textWrap) SetNewline(newline string) *textWrap {
	wrap.newline = newline
	return wrap
}

// SetWidth sets width of the wrapped text so the content does not exceedes it.
func (wrap *textWrap) SetWidth(width int) *textWrap {
	wrap.width = width
	return wrap
}

// SetTabSpacesWidth sets replace whitespace property
func (wrap *textWrap) SetTabSpacesWidth(width int) *textWrap {
	wrap.tabSpacesWidth = width
	return wrap
}

// SetDropWhitespace sets drop whitespace property
func (wrap *textWrap) SetDropWhitespace(drop bool) *textWrap {
	wrap.dropwWhitespace = drop
	return wrap
}

// SetInitialIndent sets initial indent property
func (wrap *textWrap) SetInitialIndent(indent string) *textWrap {
	wrap.initialIndent = indent
	return wrap
}

// SetReplaceWhitespace sets replace whitespace property
func (wrap *textWrap) SetReplaceWhitespace(replace bool) *textWrap {
	wrap.replaceWhitespace = replace
	return wrap
}

// Wrap method sraps the single paragraph in text (a string)
// so every line is at most width characters long.
// Returns a list of output lines, without final newlines.
func (wrap *textWrap) Wrap(text string) []string {
	buff := make([]string, 0)
	line := ""
	for _, word := range regexp.MustCompile(" ").Split(text, -1) {
		if len(line+word) < wrap.width {
			line += word + " "
		} else {
			line = strings.TrimSpace(line)
			if line != "" {
				buff = append(buff, strings.TrimSpace(line))
			}
			line = word + " "
		}
	}
	line = strings.TrimSpace(line)
	if line != "" {
		buff = append(buff, strings.TrimSpace(line))
	}
	return buff
}

// Fill method wraps the single paragraph in text, and returns
// a single string containing the wrapped paragraph.
func (wrap *textWrap) Fill(text string) string {
	return strings.Join(wrap.Wrap(text), wrap.newline)
}

// Internal method. Gets configured whitespace.
func (wrap *textWrap) getCurrentWhitespace() string {
	var ws string
	if !wrap.expandTabs {
		ws = strings.Replace(WHITESPACE, "\t", "", -1)
	} else {
		ws = WHITESPACE
	}
	return ws
}

// TrimLeft trimming leading whitespace from the text line
func (wrap *textWrap) TrimLeft(line string) string {
	var buff strings.Builder
	ws := false
	currentWhitespace := wrap.getCurrentWhitespace()
	for idx, char := range line {
		if strings.Contains(currentWhitespace, string(char)) {
			if idx == 0 {
				ws = true
			}

			if ws {
				continue
			}
		}
		buff.WriteRune(char)
		ws = false
	}

	return buff.String()
}

// Internal method. Reverses a string
func (wrap *textWrap) reverseString(text string) string {
	rns := []rune(text)
	var buff string
	for i := len(rns) - 1; i >= 0; i-- {
		buff += string(rns[i])
	}
	return buff
}

// TrimRight trimms trailing whitespace from the text line
func (wrap *textWrap) TrimRight(line string) string {
	return wrap.reverseString(wrap.TrimLeft(wrap.reverseString(line)))
}

// ExpandTabs converts tabs to an amount of spaces, set by SetTabSpacesWidth
func (wrap *textWrap) ExpandTabs(line string) string {
	if wrap.expandTabs {
		line = strings.Replace(line, "\t", strings.Repeat(" ", wrap.tabSpacesWidth), -1)
	}

	return line
}

// Method Dedent removes any common leading whitespace from every line in text.
func (wrap *textWrap) Dedent(text string) string {
	buff := make([]string, 0)
	wsbuff := make([]int, 0)
	for _, line := range strings.Split(text, wrap.newline) {
		if len(line) > 0 {
			line = wrap.ExpandTabs(line)
			wsbuff = append(wsbuff, len(line)-len(wrap.TrimLeft(line)))
		}
		buff = append(buff, line)
	}
	sort.Ints(wsbuff)
	minws := wsbuff[0]
	for idx, val := range wsbuff {
		if idx == 0 || val < minws {
			minws = val
		}
	}

	var sbuff strings.Builder
	buffLen := len(buff)
	for idx, line := range buff {
		if len(line) > 0 {
			line = line[minws:]
		}
		if idx < buffLen-1 {
			sbuff.WriteString(line + wrap.newline)
		} else {
			sbuff.WriteString(line)
		}
	}

	return sbuff.String()
}

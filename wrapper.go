package textwrap

import (
	"regexp"
	"sort"
	"strings"
)

const (
	WHITESPACE = " \t\n\r\x0b\x0c"
)

type textWrap struct {
	width             int
	replaceWhitespace bool
	dropwWhitespace   bool
	expandTabs        bool
	initialIndent     string
	tabSpacesWidth    int
	newline           string
}

// Constructor
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

/*
  Sets tab expansion to spaces. The width of the spaces is set
  by SetTabSpacesWidth method.
*/
func (wrap *textWrap) SetExpandTabs(expand bool) *textWrap {
	wrap.expandTabs = expand
	return wrap
}

/*
  Sets newline character. Default is "\n".
*/
func (wrap *textWrap) SetNewline(newline string) *textWrap {
	wrap.newline = newline
	return wrap
}

/*
  Sets width of the wrapped text so the content does not exceedes it.
*/
func (wrap *textWrap) SetWidth(width int) *textWrap {
	wrap.width = width
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
	buff := make([]string, 0)
	line := ""
	for _, word := range regexp.MustCompile(" ").Split(text, -1) {
		if len(line+word) < wrap.width {
			line += word + " "
		} else {
			buff = append(buff, strings.TrimSpace(line))
			line = word + " "
		}
	}
	buff = append(buff, strings.TrimSpace(line))

	return buff
}

/*
  Wraps the single paragraph in text, and returns
  a single string containing the wrapped paragraph.
*/
func (wrap *textWrap) Fill(text string) string {
	return strings.Join(wrap.Wrap(text), wrap.newline)
}

/*
  Get configured whitespace
*/
func (wrap *textWrap) getCurrentWhitespace() string {
	var ws string
	if !wrap.expandTabs {
		ws = strings.Replace(WHITESPACE, "\t", "", -1)
	} else {
		ws = WHITESPACE
	}
	return ws
}

// Trim leading whitespace from the text line
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

// Reverses a string
func (wrap *textWrap) reverseString(text string) string {
	rns := []rune(text)
	var buff string
	for i := len(rns) - 1; i >= 0; i-- {
		buff += string(rns[i])
	}
	return buff
}

// Trim trailing whitespace from the text line
func (wrap *textWrap) TrimRight(line string) string {
	return wrap.reverseString(wrap.TrimLeft(wrap.reverseString(line)))
}

func (wrap *textWrap) ExpandTabs(line string) string {
	if wrap.expandTabs {
		line = strings.Replace(line, "\t", strings.Repeat(" ", wrap.tabSpacesWidth), -1)
	}

	return line
}

/*
  Remove any common leading whitespace from every line in text.
*/
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

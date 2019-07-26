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
	initialIndent     string
	tabSpacesWidth    int
	newline           string
}

// Constructor
func NewTextWrap() *textWrap {
	wrap := new(textWrap)
	wrap.width = 70
	wrap.replaceWhitespace = true
	wrap.initialIndent = ""
	wrap.dropwWhitespace = true
	wrap.tabSpacesWidth = 4
	wrap.newline = "\n"

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
	var pl *string
	for _, word := range regexp.MustCompile(" ").Split(text, -1) {
		if len(line+word) < wrap.width {
			line += word + " "
			if pl == nil {
				pl = &line
			}
		} else {
			buff = append(buff, strings.TrimSpace(line))
			line = ""
			pl = nil
		}
	}
	if pl != nil {
		buff = append(buff, strings.TrimSpace(*pl))
	}

	return buff
}

/*
  Wraps the single paragraph in text, and returns
  a single string containing the wrapped paragraph.
*/
func (wrap *textWrap) Fill(text string) string {
	return strings.Join(wrap.Wrap(text), wrap.newline)
}

// Trim leading whitespace from the text line
func (wrap *textWrap) TrimLeft(line string) string {
	var buff strings.Builder
	ws := false
	for idx, char := range line {
		if strings.Contains(WHITESPACE, string(char)) {
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
	return strings.Replace(line, "\t", "    ", -1)
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
	for _, line := range buff {
		if len(line) > 0 {
			line = line[minws:]
		}
		sbuff.WriteString(line + wrap.newline)
	}

	return sbuff.String()
}

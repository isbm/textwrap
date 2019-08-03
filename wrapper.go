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
	_ansiregex = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"
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
	stripAnsiRegex    *regexp.Regexp
	ansiSavvy         bool
	ansiStrip         bool
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
	wrap.stripAnsiRegex = regexp.MustCompile(_ansiregex)
	wrap.ansiSavvy = false
	wrap.ansiStrip = false

	return wrap
}

// This method works only for plain text wrapping. In case incoming text
// contains ANSI escapes, it will be just stripped away.
func (wrap *textWrap) SetStripANSI(strip bool) *textWrap {
	wrap.ansiStrip = strip
	return wrap
}

// Set ANSI savvy (or not). Default: OFF
func (wrap *textWrap) SetANSISavvy(mode bool) *textWrap {
	wrap.ansiSavvy = mode
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

// ANSI wrap
/*
  Limitation: indexes are isolated at the moment, so each ANSI tag is treated individually.
			  That leads to only one tag carry-over, all others are lost.

			  To fix this, indexes needs to be re-grouped if they care one tag after another,
			  so then this structure is treated at once, and thus all grouped tags are carried over.
*/
func (wrap *textWrap) ansiWrap(text string) []string {
	// Get all tags
	indexes := wrap.stripAnsiRegex.FindAllStringSubmatchIndex(text, -1)
	ansitags := wrap.stripAnsiRegex.FindAllStringSubmatch(text, -1)

	buff := make([]string, 0)
	line := ""

	// Wrap plain text
	textWords := regexp.MustCompile(" ").Split(wrap.stripAnsi(text), -1)
	for idx, word := range textWords {
		if strings.TrimSpace(word) == "" { // Throw away extra spaces
			continue
		}

		if len(line+word) < wrap.width {
			line += word
			if idx < len(textWords)-1 {
				line += " "
			}
		} else {
			buff = append(buff, line)
			line = word
			if idx < len(textWords)-1 {
				line += " "
			}
		}
	}
	buff = append(buff, line)

	// Re-install ANSI tags
	carryover := make(map[int]string)
	linepath := 0
	rest := 0
	lastLine := ""
	for idx, line := range buff {
		lastLine = line
		for escIdx, escSetOff := range indexes {
			escOff := escSetOff[0]
			if escOff < linepath {
				continue
			}
			if escOff > len(line)+linepath {
				rest = escIdx
				carryover[idx+1] = ansitags[escIdx-1][0]
				break
			}

			var ansiOffset int
			if linepath == 0 { // first line
				ansiOffset = escOff
			} else {
				ansiOffset = escOff - linepath
			}
			line = line[:ansiOffset] + ansitags[escIdx][0] + line[ansiOffset:]
		}

		linepath += len(line)
		buff[idx] = line
	}

	// Fetch the rest of the tags for the last line (this is quite buggy)
	for escIdx, escSetOff := range indexes[rest:] {
		escOff := linepath - escSetOff[0] - (len(lastLine) - 1)

		if escOff < len(lastLine) {
			lastLine = lastLine[:escOff] + ansitags[escIdx+rest][0] + lastLine[escOff:]
		}
	}
	buff[len(buff)-1] = lastLine

	// Install ANSI-terminator at each end of the line
	for idx, line := range buff {
		buff[idx] = line + "\x1b[0m"
	}

	// Install carryover tags
	for line, tag := range carryover {
		buff[line] = tag + buff[line]
	}

	return buff
}

// Wrap method sraps the single paragraph in text (a string)
// so every line is at most width characters long.
// Returns a list of output lines, without final newlines.
func (wrap *textWrap) Wrap(text string) []string {
	var out []string
	if wrap.ansiSavvy {
		out = wrap.ansiWrap(text)
	} else {
		out = wrap.plainWrap(text)
	}
	return out
}

// Plain wrap
func (wrap *textWrap) plainWrap(text string) []string {
	if wrap.ansiStrip {
		text = wrap.stripAnsi(text)
	}
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

// Allow support ANSI-colored data. If the data is not stripped out,
// all the widths will be wrongly calculated
func (wrap *textWrap) stripAnsi(data string) string {
	return wrap.stripAnsiRegex.ReplaceAllString(data, "")
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

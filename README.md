# Text Wrap

This is a port of Python's "textwrap" module for Go. Well, sort of...

# Limitations

This modules (at least for now) is not wrapping on whitespaces and
right after hyphens in compound words, as it is customary in
English. That said, `break_on_hyphens` and `break_long_words` are not
yet supported.

Also `fix_sentence_endings` is not supported as well for now, which
doesn't work reliably in Python anyways (since it requires two spaces
and other conditions nobody cares of).

The implementation for hyphens support is planned however, while
`fix_sentence_endings` is not (but! your PRs are welcome and free to
implement it).

# Usage

The usage is quite similar as in Python:

```go
import (
	"fmt"
	"github.com/isbm/textwrap"
)

...

text := "Your very long text here"
wrapper := textwrap.NewTextWrap() // Defaults to 70
fmt.Println(wrapper.Fill(text))      // Returns string

// Get each line
for idx, line := range wrapper.Wrap(text) {
	fmt.Println(idx, line)
}

```

De-dent is also implemented and works exactly the same as in Python:

```go
multilineText := `
    There is some multiline text
  with different identation
      everywhere. So it will be
    aligned to the minimal.
`

// This will remove two leading spaces from each line
fmt.Println(wrapper.Dedent(multilineText))

```


# Configuration

You can setup wrapper object constructor the following way (given values are its
defaults, so you can change it to whatever you want):

```go
wrapper := textwrap.NewTextWrap().
	SetNewLine("\n").
	SetWidth(70),
	SetTabSpacesWidth(4).
	SetDropWhitespace(true).
	SetInitialIndent("").
	SetReplaceWhitespace(true)
```

Have fun.

# Bonus Functions

While it is possible to do it differently, this module also gives you
string whitespace trimming for _only_ leading whitespace (`TrimLeft`) or _only_
trailing (`TrimRight`), as contrary to `strings.TrimSpace` that trims
everything.

The whitespace is the same as defined in Python's `strings.whitespace`.


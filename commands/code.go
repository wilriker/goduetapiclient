package commands

import (
	"fmt"
	"strings"

	"github.com/wilriker/goduetapiclient/types"
)

type Message struct {
	Type    types.MessageType
	Content string
}

func (m Message) String() string {
	switch m.Type {
	case types.Error:
		return fmt.Sprintf("Error: %s", m.Content)
	case types.Warning:
		return fmt.Sprintf("Warning: %s", m.Content)
	default:
		return m.Content
	}
}

type CodeResult []Message

func (cr CodeResult) String() string {
	var b strings.Builder
	for _, m := range cr {
		if m.Content != "" {
			b.WriteString(m.String())
		}
	}
	return b.String()
}

type CodeFlags int64

const (
	Asynchronous CodeFlags = 1 << iota
	IsPreProcessed
	IsPostProcessed
	IsFromMacro
	IsNestedMacro
	IsFromConfig
	IsFromConfigOverride
	EnforceAbsolutePosition
	IsPrioritized
	CodeFlagsNone = 0
)

type KeywordType byte

const (
	KeywordTypeNone KeywordType = iota
	If
	ElseIf
	Else
	While
	Break
	Return
	Abort
	Var
	Set
)

type Code struct {
	BaseCommand
	Result          CodeResult
	Type            types.CodeType
	Channel         types.CodeChannel
	LineNumber      *int64
	Indent          byte
	Keyword         KeywordType
	KeywordArgument string
	MajorNumber     *int64
	MinorNumber     *int8
	Flags           CodeFlags
	Comment         string
	FilePosition    *int64
	Parameters      []CodeParameter
}

func NewCode() Code {
	return Code{
		Type:    types.Comment,
		Channel: types.DefaultChannel,
		Keyword: KeywordTypeNone,
		Flags:   CodeFlagsNone,
	}
}

func (c *Code) Parameter(letter string) *CodeParameter {
	l := strings.ToUpper(letter)
	for _, p := range c.Parameters {
		if l == strings.ToUpper(p.Letter) {
			return &p
		}
	}
	return nil
}

func (c *Code) ParameterOrDefault(letter string, value interface{}) *CodeParameter {
	p := c.Parameter(letter)
	if p != nil {
		return p
	}
	return NewSimpleCodeParameter(letter, value)
}

func (c *Code) GetUnprecedentedString(quote bool) string {
	var b strings.Builder
	for _, p := range c.Parameters {
		if b.Len() > 0 {
			b.WriteString(" ")
		}
		b.WriteString(p.Letter)
		if quote && p.IsString {
			b.WriteString(`"`)
		}
		b.WriteString(p.AsString())
		if quote && p.IsString {
			b.WriteString(`"`)
		}
	}
	return b.String()
}

func (c Code) String() string {
	if c.Type == types.Comment {
		return ";" + c.Comment
	}
	var b strings.Builder
	b.WriteString(c.ShortString())

	for _, p := range c.Parameters {
		b.WriteString(" ")
		b.WriteString(p.String())
	}

	if c.Comment != "" {
		b.WriteString(" ;")
		b.WriteString(c.Comment)
	}

	if len(c.Result) > 0 {
		b.WriteString("\n")
		b.WriteString(strings.TrimRight(c.Result.String(), " "))
	}

	return b.String()
}

func (c Code) ShortString() string {
	if c.Type == types.Comment {
		return "(comment)"
	}

	if c.MinorNumber != nil {
		return fmt.Sprintf("%s%d.%d", c.Type, *c.MajorNumber, *c.MinorNumber)
	}
	return fmt.Sprintf("%s%d", c.Type, *c.MajorNumber)
}

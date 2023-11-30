package helpers

import (
	"bytes"
	"fmt"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

func HighlightCode(pasteContent string) (string, bool) {
	var buf bytes.Buffer
	lexer := lexers.Analyse(pasteContent)
	if lexer == nil {
		return pasteContent, false
	}
	lexer = chroma.Coalesce(lexer)

	fmt.Println(lexer.Config().Name)
	style := styles.Get("igor")

	formatter := html.New(html.Standalone(false))

	iterator, err := lexer.Tokenise(nil, pasteContent)
	if err != nil {
		return pasteContent, false
	}

	if err := formatter.Format(&buf, style, iterator); err != nil {
		return pasteContent, false
	}

	return buf.String(), true
}

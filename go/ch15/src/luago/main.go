package main

import (
	"fmt"
	"io/ioutil"
	_lexer "luago/compiler/lexer"
	"os"
)

func testLexer(chunk, chunkName string) {
	lexer := _lexer.NewLexer(chunk, chunkName)
	for {
		line, kind, token := lexer.NextToken()
		fmt.Printf("[%2d] [%-10s] %s\n", line, kindToCategory(kind), token)
		if kind == _lexer.TOKEN_EOF {
			break
		}
	}
}

func kindToCategory(kind int) string {
	switch {
	case kind < _lexer.TOKEN_SEP_SEMI:
		return "other"
	case kind <= _lexer.TOKEN_SEP_RCURLY:
		return "separator"
	case kind <= _lexer.TOKEN_OP_NOT:
		return "operator"
	case kind <= _lexer.TOKEN_KW_WHILE:
		return "keyword"
	case kind <= _lexer.TOKEN_IDENTIFIER:
		return "identifier"
	case kind <= _lexer.TOKEN_NUMBER:
		return "number"
	case kind <= _lexer.TOKEN_STRING:
		return "string"
	default:
		return "other"
	}
}

func main() {
	if len(os.Args) > 1 {
		data, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		testLexer(string(data), os.Args[1])
	}
}

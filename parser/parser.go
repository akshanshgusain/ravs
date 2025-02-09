package parser

import (
	"ravs_lang/ast"
	"ravs_lang/lexer"
	"ravs_lang/token"
)

type Parser struct {
	l         *lexer.Lexer
	currToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// read two tokens, so currToken and peekToken are both set
	p.nextToke()
	p.nextToke()

	return p
}

func (p *Parser) nextToke() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}

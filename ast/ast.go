package ast

import "ravs_lang/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node // Embedding Node, composition of Node interface
	statementNode()
}

type Expression interface {
	Node // Embedding Node, composition of Node interface
	expressionNode()
}

// implementations of various types of Nodes:

// Program is the root node of every AST
type Program struct {
	Statements []Statement // very program is series of statements(Nodes of statement type)
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// implementations of various types of Statement and Expression Nodes:

type LetStatement struct {
	Token token.Token // Token.LET
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
func (ls *LetStatement) statementNode() {}

type Identifier struct {
	Token token.Token // Token.IDENT
	Value string
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (i *Identifier) expressionNode() {}

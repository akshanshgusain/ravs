package ast

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

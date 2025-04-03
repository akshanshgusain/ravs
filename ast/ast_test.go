package ast

import (
	"ravs_lang/token"
	"testing"
)

func TestString(t *testing.T) {
	p := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "lastName"},
					Value: "lastName",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "Gusain"},
					Value: "Gusain",
				},
			},
		},
	}
	if p.String() != "let lastName = Gusain;" {
		t.Errorf("p.String() wrong. got=%q", p.String())
	}
}

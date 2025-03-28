package parser

import (
	"ravs_lang/ast"
	"ravs_lang/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram returned nil program")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Error("s.TokenLiteral not 'let'. got=", s.TokenLiteral())
		return false
	}
	letSmt, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("s.Statement is not ast.LetStatement. got=%T", s)
	}
	if letSmt.Name.Value != name {
		t.Errorf("letSmt.Name.Value not '%s'. got=%s", name, letSmt.Name.Value)
		return false
	}
	if letSmt.Name.TokenLiteral() != name {
		t.Errorf("letSmt.Name.TokenLiteral not '%s'. got=%s",
			name, letSmt.Name.TokenLiteral())
		return false
	}
	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

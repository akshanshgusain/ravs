package lexer

import "ravs_lang/token"

// Lexer readPosition always points to the next position where we're going to read from next
// position always points to the position where we last read

type Lexer struct {
	input        string
	position     int // current position in input
	readPosition int //current reading position in input (points to the next character in the input)
	currentChar  byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar gives us the next character and advance our position in the input string
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.currentChar = 0 // ASCII NUL
	} else {
		l.currentChar = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// NextToken returns Token based on the current character
func (l *Lexer) NextToken() token.Token {
	var t token.Token
	l.eatWhitespace()
	switch l.currentChar {
	case '=':
		t = token.Token{Type: token.ASSIGN, Literal: string(l.currentChar)}
	case '/':
		t = token.Token{Type: token.SLASH, Literal: string(l.currentChar)}
	case '*':
		t = token.Token{Type: token.ASTERISK, Literal: string(l.currentChar)}
	case '<':
		t = token.Token{Type: token.LT, Literal: string(l.currentChar)}
	case '>':
		t = token.Token{Type: token.GT, Literal: string(l.currentChar)}
	case ';':
		t = token.Token{Type: token.SEMICOLON, Literal: string(l.currentChar)}
	case ',':
		t = token.Token{Type: token.COMMA, Literal: string(l.currentChar)}
	case '{':
		t = token.Token{Type: token.LBRACE, Literal: string(l.currentChar)}
	case '}':
		t = token.Token{Type: token.RBRACE, Literal: string(l.currentChar)}
	case '(':
		t = token.Token{Type: token.LPAREN, Literal: string(l.currentChar)}
	case ')':
		t = token.Token{Type: token.RPAREN, Literal: string(l.currentChar)}
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:
		if isLetter(l.currentChar) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdentifier(t.Literal)
			return t
		} else if isDigit(l.currentChar) {
			t.Literal = l.readNumber()
			if l.currentChar == '.' {
				l.readChar()
				if isDigit(l.currentChar) {
					t.Literal += "." + l.readNumber()
					t.Type = token.FLOAT
				} else {
					t.Type = token.ILLEGAL
				}
			} else {
				t.Type = token.INT
			}
			return t
		} else {
			t = token.Token{Type: token.ILLEGAL, Literal: string(l.currentChar)}
		}
	}
	return t
}

func (l *Lexer) readIdentifier() string {
	p := l.position
	for isLetter(l.currentChar) {
		l.readChar()
	}
	return l.input[p:l.position]
}

func isLetter(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_'
}

func (l *Lexer) eatWhitespace() {
	for l.currentChar == ' ' || l.currentChar == '\t' || l.currentChar == '\n' || l.currentChar == '\r' {
		l.readChar()
	}
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func (l *Lexer) readNumber() string {
	p := l.position
	for isDigit(l.currentChar) {
		l.readChar()
	}
	return l.input[p:l.position]
}

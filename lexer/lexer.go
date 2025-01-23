package lexer

type Lexer struct {
	input        string
	position     int // current position in input
	readPosition int //current reading position in input (points to the next character in the input)
	currentChar  byte
}

func New(input string) *Lexer {
	return &Lexer{input: input}
}

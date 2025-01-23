package lexer

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

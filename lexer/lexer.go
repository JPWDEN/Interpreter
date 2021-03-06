package lexer

import (
	"strings"

	"github.com/Interpreter/token"
)

//Lexer struct
type Lexer struct {
	input        string
	position     int  //Current position in input string above
	readPosition int  //Reading position for looking forward in input string after ch
	ch           byte //Current char/byte under examination
}

//New returns a new Lexer object
func NewLexer(inputSourceCode string) *Lexer {
	l := &Lexer{input: inputSourceCode}
	l.readChar()
	return l
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	//All of the alphabetic letters we want to allow for use in identifiers and keywords
	//Currently:  a-z, A-Z, _, and .    Potentially add:  0-9, *, &
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '.'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 //ASCII null character to signify end of input
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	if l.ch == '.' { //Process remaining digits of float in the case we find a '.'
		for isDigit(l.ch) {
			l.readChar()
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

//NextToken returns a struct for the next token (literal and actual value) in the input
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhiteSpace()
	switch l.ch {
	//Indentifiers and Literals
	case '=':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.EQ, Literal: string(l.ch) + string(l.peekChar())}
			l.readChar()
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.NOTEQ, Literal: string(l.ch) + string(l.peekChar())}
			l.readChar()
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
		tok = newToken(token.BANG, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LESS, l.ch)
	case '>':
		tok = newToken(token.GREATER, l.ch)
	case '^':
		tok = newToken(token.POINTER, l.ch)

	//Delimeters
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			if strings.Contains(tok.Literal, ".") {
				tok.Type = token.FLOAT
			} else {
				tok.Type = token.INT
			}
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

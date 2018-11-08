package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//Identifiers and literals
	IDENT = "IDENT" //Function names, variables
	INT   = "INT"   //Integer number
	FLOAT = "FLOAT" //32-bit floating point number

	//Operators
	ASSIGN   = "="
	ASTERISK = "*"
	BANG     = "!"
	EQ       = "=="
	GREATER  = ">"
	LESS     = "<"
	MINUS    = "-"
	NOTEQ    = "!="
	PLUS     = "+"
	POINTER  = "^" //Pointer to memory address
	SLASH    = "/"

	//Delimeters
	COMMA     = ","
	LBRACE    = "{"
	LPAREN    = "("
	RBRACE    = "}"
	RPAREN    = ")"
	SEMICOLON = ";"

	//Keywords
	ELSE     = "ELSE"
	FALSE    = "FALSE"
	FUNCTION = "FUNCTION"
	IF       = "IF"
	LET      = "LET"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
)

//TokenType value holds to type of token we are dealing  (integer, brace, comma, etc)
type TokenType string

//Token is the struct that defines the type fo token and the actual value of the token
type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"else":   ELSE,
	"false":  FALSE,
	"fn":     FUNCTION,
	"if":     IF,
	"let":    LET,
	"return": RETURN,
	"true":   TRUE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

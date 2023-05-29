package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//identifiers + literals
	IDENT = "IDENT" // add, foo, bar, x y
	INT   = "INT"   //123

	//Operator
	ASSIGN = "="
	PLUS   = "+"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPATEN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Keywords
	LET      = "LET"
	FUNCTION = "FUNCTION"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if t, ok := keywords[ident]; ok {
		return t
	}
	return IDENT
}

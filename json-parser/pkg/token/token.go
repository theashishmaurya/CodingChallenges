package token

const (
	// Token/character we don't know about
	Illegal Type = "ILLEGAL"

	// End of file
	EOF Type = "EOF"

	// Literals
	String Type = "STRING"
	Number Type = "NUMBER"

	// The six structural tokens
	LeftBrace    Type = "{"
	RightBrace   Type = "}"
	LeftBracket  Type = "["
	RightBracket Type = "]"
	Comma        Type = ","
	Colon        Type = ":"

	// Structural
	Whitespace Type = "WHITESPACE"

	// Comments are not valid in JSON
	// LineComment  Type = "//"
	// BlockComment Type = "/*"

	// Values
	True  Type = "TRUE"
	False Type = "FALSE"
	Null  Type = "NULL"
)

// Structure of the tokens
type Type string

type Token struct {
	Type    Type
	Literal string
	Line    int
	Start   int
	End     int
}

func main() {

	// Supose to get the data someHow.

	dataStream := "JSON"

	// Read bit by bit Okay...

}

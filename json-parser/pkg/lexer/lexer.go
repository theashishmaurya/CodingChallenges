package lexer

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

}

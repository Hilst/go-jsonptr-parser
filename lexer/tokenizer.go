package lexer

type Tokenizer struct {
	position     int
	readPosition int
	ch           rune
	input        string
}

const _0 = int('0')
const _9 = int('9')
const dot = int('.')

const a = int('a')
const z = int('z')
const A = int('A')
const Z = int('Z')
const __ = int('_')
const bar = int('/')
const ddot = int(':')

func isNumber(ch rune) bool {
	var char = int(ch)
	return _0 <= char && char <= _9 || char == dot
}
func isLetter(ch rune) bool {
	var char = int(ch)
	return a <= char && z >= char ||
		A <= char && Z >= char ||
		char == __ ||
		char == bar ||
		char == ddot
}

var keywords = map[string]Token{
	"where:": CreateToken(Where, "where:"),
	"index":  CreateToken(Index, "index"),
}

func InitTokenizer(input string) Tokenizer {
	var t = Tokenizer{
		position:     0,
		readPosition: 0,
		input:        input,
	}
	t.readChar()
	return t
}
func (tzer *Tokenizer) NextToken() Token {
	for tzer.ch == ' ' {
		tzer.readChar()
	}

	var t Token
	var tNil bool = true
	switch tzer.ch {
	case '{':
		t = CreateToken(LSql, string(tzer.ch))
		tNil = false
	case '}':
		t = CreateToken(RSql, string(tzer.ch))
		tNil = false
	case '=':
		t = CreateToken(Equal, string(tzer.ch))
		tNil = false
	case '>':
		t = CreateToken(Greater, string(tzer.ch))
		tNil = false
	case '<':
		t = CreateToken(Less, string(tzer.ch))
		tNil = false
	case '\'':
		t = CreateToken(SQuote, string(tzer.ch))
		tNil = false
	case '\x00':
		t = CreateToken(EOF, "eof")
		tNil = false
	}

	if isLetter(tzer.ch) {
		var ident = tzer.readIdent()
		var word, exists = keywords[ident]
		if exists {
			return word
		} else {
			return CreateToken(Ident, ident)
		}
	} else if isNumber(tzer.ch) {
		return CreateToken(Number, tzer.readNumber())
	} else if tNil {
		return CreateToken(Illegal, string(tzer.ch))
	}

	tzer.readChar()
	return t
}

func (tzer *Tokenizer) readChar() {
	if tzer.readPosition >= len(tzer.input) {
		tzer.ch = '\x00' // NULL TERMINATION
	} else {
		tzer.ch = rune(tzer.input[tzer.readPosition])
	}
	tzer.position = tzer.readPosition
	tzer.readPosition++
}
func (tzer *Tokenizer) readIdent() string {
	var first = tzer.position
	for isLetter(tzer.ch) {
		tzer.readChar()
	}
	return tzer.input[first:tzer.position]
}
func (tzer *Tokenizer) readNumber() string {
	var first = tzer.position
	for isNumber(tzer.ch) {
		tzer.readChar()
	}
	return tzer.input[first:tzer.position]
}

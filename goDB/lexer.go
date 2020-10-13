package gosql

import (
	"fmt"
	"strings"
)

// define struct to handle location of tokens in string
type location struct {
	line uint
	col  uint
}

// Define token keywords for querying
type keyword string

const (
	selectKeyword     keyword = "select"
	fromKeyword       keyword = "from"
	asKeyword         keyword = "as"
	tableKeyword      keyword = "table"
	createKeyword     keyword = "create"
	dropKeyword       keyword = "drop"
	insertKeyword     keyword = "insert"
	intoKeyword       keyword = "into"
	valuesKeyword     keyword = "values"
	intKeyword        keyword = "int"
	textKeyword       keyword = "text"
	boolKeyword       keyword = "boolean"
	whereKeyword      keyword = "where"
	andKeyword        keyword = "and"
	orKeyword         keyword = "or"
	trueKeyword       keyword = "true"
	falseKeyword      keyword = "false"
	uniqueKeyword     keyword = "unique"
	indexKeyword      keyword = "index"
	onKeyword         keyword = "on"
	primarykeyKeyword keyword = "primary key"
	nullKeyword       keyword = "null"
)

// define token match for symbols
type symbol string

const (
	semicolonSymbol  symbol = ";"
	asteriskSymbol   symbol = "*"
	commaSymbol      symbol = ","
	leftParenSymbol  symbol = "("
	rightParenSymbol symbol = ")"
	eqSymbol         symbol = "="
	neqSymbol        symbol = "<>"
	neqSymbol2       symbol = "!="
	concatSymbol     symbol = "||"
	plusSymbol       symbol = "+"
	ltSymbol         symbol = "<"
	lteSymbol        symbol = "<="
	gtSymbol         symbol = ">"
	gteSymbol        symbol = ">="
)

// Identifies the kind of token
type tokenKind uint

const (
	keywordKind tokenKind = iota
	symbolKind
	identifierKind
	stringKind
	numericKind
	boolKind
	nullKind
)

// object for handling tokens
type token struct {
	value string
	kind  tokenKind
	loc   location
}

// cursor to track position in the lexer input
type cursor struct {
	pointer uint
	loc     location
}

// takes in token and checks if it equals another
func (t *token) equals(other *token) bool {
	return t.value == other.value && t.kind == other.kind
}

// ? This is confusing
type lexer func(string, cursor) (*token, cursor, bool)

// main function that takes a string and returns an array of token structs and and error message
func lex(source string) ([]*token, error) {
	// initialize empty array of token structs
	tokens := []*token{}

	// init cursor
	cur := cursor{}

	// ? key...mayvbe
lex:
	// iterate over a the source string
	for cur.pointer < uint(len(source)) {
		// create an array of lexers that will have the lexing analysis done on a token in the source
		lexers := []lexer{lexKeywords, lexSymbols, lexString, lexNumeric, lexIdentifier}
		//  gets each lexer from the array, if analysis works, append token, else continue
		for _, l := range lexers {
			if token, newCursor, ok := l(source, cur); ok {
				cur = newCursor

				if token != nil {
					tokens = append(tokens, token)
				}
				continue lex
			}
		}
		hint := ""
		if len(tokens) > 0 {
			hint = " after " + tokens[len(tokens)-1].value
		}
		return nil, fmt.Errorf("Unable to lex token%s, at %d:%d", hint, cur.loc.line, cur.loc.col)
	}
	return tokens, nil
}

// analyze if token is numeric
func isNumeric(source string, ic cursor) (*token, cursor, bool) {
	// get the current position
	cur := ic

	// mark if punctuation found
	periodFound := false
	expMarkerFound := false

	for ; cur.pointer < uint(len(source)); cur.pointer++ {
		c := source[cur.pointer]
		cur.loc.col++

		isDigit := c >= '0' && c <= '9'
		isPeriod := c == '.'
		isExpMarker := c == 'e'

		// must start with a digit or period
		if cur.pointer == ic.pointer {
			if !isDigit && !isPeriod {
				return nil, ic, false
			}

			periodFound = isPeriod
			continue
		}

		if isPeriod {
			if periodFound {
				return nil, ic, false
			}

			periodFound = true
			continue
		}

		if isExpMarker {
			if expMarkerFound {
				return nil, ic, false
			}

			// No periods allowed after expMarker
			periodFound = true
			expMarkerFound = true

			// expMarker must be followed by digits
			if cur.pointer == uint(len(source)-1) {
				return nil, ic, false
			}

			cNext := source[cur.pointer+1]

			if cNext == '-' || cNext == '+' {
				cur.pointer++
				cur.loc.col++
			}

			continue
		}

		if !isDigit {
			break
		}
	}

	// No characters accumulated
	if cur.pointer == ic.pointer {
		return nil, ic, false
	}

	return &token{
		value: source[ic.pointer:cur.pointer],
		loc:   ic.loc,
		kind:  numericKind,
	}, cur, true
}

// lexer for characters
func lexCharacterDelimited(source string, ic cursor, delimiter byte) (*token, cursor, bool) {
	cur := ic

	// check if source string is empty
	if len(source[cur.pointer:]) == 0 {
		return nil, ic, false
	}

	// if the source str
	if source[cur.pointer] != delimiter {
		return nil, ic, false
	}

	// increment pointer locations and source string location
	cur.loc.col++
	cur.pointer++

	// create an array for char bytes
	var value []byte
	// iterate over souurce string
	for ; cur.pointer < uint(len(source)); cur.pointer++ {
		c := source[cur.pointer]

		if c == delimiter {
			// SQL escapes are via double characters, not backslash.
			if cur.pointer+1 >= uint(len(source)) || source[cur.pointer+1] != delimiter {
				return &token{
					value: string(value),
					loc:   ic.loc,
					kind:  stringKind,
				}, cur, true
			} else {
				value = append(value, delimiter)
				cur.pointer++
				cur.loc.col++
			}
		}

		value = append(value, c)
		cur.loc.col++
	}

	return nil, ic, false
}

func lexString(source string, ic cursor) (*token, cursor, bool) {
	return lexCharacterDelimited(source, ic, '\'')
}

func lexSymbol(source string, ic cursor) (*token, cursor, bool) {
	// point c at the appropriate part of
	c := source[ic.pointer]
	// pointer to curser object
	cur := ic

	// Will get overwritten later if not an ignored syntax
	cur.pointer++
	cur.loc.col++

	// Syntax that should be thrown away
	switch c {
	case '\n':
		cur.loc.line++
		cur.loc.col = 0
		fallthrough
	case '\t':
		fallthrough
	case ' ':
		return nil, cur, true
	}

	// syntax to keep
	symbols := []symbol{
		commaSymbol,
		leftParenSymbol,
		rightParenSymbol,
		semicolonSymbol,
		asteriskSymbol,
	}

	var options []string
	for _, s := range symbols {
		options = append(options, string(s))
	}

	// use 'ic' not cur....interesting
	match := longestMatch(source, ic, options)

	// unknown chars
	if match == "" {
		return nil, ic, false
	}

	// current pointer is set to
	cur.pointer = ic.pointer + uint(len(match))
	cur.loc.col = ic.loc.col + uint(len(match))

	return &token{
		value: match,
		loc:   ic.loc,
		kind:  symbolKind,
	}, cur, true
}

func lexKeyword(source string, ic cursor) (*token, cursor, bool) {
	cur := ic              // get your cursor
	keywords := []keyword{ // array of keywords
		selectKeyword,
		insertKeyword,
		valuesKeyword,
		tableKeyword,
		createKeyword,
		whereKeyword,
		fromKeyword,
		intoKeyword,
		textKeyword,
	}

	var options []string
	for _, k := range keywords { // putting keywords in options
		options = append(options, string(k))
	}

	match := longestMatch(source, ic, options)
	if match == "" {
		return nil, ic, false
	}

	// inc cursor
	cur.pointer = ic.pointer + uint(len(match))
	cur.loc.col = ic.loc.col + uint(len(match))

	// kind of token
	kind := keywordKind
	if match == string(trueKeyword) || match == string(falseKeyword) {
		kind = boolKind
	}

	// return  the token, type and location
	return &token{
		value: match,
		kind:  kind,
		loc:   ic.loc,
	}, cur, true
}

// longestMatch iterates through a source string starting at the given
// cursor to find the longest matching substring among the provided options
func longestMatch(source string, ic cursor, options []string) string {
	var value []byte
	var skipList []int
	var match string

	cur := ic

	// start at the string its pointing at
	for cur.pointer < uint(len(source)) {
		// value will  be lower case version of source string starting at the pointer location
		value = append(value, strings.ToLower(string(source[cur.pointer]))...)
		cur.pointer++

	match:
		for i, option := range options {
			for _, skip := range skipList {
				if i == skip {
					continue match
				}
			}

			// deal with case like INT v INTO
			if option == string(value) {
				skipList = append(skipList, i)
				if len(option) > len(match) {
					match = option
				}

				continue
			}

			sharesPrefix := string(value) == option[:cur.pointer-ic.pointer]
			tooLong := len(value) > len(option)
			if tooLong || !sharesPrefix {
				skipList = append(skipList, i)
			}
		}

		if len(skipList) == len(options) {
			break
		}

	}
	return match
}

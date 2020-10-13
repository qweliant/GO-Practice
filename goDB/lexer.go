package gosql


import (
	"fmt"
	"strings"
)


// define struct to handle location of tokens in string
type location struct {
	line uint
	col uint
}


// Define token keywords for querying
type keyword string


const (
    selectKeyword keyword = "select"
    fromKeyword   keyword = "from"
    asKeyword     keyword = "as"
    tableKeyword  keyword = "table"
    createKeyword keyword = "create"
    insertKeyword keyword = "insert"
    intoKeyword   keyword = "into"
    valuesKeyword keyword = "values"
    intKeyword    keyword = "int"
    textKeyword   keyword = "text"	
)

// define token match for symbols
type symbol string

const (
	semicolonSymbol  symbol = ";"
    asteriskSymbol   symbol = "*"
    commaSymbol      symbol = ","
    leftparenSymbol  symbol = "("
    rightparenSymbol symbol = ")"
)

// Identifies the kind of token
type tokenKind uint

const (
	keywordKind tokenKind = iota
    symbolKind
    identifierKind
    stringKind
    numericKind
)

// object for handling tokens
type token struct{
	value string
	kind tokenKind
	loc location
}

// cursor to track position in the lexer input
type cursor struct {
	pointer uint
	loc location
}

// takes in token and checks if it equals another
func (t *token) equals(other *token) bool {
	return t.value == other.value && t.kind == other.kind
}


// ? This is confusing
type lexer func(string, cursor) (*token, cursor, bool)


// main function that takes a string and returns an array of token structs and and error message
func lex(source string) ([]*token, error){
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

				if token != nil{
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
func isNumeric(source string, ic cursor) (*token, cursor, bool){
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
			if !isDigit && !isPeriod{
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

	return &token {
		value: source[ic.pointer:cur.pointer],
		loc: ic.loc,
		kind: numericKind,
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

    cur.loc.col++
    cur.pointer++

    var value []byte
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
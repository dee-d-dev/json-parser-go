package main

import (
	"fmt"
	"regexp"
	"strings"
)

    const (
        LEFT_BRACE = "LEFT BRACE"
        RIGHT_BRACE = "RIGHT BRACE"
        LEFT_BRACKET = "LEFT BRACKET"
        RIGHT_BRACKET = "RIGHT BRACKET"
        COMMA = "COMMA"
        COLON = "COLON"
        STRING = "STRING"
        NUMBER = "NUMBER"
        TRUE = "TRUE"
        FALSE = "FALSE"
        NULL = "NULL"
    )

    type Token struct {
        value string
        tokenType string
    }

    func Lexer(s string) []Token {
        currentPosition := 0
        tokens := []Token{}
    
        for currentPosition < len(s) {
            char := string(s[currentPosition])
            switch char {
                case "{":
                    tokens = append(tokens, Token{value: char, tokenType: LEFT_BRACE})
                    currentPosition++
                case "}":
                    tokens = append(tokens, Token{value: char, tokenType: RIGHT_BRACE})
                    currentPosition++
                case "[":
                    tokens = append(tokens, Token{value: char, tokenType: LEFT_BRACKET})
                    currentPosition++
                case "]":
                    tokens = append(tokens, Token{value: char, tokenType: RIGHT_BRACKET})
                    currentPosition++
                case ",":
                    tokens = append(tokens, Token{value: char, tokenType: COMMA})
                    currentPosition++
                case ":":
                    tokens = append(tokens, Token{value: char, tokenType: COLON})
                    currentPosition++
                case `"`:
                    stringSequence := ""
                    currentPosition++

                    for currentPosition < len(s) {
                        if string(s[currentPosition]) == `"` {
                            currentPosition++
                            break
                        }

                        stringSequence += string(s[currentPosition])
                        currentPosition++
                    }
                    tokens = append(tokens, Token{value: stringSequence, tokenType: STRING})

                default:
                    if regexp.MustCompile(`\s`).MatchString(char) {
                        currentPosition++
                        continue
                    } else if regexp.MustCompile(`[0-9]`).MatchString(char) {
                        numberSequence := ""
                        for currentPosition < len(s) {
                            if regexp.MustCompile(`[0-9]`).MatchString(string(s[currentPosition])) {
                                numberSequence += string(s[currentPosition])
                                currentPosition++
                            } else {
                                break
                            }
                        }
                        tokens = append(tokens, Token{value: numberSequence, tokenType: NUMBER})

                    } else if s[currentPosition:currentPosition+4] == "true" {
                        tokens = append(tokens, Token{value: "true", tokenType: TRUE})
                        currentPosition += 4
                    } else if s[currentPosition:currentPosition+5] == "false" {
                        tokens = append(tokens, Token{value: "false", tokenType: FALSE})
                        currentPosition += 5
                    } else if s[currentPosition:currentPosition+4] == "null" {
                        tokens = append(tokens, Token{value: "null", tokenType: NULL})
                        currentPosition += 4
                    } else {
                        panic("Invalid character")
                    }
                
            }
        }
    
        return tokens
    }

    type ObjectNode struct {
        nodeType string
        value    Token
        children map[string]*ObjectNode
    }

    func buildNode(tokenType string, value string) *ObjectNode {
        return &ObjectNode{
            nodeType: strings.ToLower(tokenType),
                value: Token{
                tokenType: tokenType,
                value:     value,
            },
            children: make(map[string]*ObjectNode),
        }
    }

    func parseValue(tokens []Token, currentPosition *int) *ObjectNode {
        currentTokenType := tokens[*currentPosition].tokenType
        switch currentTokenType {
        case STRING, NUMBER, TRUE, FALSE:
            node := buildNode(currentTokenType, tokens[*currentPosition].value)
            *currentPosition++
            return node
        case LEFT_BRACE:
            return parseObject(tokens, currentPosition)
        default:
            panic("Not token to parse")
        }
    }

    func parseObject(tokens []Token, currentPosition *int) *ObjectNode {
        *currentPosition++

        objectNode := &ObjectNode{
            nodeType: "object",
            children: make(map[string]*ObjectNode),
        }

        for tokens[*currentPosition].tokenType != RIGHT_BRACE {
            if tokens[*currentPosition].tokenType == COMMA {
                *currentPosition++
            }

            if tokens[*currentPosition].tokenType == STRING {
                key := tokens[*currentPosition].value
                *currentPosition++

                if tokens[*currentPosition].tokenType != COLON {
                panic("Invalid JSON, expected a colon")
                }
                *currentPosition++

                value := parseValue(tokens, currentPosition)

                objectNode.children[key] = value
            } else {
                panic("Invalid JSON, expected a string key")
            }

            *currentPosition++

            if *currentPosition >= len(tokens) {
                break
            }
        }

        return objectNode
    }

    func Parser(tokens []Token) *ObjectNode {
        if len(tokens) == 0 {
            panic("No tokens to parse")
        }

        i := 0
        return parseValue(tokens, &i)
    }

    func main() {
        parser := Lexer(`{ "name": "John", "age": 30, "isStudent": false`)

        fmt.Println(parser)

        parsed := Parser(parser)
        fmt.Println(parsed)
    }
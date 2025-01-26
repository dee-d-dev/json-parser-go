package main

import (
	"fmt"
	"regexp"
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

    func main() {
        parser := Lexer(`{
            "key1": true,
            "key2": false,
            "key3": null,
            "key4": "value",
            "key5": 101
            }`)
        fmt.Println(parser)
    }
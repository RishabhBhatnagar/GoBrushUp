// 2) Conversion of Infix to Postfix.

package data_structures

import (
	"errors"
	"fmt"
	"strings"
)

const BRACKETPREC = 2946 // precedence value for brackets.
const OTHERCHAR = 3489

func getPrecedence(char rune) (int, bool) {
	/**
	 * returns the precedence value of the character.
	 * Higher the precedence, higher will be the value.
	 * Args:
	 * 	 char: character whose precedence is required.
	 * Returns:
	 *   int: precedence value.
	 *   bool represents isOperator
	 */
	switch char {
	case '+', '-':
		return 1, true
	case '/', '*', '%':
		return 2, true
	case '^':
		return 3, true
	case '(', ')':
		return BRACKETPREC, false
	default:
		return OTHERCHAR, false
	}
	return -1, false
}

func Infix2postfix(infix string) (postfix string, err error) {
	/**
	 * Converts given infix expression to postfix.
	 */
	infix = fmt.Sprintf("(%s)", infix)          // adding brackets to the expression.
	infix = strings.Replace(infix, " ", "", -1) // removing all the spaces.
	st := stack{stackSize: len(infix)}
	var sb strings.Builder
	for _, char := range infix {
		if prec, isOp := getPrecedence(char); isOp {
			// an operator is encountered.
			if st.isEmpty(){
				st.push(int(char))   // push the first operator
			} else {
				prec2, isOp2 := getPrecedence(rune(st.peek()))
				if isOp2 {
					// if last appended char is an operator,
					if prec2 < prec {
						st.push(int(char))
					} else {
						// pop element from stack until element in the stack has lower precedence.
						for prec2, _ := getPrecedence(rune(st.peek())); !st.isEmpty(); {
							if prec2 < prec{
								st.push(int(char))
								break
							}
							sb.WriteRune(rune(st.pop())) // removing last entry in the stack
						}
						st.push(int(char))
					}
				} else {
					// last element in the stack is not an operator,
					// appending the current element as it is.
					st.push(int(char))
				}
			}
		} else if prec == BRACKETPREC {
			if char == '(' {
				// current character is an opening brace.
				st.push(int(char))
			} else {
				// closing bracket encountered
				// pop until opening brace is encountered.
				for st.peek() != int('(') {
					if st.isEmpty(){
						return "", errors.New("Imbalanced infix Expression")
					}
					sb.WriteRune(rune(st.pop()))
				}
				if st.isEmpty(){
					return "", errors.New("Imbalanced infix Expression")
				}
				st.pop()  // removing the opening brace.
			}
		} else {
			sb.WriteRune(char)
		}
	}
	if !st.isEmpty(){
		return "", errors.New("Imbalanced infix Expression")
	}
	return strings.ReplaceAll(sb.String(), "(", ""), nil
}

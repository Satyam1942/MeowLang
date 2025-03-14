package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var variables = make(map[string]string)
var encounteredIf = false
var insideIf = false
var executedIf = false

// to evaluate only binary arithmetic Expression
func evaluateArithmeticExpression(expression string) string {

	tokens := strings.Fields(expression)
	//case of string assigned to variable
	if tokens[0][0] == '"' {
		return expression
	}

	// if expression has only one variable
	if len(tokens) == 1 {
		val, exists := variables[tokens[0]]
		if exists {
			return val
		} else {
			return tokens[0]
		}
	}

	//invalid expressions
	if len(tokens) > 3 {
		panic("Invalid Cat Expression 😿!")
	}

	//boolean expressions
	var operand1 int
	var operand2 int
	var err1 error
	var err2 error

	val, exists := variables[tokens[0]]
	if exists {
		operand1, err1 = strconv.Atoi(val)
	} else {
		operand1, err1 = strconv.Atoi(tokens[0])
	}

	val2, exists2 := variables[tokens[2]]
	if exists2 {
		operand2, err2 = strconv.Atoi(val2)
	} else {
		operand2, err2 = strconv.Atoi(tokens[2])
	}

	if err1 != nil || err2 != nil {
		panic("Invalid Cat Expression 😿!")
	}

	operation := tokens[1]

	var result int

	switch operation {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		if operand2 == 0 {
			panic("Cats can't divide by Zero 😾!")
		}
		result = operand1 / operand2
	case "^":
		result = operand1 ^ operand2
	case "|":
		result = operand1 | operand2
	case "&":
		result = operand1 & operand2
	default:
		panic("Unknown Cat Operator 😾!")
	}

	return strconv.FormatInt(int64(result), 10)
}

func evaluateBooleanExpression(expression string) bool {
	tokens := strings.Fields(expression)
	// length is 1 means its either a constant or a variable which is always true
	if len(tokens) == 1 {
		return true
	}

	//invalid expressions
	if len(tokens) > 3 {
		panic("Invalid Cat Expression 😿!")
	}

	//boolean expressions
	var operand1 int
	var operand2 int
	var err1 error
	var err2 error

	val, exists := variables[tokens[0]]
	if exists {
		operand1, err1 = strconv.Atoi(val)
	} else {
		operand1, err1 = strconv.Atoi(tokens[0])
	}

	val2, exists2 := variables[tokens[2]]
	if exists2 {
		operand2, err2 = strconv.Atoi(val2)
	} else {
		operand2, err2 = strconv.Atoi(tokens[2])
	}

	if err1 != nil || err2 != nil {
		panic("Invalid Cat Expression 😿!")
	}

	operation := tokens[1]

	var result bool

	switch operation {
	case ">":
		result = operand1 > operand2
	case "<":
		result = operand1 < operand2
	case ">=":
		result = operand1 >= operand2
	case "<=":
		result = operand1 <= operand2
	case "==":
		result = operand1 == operand2
	case "!=":
		result = operand1 != operand2
	default:
		panic("Unknown Cat Operator 😾!")
	}

	return result
}

// 0 => returned if I want to process next line
// 1 => return if I want to skip execution of next line
func process(line string) int {
	tokens := strings.Fields(line)
	if len(tokens) == 0 {
		return 0
	}

	switch tokens[0] {
	case "meow":
		if !insideIf {
			encounteredIf = false
		} else {
			insideIf = false
		}

		if len(tokens) >= 4 && tokens[2] == "=" {
			variables[tokens[1]] = evaluateArithmeticExpression(strings.Join(tokens[3:], " "))
		} else {
			fmt.Println("Syntax error in meow 😿!!")
		}

	case "purrln":
		if !insideIf {
			encounteredIf = false
		} else {
			insideIf = false
		}

		if len(tokens) == 2 {
			val, exists := variables[tokens[1]]
			if exists {
				fmt.Println(val)
			} else {
				fmt.Println(tokens[1])
			}
		} else {
			fmt.Println(evaluateArithmeticExpression(strings.Join(tokens[1:], " ")))
		}
	case "purr":
		if !insideIf {
			encounteredIf = false
		} else {
			insideIf = false
		}

		if len(tokens) == 2 {
			val, exists := variables[tokens[1]]
			if exists {
				fmt.Print(val)
			} else {
				fmt.Print(tokens[1])
			}
		} else {
			fmt.Print(evaluateArithmeticExpression(strings.Join(tokens[1:], " ")))
		}
	case "meowIf":
		encounteredIf = true
		insideIf = true
		result := evaluateBooleanExpression(strings.Join(tokens[1:], " "))
		if result {
			//evaulate if portion => execute next line
			executedIf = true
			break
		} else {
			//evaluate else portion => skip the next line
			return 1
		}
	case "otherwiseScratch":
		if !encounteredIf {
			panic("Cats don't scratch without meowIf 😾!")
		}
		encounteredIf = false
		if executedIf {
			//already evaluated if portion => skip the next line
			executedIf = false
			return 1
		} else {
			break
		}
	case "//":
		break
	default:
		fmt.Println("🐱 Unknown command!")
	}

	return 0
}

func main() {
	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var startProcessing = false
	var skipLine = false

	for scanner.Scan() {
		line := scanner.Text()
		if skipLine {
			skipLine = false
			continue
		}

		line = strings.TrimSpace(line)
		if line == "Hello Kitty!" {
			startProcessing = true
			continue
		} else if line == "Bye Kitty!" {
			break
		}

		if startProcessing {
			retVal := process(line)
			if retVal == 1 {
				// skip the next line
				skipLine = true
			}
		}
	}
}

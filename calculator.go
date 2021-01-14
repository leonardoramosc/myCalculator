package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (calc) Operate(input string, op string) float64 {

	values := strings.Split(input, op)

	op1 := parse(values[0])
	op2 := parse(values[1])

	switch op {
	case "+":
		fmt.Println(op1 + op2)
		return op1 + op2
	case "-":
		fmt.Println(op1 - op2)
		return op1 - op2
	case "*":
		fmt.Println(op1 * op2)
		return op1 * op2
	case "/":
		fmt.Println(op1 / op2)
		return op1 / op2
	default:
		fmt.Println(op, " operator is not supported")
		return 0
	}
}

func formatInput(str string) string {
	splitted := strings.Split(strings.TrimSpace(str), " ")
	formatted := strings.Join(splitted, "")

	return formatted
}

func GetUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingresa operacion:")

	scanner.Scan()

	return formatInput(scanner.Text())
}

func GetOperator(operation string) string {
	operators := []string{"+", "-", "*", "/"}
	var op string

	for i := 0; i < len(operators); i++ {
		operator := operators[i]

		if strings.Contains(operation, operator) {
			index := strings.Index(operation, operator)
			splitted := strings.Split(operation, "")
			op = splitted[index]
		}
	}

	return op
}

func parse(str string) float64 {
	op, _ := strconv.ParseFloat(str, 64)

	return op
}

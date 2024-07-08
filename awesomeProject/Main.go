package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var dictionary = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

func Input_bus() (a, op, b string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Split(input, " ")
	if len(parts) < 3 {
		fmt.Println("Выдача паники, так как строка не является математической операцией.")
		return
	}
	if len(parts) != 3 {
		fmt.Println("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		return
	}
	a = parts[0]
	op = parts[1]
	b = parts[2]
	return
}

func Math_bus(a, op, b string) (string, string) {
	var err, output string
	var result int
	var num1, num2 bool
	result, num1 = dictionary[a] //проверка наличия римских цифр
	result, num2 = dictionary[b]
	if num1 == num2 {
		if num1 == true {
			a = strconv.Itoa(dictionary[a])
			b = strconv.Itoa(dictionary[b])
		} //конвертация римских при наличии
		aa, _ := strconv.Atoi(a) // конвертация с строки в число
		bb, _ := strconv.Atoi(b)
		if (op == "-" || op == "/") && num1 == true && (aa-bb <= 0 || aa/bb <= 0) {
			err = "Выдача паники, так как в римской системе нет отрицательных чисел."
			return err, output
		} // проверка на допустимость отрицательных римсчких чисел
		switch op {
		case "+":
			result = aa + bb
		case "-":
			result = aa - bb
		case "*":
			result = aa * bb
		case "/":
			result = aa / bb
		default:
			err = "Выдача паники, оператор не найден"
		}
	} else {
		err = "Выдача паники, так как используются одновременно разные системы счисления."
	}
	if num1 == true {
		output = intToRoman(result)
	} else {
		output = strconv.Itoa(result)
	}

	return output, err
}

func intToRoman(num int) string {
	var values = []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	var numerals = []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var result string

	for i, v := range values {
		count := int(num / v)
		result += strings.Repeat(numerals[i], count)
		num -= v * count
	}
	return result
}

func main() {
	fmt.Println(Math_bus(Input_bus()))
}

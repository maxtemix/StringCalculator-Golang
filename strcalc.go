package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var n1, operator string

	fmt.Scanf("%q%s", &n1, &operator)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n2 := scanner.Text()

	w1 := []rune(n1)
	w2 := []rune(n2)

	if n1 == "" {

		panic("Неправильный ввод первой строки")

	} else if len(w1) <= 10 && len(w2) <= 12 {

		b := rune(n2[0])
		e := rune(n2[len(n2)-1])

		switch operator {

		case "+":
			if b == '"' && e == '"' {

				n2 = n2[1 : len(n2)-1]

				fmt.Printf("\"%s%s\"", n1, n2)

			} else {

				panic("Неправильный ввод второй строки")
			}

		case "-":
			if b == '"' && e == '"' {

				n2 = n2[1 : len(n2)-1]

				n1 = strings.ReplaceAll(n1, n2, "")

				fmt.Printf("%q", n1)

			} else {

				panic("Неправильный ввод второй строки")
			}

		case "*":
			if n2 == "1" || n2 == "2" || n2 == "3" || n2 == "4" || n2 == "5" || n2 == "6" || n2 == "7" || n2 == "8" || n2 == "9" || n2 == "10" {

				h, _ := strconv.Atoi(n2)

				w1 = []rune(strings.Repeat(n1, h))

				if len(w1) > 40 {

					w1 = append(w1[:40])

					fmt.Printf("\"%s...\"", string(w1))

				} else {

					fmt.Printf("%q", string(w1))

				}

			} else {

				panic("Некорректное число")

			}

		case "/":
			if n2 == "1" || n2 == "2" || n2 == "3" || n2 == "4" || n2 == "5" || n2 == "6" || n2 == "7" || n2 == "8" || n2 == "9" || n2 == "10" {

				h, _ := strconv.Atoi(n2)

				w1 = w1[:len(w1)/h]

				fmt.Printf("%q", string(w1))

			} else {

				panic("Некорректное число")

			}

		default:
			panic("Некорректная арифметическая операция")

		}

	} else {

		panic("Некорректная длина введенных строк")

	}
}

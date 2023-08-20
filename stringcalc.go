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

	if n1 == "" {

		fmt.Print("Исключение. Работа программы завершена.")

	} else if len(n1) <= 10 && len(n2) <= 12 {

		b := rune(n2[0])
		e := rune(n2[len(n2)-1])

		switch operator {

		case "+":
			if b == '"' && e == '"' {

				n2 = n2[1 : len(n2)-1]

				fmt.Printf("\"%s%s\"", n1, n2)

			} else {

				fmt.Print("Исключение. Работа программы завершена.")
			}

		case "-":
			if b == '"' && e == '"' {

				n2 = n2[1 : len(n2)-1]

				n1 = strings.ReplaceAll(n1, n2, "")

				fmt.Printf("%q", n1)

			} else {

				fmt.Print("Исключение. Работа программы завершена.")
			}

		case "*":
			if n2 == "1" || n2 == "2" || n2 == "3" || n2 == "4" || n2 == "5" || n2 == "6" || n2 == "7" || n2 == "8" || n2 == "9" || n2 == "10" {

				h, _ := strconv.Atoi(n2)

				n1 = strings.Repeat(n1, h)

				if len(n1) > 40 {

					n1 = n1[:40]

					fmt.Printf("\"%s...\"", n1)

				} else {

					fmt.Printf("%q", n1)

				}

			} else {

				fmt.Print("Исключение. Работа программы завершена.")

			}

		case "/":
			if n2 == "1" || n2 == "2" || n2 == "3" || n2 == "4" || n2 == "5" || n2 == "6" || n2 == "7" || n2 == "8" || n2 == "9" || n2 == "10" {

				h, _ := strconv.Atoi(n2)

				n1 = n1[:len(n1)/h]

				fmt.Printf("%q", n1)

			} else {

				fmt.Print("Исключение. Работа программы завершена.")

			}

		default:
			fmt.Print("Исключение. Работа программы завершена.")

		}

	} else {

		fmt.Print("Исключение. Работа программы завершена.")

	}
}

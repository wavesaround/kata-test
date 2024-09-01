package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	// просим пользователя ввести математическое выражение, принимаем строку

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Пожалуйста, введите одно арифметическое выражение с двумя числами до 10. \nДля получения результата нажмите Enter")

	text, _ := reader.ReadString('\n')

	// убираем пробелы и табуляцию, меняем регистр

	spaces := regexp.MustCompile(`[\t\n\f\r ]+`)
	text = spaces.ReplaceAllString(text, "")
	text = strings.ToUpper(text)

	// поиск дробных чисел

	findDouble, _ := regexp.MatchString(`\d[.,]\d`, text)
	if findDouble {
		panic("calculator only works with integers")
	}

	// проверка на наличие и количество арифметических операторов

	if findOperator(text) != 1 {
		panic("more than 1 arithmetic operator or operator not found")
	}

	// проверка на смешивание чисел

	findMix, _ := regexp.MatchString(`[XVI].[0-9]|[0-9].[XVI]`, text)
	if findMix {
		panic("mixed digits")
	}

	// поиск чисел и оператора

	numbers, _ := regexp.Compile(`\d+|\w+|[+-–/*]`)
	twoNumbers := numbers.FindAllString(text, -1)

	// конвертируем римские в арабские, получаем целые числа
	// сразу проверяем, не встречаются ли среди них больше 10

	romeState, _ := regexp.MatchString(`[XVILCM]`, twoNumbers[0])

	var a int
	var b int

	if romeState {
		romeMoreTen1, _ := regexp.MatchString(`X{1}[XIV]{1,}|[LCM]`, twoNumbers[0])
		romeMoreTen2, _ := regexp.MatchString(`X{1}[XIV]{1,}|[LCM]`, twoNumbers[2])
		if romeMoreTen1 || romeMoreTen2 {
			panic("numbers are more the 10")
		}
		a = romeToArab(twoNumbers[0])
		b = romeToArab(twoNumbers[2])
	} else {
		a, _ = strconv.Atoi(twoNumbers[0])
		b, _ = strconv.Atoi(twoNumbers[2])
		if a > 10 || b > 10 {
			panic("numbers are more the 10")
		}
	}

	// производим арифметическую операцию

	res := 0

	switch twoNumbers[1] {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "–":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	default:
		panic("wrong operator")
	}

	// не допускаем нулевых и отрицательных значений, если цифры римские

	if romeState && res < 1 {
		panic("result is less than 1")
	}

	// если ввод был арабским, выводим результат, в противном случае переводим в римские

	if !romeState {
		fmt.Println("Результат:", res)
	} else {
		fmt.Println("Результат:", arabToRome(res))
	}
}

func findOperator(str string) int {
	operators := func(r rune) bool {
		return r == '+' || r == '-' || r == '*' || r == '/'
	}
	count := 0
	for i := 0; i < len(str); i++ {
		if strings.ContainsFunc(string(str[i]), operators) {
			count++
		}
	}
	return count
}

func romeToArab(x string) int {
	x = strings.ToUpper(x)
	arabToRome := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	return arabToRome[x]
}

func arabToRome(x int) string {

	arabToRomeUnit := map[int]string{
		0: "",
		1: "I",
		2: "II",
		3: "III",
		4: "IV",
		5: "V",
		6: "VI",
		7: "VII",
		8: "VIII",
		9: "IX",
	}

	arabToRomeTens := map[int]string{
		1:  "X",
		2:  "XX",
		3:  "XXX",
		4:  "XL",
		5:  "L",
		6:  "LX",
		7:  "LXX",
		8:  "LXXX",
		9:  "XC",
		10: "C",
	}

	if x < 11 {
		return arabToRomeUnit[x]
	} else {
		tens := x / 10
		units := x % 10
		return arabToRomeTens[tens] + arabToRomeUnit[units]
	}

}

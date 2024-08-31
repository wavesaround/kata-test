package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Пожалуйста, введите одно арифметическое выражение с двумя числами. \nДля получения результата нажмите Enter")

	text, _ := reader.ReadString('\n')

	// убираем пробелы и табуляцию

	spaces := regexp.MustCompile(`[\t\n\f\r ]+`)
	text = spaces.ReplaceAllString(text, "")
	fmt.Println(text)

	// проверка на наличие и количество арифметических операторов

	if findOperator(text) > 1 || findOperator(text) == 0 {
		panic("more than 1 arithmetic operator or operator not found")
	}

	// проверка на смешивание чисел

	findMix, _ := regexp.MatchString(`[XVIxvi].[0-9]|[0-9].[XVIxvi]`, text)
	if findMix {
		panic("mixed digits")
	}

	// поиск чисел

	re, _ := regexp.Compile(`\d+`)
	res := re.FindAllString(text, -1)
	fmt.Println(res) // [cat cat]



}

func findOperator(str string) int {
	operators := func(r rune) bool {
		return r == '+' || r == '-' || r == '*' || r == '/' || r == '='
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

	arabToRome := map[string]int{
		"I":     1,
		"II":    2,
		"III":   3,
		"IV":    4,
		"V":     5,
		"VI":    6,
		"VII":   7,
		"VIII":  8,
		"IX":    9,
		"X":     10,
		"XI":    11,
		"XII":   12,
		"XIII":  13,
		"XIV":   14,
		"XV":    15,
		"XVI":   16,
		"XVII":  17,
		"XVIII": 18,
		"XIX":   19,
		"XX":    20,
	}
	return arabToRome[x]
}

func arabToRome(x int) string {

	arabToRome := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
		11: "XI",
		12: "XII",
		13: "XIII",
		14: "XIV",
		15: "XV",
		16: "XVI",
		17: "XVII",
		18: "XVIII",
		19: "XIX",
		20: "XX",
	}
	return arabToRome[x]
}

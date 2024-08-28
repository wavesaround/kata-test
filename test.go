package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Пожалуйста, введите выражениe")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	toNumber, _ := strconv.Atoi(text)
	fmt.Println(arabToRome(toNumber))

}

package check

import (
	"fmt"
	"strings"
)

func Check(val string, generated string) {
	// fmt.Println(val, generated)

	var uppercase = strings.ToUpper(val)
	// var equality = strings.EqualFold(val, generated)
	// fmt.Println(equality)
	generatedArr := strings.Split(generated, "")
	promptedArr := strings.Split(uppercase, "")
	var posCorrect int

	if len(promptedArr) > 4 {
		fmt.Println("Enter 4 letters")
	} else if len(promptedArr) < 4 {
		fmt.Println("Enter 4 letters")
	}
	if generatedArr[0] == promptedArr[0] {
		posCorrect++
	}
	if generatedArr[1] == promptedArr[1] {
		posCorrect++
	}
	if generatedArr[2] == promptedArr[2] {
		posCorrect++
	}
	if generatedArr[3] == promptedArr[3] {
		posCorrect++
	}

	fmt.Println("On position: ", posCorrect)
	return
}

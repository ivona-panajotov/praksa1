package main

import (
	"example.com/kombinacije/check"
	"example.com/kombinacije/randomize"

	"errors"
	"fmt"
	"math/rand"
	// "strings"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	// test.RandomString(4)
	var generatedStr = random.RandomString(4)
	fmt.Println(generatedStr)
	var letters string
	fmt.Println("Type letter combination")
	fmt.Scanln(&letters)
	var shortStr = errors.New("Letter combination must be 4")
	if len(letters) < 4 {
		fmt.Println(shortStr)
	}
	err := check.Check(letters, generatedStr)
	if err != nil {
		fmt.Println(shortStr)
	}
	// fmt.Println("Letter combo is ", letters)
	// letters = strings.EqualFold(value, letters)
	check.Check(letters, generatedStr)

	// fmt.Println(letters == value)
}

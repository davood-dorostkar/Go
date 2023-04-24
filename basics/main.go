package main

import (
	"fmt"
)

var level = map[string]int{
	"a":  1,
	"b:": 2,
}

type Car struct {
	length int
	mass   float32
}

func check(in int) {
	if in < 20 {
		fmt.Print("low value")
	} else if in > 30 {
		fmt.Print("hight value")
	} else {
		fmt.Print("mid value")
	}
}

func loop(end int) {
	for i := 0; i <= end; i++ {
		fmt.Println(i)
	}
}

func loopArray() {
	var nums = []int{11, 22, 33}
	for i := range nums {
		fmt.Println(i)
	}
}

func loopArrayValues() {
	var nums = []int{11, 22, 33}
	for _, i := range nums {
		fmt.Println(i)
	}
}

func loopArrayIndexValues() {
	var nums = []int{11, 22, 33}
	for index, value := range nums {
		fmt.Printf("%v, %v\n", index, value)
	}
}

func while() {
	i := 0
	for i <= 6 {
		fmt.Println(i)
		i++
	}
}

func interrupt() {
	for {
		var input int
		fmt.Scanln(&input)
		if input == -1 {
			break
		} else {
			fmt.Printf("you entered: %v\n", input)
		}
	}
}

func switcher(input int) {
	// var input int
	// fmt.Scanln(&input)
	switch input {
	// switch input {
	case 1:
		fmt.Println("low value")
	case 2:
		fmt.Println("mid value")
	default:
		fmt.Println("high value")
	}
}

func say(message string) {
	fmt.Println(message)
}

func makeMessage(name string) string {
	message := "hello, " + name
	return message
}

func calc(a int, b int) (x float32, y int) {
	x = float32(a) / float32(b)
	y = a * b
	return x, y
}

func main() {
	// name := "world!"
	// count := 42
	// fmt.Printf("Hello, %s, %v\n", name, count)
	// var benz Car
	// benz.length = 4
	// check(25)
	// loop(5)
	// loopArray()
	// loopArrayValues()
	// loopArrayIndexValues()
	// while()
	// interrupt()
	// switcher(2)
	// say("hello")
	// message := makeMessage("david")
	// fmt.Println(message)
div, mult := calc(11, 34)
values := fmt.Sprintf("divide: %v, multiply: %v", div, mult)
fmt.Println(values)
fmt.Println(calc(11, 34))
}

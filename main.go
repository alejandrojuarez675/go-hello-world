package main

import (
	"fmt"
)

func main() {
	fmt.Println("Init example")

	fmt.Println("Hello world")

	fmt.Println("Creating dog")
	toby := dog{name: "toby"}
	fmt.Println(toby)
	toby.Bark()

	fmt.Println("Creating calculator")
	casioCalculator := calculator{branding: "casio", price: 15.2}
	numbersArray := make([]int16, 0, 15)
	for i := int16(0); i < 15; i++ {
		numbersArray = append(numbersArray, i)
	}
	fmt.Printf("Numbers to sum %v\n", numbersArray)
	res := casioCalculator.GetSummation(numbersArray...)
	fmt.Printf("The summation of the first 15 numbers is %d\n", res)
}

type dog struct {
	name string
}

func (d dog) Bark() {
	fmt.Printf("%s: is barking\n", d.name)
}

type calculator struct {
	branding string
	price    float32
}

func (calculator) GetSummation(numbers ...int16) (result int32) {
	for _, v := range numbers {
		result += int32(v)
	}
	return
}

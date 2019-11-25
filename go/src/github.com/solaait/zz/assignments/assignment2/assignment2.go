package main

import "fmt"

func main() {

	Fibonacci()

	Combine()
}

func Fibonacci() {

	var nthFibonacci int

	fmt.Println("enter the number:")

	fmt.Scanln(&nthFibonacci)
	a := 0
	b := 1
	fmt.Printf("%d %d ", a, b)
	for i := 2; i < nthFibonacci; i++ {

		c := a + b

		a = b

		b = c

		fmt.Printf("%d ", b)

	}

	fmt.Println()

}

func Combine() {

	charArray := []string{}
	numArray := []int{}

	var charsLength int

	fmt.Println("enter the number characters:")

	fmt.Scanln(&charsLength)

	var numsLength int

	fmt.Println("enter the number integers:")

	fmt.Scanln(&numsLength)

	var chars string

	fmt.Println("enter the characters:")
	for i := 0; i < charsLength; i++ {

		fmt.Scanln(&chars)

		charArray = append(charArray, chars)
	}

	var numbers int
	fmt.Println("enter the numbers:")

	for i := 0; i < numsLength; i++ {

		fmt.Scanln(&numbers)

		numArray = append(numArray, numbers)
	}

	fmt.Printf("%s,%d", charArray[0], numArray[0])

	for value := 1; value < charsLength; value++ {

		fmt.Printf(",%s,%d", charArray[value], numArray[value])
	}
}

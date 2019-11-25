package main

import (
	"fmt"
)

var glob string ="solomon"

func main() {

	// x := make([]float32, 5)
	// x[0] = 1.2
	// x[1] = 3.0
	// x[2] = 2

	// for index,element := range x {
		
	// 	fmt.Println(index," :",element)
	// }
    fmt.Println(glob)
	ch := make(chan string)

	ch2 := make(chan string)

	ch3 := make(chan string)

	go greet(ch)

	ch <- "solomon"

	go name(ch2)
	ch2 <- "Solomon"

	ch2 <- "Kindie"

	go sendFname(ch3)

	<-ch3

	test1 := Rectangle{2.5, 4.5}

	fmt.Println(test1)

	test1.scaleOne()

	fmt.Println(test1)

}

type Rectangle struct {
	height float32
	width  float32
}

func (r *Rectangle) scaleOne() {

	r.height *= 2
	r.width *= 2

}

func greet(ch chan string) {

	message := "hello " + <-ch + " !"

	fmt.Println(message)

}

func name(c chan string) {

	output := "firstname" + ": " + <-c + " lastname" + ": " + <-c

	fmt.Println(output)
}

func sendFname(c chan string) {

	c <- "solomon"

}

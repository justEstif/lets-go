package main

import "fmt"

const Pi = 3.14

func typeInterface() {
	const World = "FD"
	fmt.Println("Hello,", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

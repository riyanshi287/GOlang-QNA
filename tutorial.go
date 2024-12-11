package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to my quiz game!")

	score := 0

	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome to the game!\n", name)
	
	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age>= 10{
		fmt.Println("Yay you can play!")
	} else {
		fmt.Println("oh no! you cannot play...try again in some years")
		return
	}

	fmt.Printf("What is better for gaming: SSD or HDD? ")
	var ans string
	fmt.Scan(&ans)

	upperstring := strings.ToUpper(ans)
	if upperstring == "SSD"{
		fmt.Println("CORRECT!")
		score++
	} else{
		fmt.Println("INCORRECT!")
	}

	fmt.Printf("which Latency is lower: 50 ms or 10 ms? ")
	var ans1, ans2 string
	fmt.Scan(&ans1, &ans2)
	lowerString := strings.ToLower(ans1+" "+ ans2)
	if lowerString == "10 ms"{
		fmt.Println("CORRECT!")
		score++
	} else{
		fmt.Println("INCORRECT!")
	}
	no_of_ques := 2
	result := (float64(score) / float64(no_of_ques))* 100
	fmt.Printf("you have score %v%%.", result)

}
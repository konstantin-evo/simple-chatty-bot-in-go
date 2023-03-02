package main

import (
	"fmt"
)

func greet(name string) {
	year := 2022
	fmt.Printf("Hello! My name is %s.\n", name)
	fmt.Printf("I was created in %d.\n", year)
}

func showName() {
	var name string
	fmt.Println("Please, remind me your name.")
	fmt.Scan(&name)
	fmt.Println("What a great name you have, " + name + "!")
}

func guessAge() {
	var rem3, rem5, rem7, age int

	fmt.Println("Let me guess your age.")
	fmt.Println("Enter remainders of dividing your age by 3, 5 and 7.")
	fmt.Scan(&rem3, &rem5, &rem7)

	age = (rem3*70 + rem5*21 + rem7*15) % 105
	fmt.Printf("Your age is %d; that's a good time to start programming!\n", age)
}

func count() {
	var n int

	fmt.Println("Now I will prove to you that I can count to any number you want.")
	fmt.Scan(&n)
	for i := 0; i <= n; i++ {
		fmt.Printf("%d!\n", i)
	}
}

func startQuiz() {
	quizQuestion := "Why do we use methods?"
	quizOptions := []string{"To repeat a statement multiple times.", "To decompose a program into several small subroutines.", "To determine the execution time of a program.", "To interrupt the execution of a program."}
	quizAnswer := 2
	for {
		fmt.Println("Let's test your programming knowledge.")
		fmt.Println(quizQuestion)
		for i, option := range quizOptions {
			fmt.Printf("%d. %s\n", i+1, option)
		}
		var answer int
		fmt.Scan(&answer)
		if answer == quizAnswer {
			fmt.Println("Congratulations, you got it right!")
			break
		} else {
			fmt.Println("Sorry, that's incorrect. Please try again.")
		}
	}
}

func sayGoodbye() {
	fmt.Println("Congratulations, have a nice day!")
}

func main() {
	greet("Aid")
	showName()
	guessAge()
	count()
	startQuiz()
	sayGoodbye()
}

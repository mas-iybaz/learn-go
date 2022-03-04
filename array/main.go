package main

import (
	"fmt"
)

func averageGrade(arr [3]int) float32 {
	var avg float32

	for _, score := range arr {
		score := float32(score)
		avg += score
	}

	return avg / float32(len(arr))
}

func main() {
	// Array Declaration
	var studentGrades [3]int
	var students []string
	randomNumbers := [3]int{3, 6, 9}
	drinks := [2][2]string{
		{"tea", "coffe"},
		{"milk", "coffe"},
	}

	studentGrades[0] = 80
	studentGrades[1] = 85
	studentGrades[2] = 95

	students = []string{"Alpha", "Beta", "Charlie"}

	for index, grade := range studentGrades {
		fmt.Printf("Grade for student %d: %d \n", index+1, grade)
	}

	for index, name := range students {
		fmt.Printf("Name for student %d: %v \n", index+1, name)
	}

	for _, number := range randomNumbers {
		fmt.Println("random number:", number)
	}

	for index, userDrinks := range drinks {
		fmt.Println("User", index+1, "favorite drink is:")

		for _, drink := range userDrinks {
			fmt.Printf("%v ", drink)
		}

		fmt.Println()
	}

	fmt.Println(averageGrade(studentGrades))
}

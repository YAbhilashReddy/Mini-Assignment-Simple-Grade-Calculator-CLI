package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	inputReader := bufio.NewReader(os.Stdin)

	//01 : Taking Student Name :=
	fmt.Println("Enter Your name :- ")
	studentName, _ := inputReader.ReadString('\n')
	studentName = strings.TrimSpace(studentName)

	//02 : Taking Student Marks :=
	fmt.Println("Enter your Total Marks :- ")
	studentMarks, _ := inputReader.ReadString('\n')
	studentMarks = strings.TrimSpace(studentMarks)

	marks, err := strconv.Atoi(studentMarks)

	if err != nil {
		fmt.Println("Invalid Marks? Pls Enter Correct Positive Number:-", err)
		return
	}

	grade := calculatingStudentGrading(marks)
	fmt.Println(studentName, "scored", marks, "marks and received grade:-", grade)
}

func calculatingStudentGrading(marks int) string {
	if marks >= 90  || marks <= 100{
		return "A"
	} else if marks >= 80 {
		return "B"
	} else if marks >= 70 {
		return "C"
	} else if marks >= 60 {
		return "E"
	} else {
		return "F"
	}

}

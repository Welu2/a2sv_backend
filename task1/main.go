package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Subject represents a single subject entry
type Subject struct {
	Name  string
	Score float64
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	
	fmt.Print("Enter student name: ")
	studentName, _ := reader.ReadString('\n')
	studentName = strings.TrimSpace(studentName)

	fmt.Print("Enter number of subjects: ")
	numInput, _ := reader.ReadString('\n')
	numSubjects, _ := strconv.Atoi(strings.TrimSpace(numInput))

	// Store Data in a Slice
	var subjects []Subject

	for i := 1; i <= numSubjects; i++ {
		fmt.Printf("Enter subject %d name: ", i)
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		fmt.Printf("Enter subject %d score: ", i)
		scoreStr, _ := reader.ReadString('\n')
		score, _ := strconv.ParseFloat(strings.TrimSpace(scoreStr), 64)

		subjects = append(subjects, Subject{Name: name, Score: score})
	}

	// Calculate and Display Report
	total, average := calculateScores(subjects)
	displayReport(studentName, subjects, total, average)
}

// calculateScores computes the sum and mean of the subject scores
func calculateScores(subjects []Subject) (float64, float64) {
	total := 0.0
	for _, s := range subjects {
		total += s.Score
	}
	average := total / float64(len(subjects))
	return total, average
}

// displayReport prints the formatted grade report to the terminal
func displayReport(name string, subjects []Subject, total float64, average float64) {
	fmt.Println("\n--- Grade Report ---")
	fmt.Printf("Student Name: %s\n", name)
	fmt.Printf("%-15s %-10s\n", "Subject", "Score")
	fmt.Println(strings.Repeat("-", 25))

	for _, s := range subjects {
		fmt.Printf("%-15s %-10.2f\n", s.Name, s.Score)
	}

	fmt.Println(strings.Repeat("-", 25))
	fmt.Printf("Total Score:   %.2f\n", total)
	fmt.Printf("Average Score: %.2f\n", average)
}

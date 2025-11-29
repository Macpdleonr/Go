package main

import {
	"os"
	"fmt"
	"encoding/csv"
	"strconv"
}



var functions = map[string] func(int, int) int{

	"sum": func(a int, b int) int { return a + b },
	
	"rest": func(a int, b int) int { return a - b },

	"multiply": func(a int, b int) int { return a * b },

	"divide": func(a int, b int) int { return a / b },
}

func presentResult(operation string, a int, b int){

	f, exists := functions[operation]

	if !exists {
		fmt.Println("Operation ", operation, " not found")
		return
	}
	
	fmt.Println("For a = ", a, " and b = ", b, " the ", operation," result is : ", f(a,b))
}

// Student struct for store your information
type Student struct{
	Name string
	Notes []float64
}

func readCSV(nameFile string)([]Student, error) {
	
	file, err := os.Open(nameFile)

	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)

	reader.Comma = ';'

	records, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	var students []Student

	for _, record := range records {
		name := record[0]

		var notes []float64

		for_, noteStr := range record[1:] {	
			note, err := strconv.ParseFloat(noteStr, 64)
			if err != nil {
				return nil, err
			}
			notes = append(notes, note)
		}
		students = append(students, Student{Name: name, Notes: notes})
	}

	return students, nil
}

// Function to calculate average of notes
func calculateAverage(notes []float64) float64 {

	var sum float64 = 0

	for _, note := range notes {
		sum += note
	}

	return sum / float64(len(notes))
}

// Function to calculate arithmetic mean of all students
func calculateArithmeticMean(students []Student) float64 {
	
	var sum float64
	var countNotes int

	for _, student := range students {

		for _, note := range student.Notes {

			sum += note
			countNotes++

		}

	}

	return sum / float64(countNotes)
}

// Function to write CSV with students and their averages
func writeCSV(nameFile string, students []Student) error {
	
	file, err := os.Create(nameFile)

	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	for _, student := range students {

		average := calculateAverage(student.Notes)

		record := []string{student.Name, fmt.Sprintf("%.2f", average)}

		if err := writer.Write(record); err != nil {
			return err
		}

	}

	return nil
}

func main() {

	presentResult("sum",3, 4)
	presentResult("sum",10, 20)
	presentResult("sum",-55, 83)

	presentResult("rest",5, 2)
	presentResult("rest",100, 50)
	presentResult("rest",80, 25)

	presentResult("multiply",3, 4)
	presentResult("multiply",10, 20)
	presentResult("multiply",-5, 6)

	presentResult("divide",8, 2)
	presentResult("divide",100, 5)
	presentResult("divide",81, 9)

	// Reading CSV file
	students, err := readCSV("notes.csv")

	if err != nil {
		fmt.Println("Error reading CSV file: ", err)
		return
	}

	// Calculating arithmetic mean of all students
	arithMean := calculateArithmeticMean(students)

	fmt.Printf("The arithmetic mean of all students is: %.2f\n", arithMean)

	// Writing CSV file with students and their averages
	err = writeCSV("students_averages.csv", students)

	if err != nil {
		fmt.Println("Error writing CSV file: ", err)
		return
	}

	fmt.Println("students_averages.csv file created successfully.")

}
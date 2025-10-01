# ESEP Grade Calculator

This calculates a student's final grade based on assignments, exams, and essays by calculating a weighted average and converting it into a letter grade.

- Assignments: 50%  
- Exams: 35%  
- Essays: 15%  

Grade scale:  
- A: 90–100  
- B: 80–89  
- C: 70–79  
- D: 60–69  
- F: <60  


## Setup

1. Fork this repository to your own GitHub account.  

2. Clone your fork using the following commands:
- git clone https://github.com/<your-username>/esep-grade-calculator.git
- cd esep-grade-calculator

3. Check your Go installation to ensure it is at least version 1.24
- go version


## Running Tests

Run all unit tests:
- cd grade-calculator
- go test ./...

Run unit tests with coverage (what % of code is tested):
- cd grade-calculator
- go test -cover ./...” 


## Project Files

go.mod – Module name and Go version

grade_calculator.go – Implementation of the Grade Calculator

grade_calculator_test.go – Unit tests

README.md – Documentation (this file)
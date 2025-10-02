# ESEP Grade Calculator

A tiny Go library for computing final course grades from assignments, exams, and essays. It averages each category, applies weights, then converts the final numeric score to a letter grade.

Runtime requirements:
Golang 1.24+

Run unit tests:
```
cd grade-calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade-calculator/
go test -coverprofile ".\coverage.txt" .
go tool cover -func ".\coverage.txt"
# Optional HTML view:
go tool cover -html ".\coverage.txt"
```
# Go Program: Handling Non-numeric Input in Number Check

This program checks if a given number is even or odd.  However, it has a flaw: it doesn't handle cases where the user enters non-numeric input. This leads to a runtime panic.

The `bug.go` file contains the buggy code.  The solution, which gracefully handles non-numeric input, is in `bugSolution.go`.

## How to run:

1. Save the code in `bug.go` and `bugSolution.go`.
2. Compile and run using `go run bug.go` or `go run bugSolution.go`.
3. Test with both numeric and non-numeric input to observe the difference.
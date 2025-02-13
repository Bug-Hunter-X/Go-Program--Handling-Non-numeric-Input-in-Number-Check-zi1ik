func main() {
  fmt.Println("Enter a number:")
  var input string
  fmt.Scanln(&input)

  num, err := strconv.Atoi(input)
  if err != nil {
    fmt.Println("Invalid input. Please enter a valid integer.")
    return
  }

  if num%2 == 0 {
    fmt.Println(num, "is even")
  } else {
    fmt.Println(num, "is odd")
  }
} 
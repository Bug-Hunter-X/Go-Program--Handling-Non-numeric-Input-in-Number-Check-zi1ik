func main() {



  fmt.Println("Enter a number:")
  var num int
  fmt.Scanln(&num)

  if num%2 == 0 {
    fmt.Println(num, "is even")
  } else {
    fmt.Println(num, "is odd")
  }

}
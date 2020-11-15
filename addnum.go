package main
import "fmt"
  
func main() { 
var num1 int
var num2 int
var sum int
  fmt.Println("Enter first number:")
  fmt.Scanln(&num1)
  fmt.Println("Enter second number:")
  fmt.Scanln(&num2)
  sum=num1+num2
  fmt.Println("Sum is:", sum)
}

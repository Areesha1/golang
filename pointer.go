package main 
import (
    "fmt"
    "os" )
 func main () {

   x := 1
   p := &x    // p, of type *int, points to x
  fmt.Println(*p)    // "1"
  
  *p = 6     // equivalent to x = 6
  fmt.Println(x) // "6"
  fmt.Println(p, *p)
}

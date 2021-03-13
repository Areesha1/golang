package main
import "fmt"
func main(){
  arr1 := [...]int{4, 2, 5,1}
  fmt.Println(arr1[2])
  fmt.Println(len(arr1))
  for i, v := range arr1 {
  fmt.Printf("%d %d\n", i, v)
  }
}

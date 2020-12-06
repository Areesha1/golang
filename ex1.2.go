package main
import (
  "fmt"
  "os"
  )
 func main (){
  
   for n , arg:= range os.Args[1:]{
   
   fmt.Println ("The index" , n , "contains the value" , arg )
   }
}
   

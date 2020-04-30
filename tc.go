package main

import "fmt"


func main(){

      fmt.Printf("Enter a floating point number:")
      var x float32
      var y int
      fmt.Scan(&x)
      y=int(x)
      fmt.Printf("\nThe truncated value is:")
      fmt.Printf("%d",y)
}
package main

import "fmt"

func main(){

  var a,b,time int
  
  fmt.Scan(&time)
  
    a = time/30
  
    b = (time%30)*2
  
    fmt.Println("It is", a, "hours", (b), "minutes.")
}

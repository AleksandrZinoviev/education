package main
import "fmt"
func main(){
  var a,b int
  fmt.Scan(&a) 
  fmt.Scan(&b)
  a = a * a
  b = b * b
  fmt.Println(a+b)
}

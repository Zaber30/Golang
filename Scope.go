package main
import "fmt"

var a=20
var b=30
func add(x int, y int){
	z :=x+y
	fmt.Println(z)
}
 func main(){
	fmt.Println("hello world")
	p :=30
	q :=40
	add(p,q)
	add(a,b)
	

 }
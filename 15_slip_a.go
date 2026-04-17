package main
import "fmt"

func divide(n1, n2 int) (int, int){
    return n1/n2, n1%n2
}

func main(){
    n1, n2 := 10, 3
    q, r := divide(n1, n2)
    fmt.Println("Q: ", q, "\tR: ", r)
}
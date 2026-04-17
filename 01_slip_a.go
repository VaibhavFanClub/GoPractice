package main
import "fmt"

func main(){
    var n1, n2, c int
    fmt.Print("Enter two numbers: ")
    fmt.Scan(&n1, &n2)
    fmt.Print("Enter choice (1. +, 2. -): ")
    fmt.Scan(&c)
    switch c {
        case 1: fmt.Println(n1 + n2)
        case 2: fmt.Println(n1 - n2)
    }
}
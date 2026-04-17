package main
import "fmt"

func main(){
    es, os := 0, 0
    for i := 1; i <= 100; i++ {
        if i % 2 == 0 {
            es += i
        }else {
            os += i
        }
    }
    fmt.Println("Sum of even no.: ", es)
    fmt.Println("Sum of odd no.: ", os)
}
package main
import "fmt"

func swap(n1, n2 *int){
    temp := *n1
    *n1 = *n2
    *n2 = temp
}

func main(){
    n1, n2 := 2, 6
    fmt.Println("Original: ", n1, n2)
    swap(&n1, &n2)
    fmt.Println("swap: ", n1, n2)
}
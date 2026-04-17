package main
import "fmt"

func table(n int) {
    for i:=1; i<11; i++{
        fmt.Println(n, " x ", i, " = ", n*i)
    }
}

func main(){
    table(5)
}
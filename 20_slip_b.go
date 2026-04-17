package main
import "fmt"

func ex(n int, ch chan int){
    for i:=1; i <= n; i++{
        ch <- i
    }
    close(ch)
}

func main(){
    ch := make(chan int)
    go ex(10, ch)
    for val := range ch {
        fmt.Print(val, " ")
    }
}
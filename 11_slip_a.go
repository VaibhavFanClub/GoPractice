package main
import "fmt"

func main(){
    var n int
    fmt.Print("Enter a number: ")
    fmt.Scan(&n)
    if n >= -99 && n <= -10 || n >= 10 && n <= 99 {
        fmt.Print("Number is two digit")
    }else {
        fmt.Print("Number is not two digit")
    }
}
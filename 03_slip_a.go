package main
import "fmt"

func isPalindrome(n int){
    temp := n
    rev := 0
    for n != 0{
        d := n % 10
        rev = rev * 10 + d
        n /= 10
    }
    if temp == rev {
        fmt.Println("Palindrome")
    }else {
        fmt.Println("Not Palindrome")
    }
}

func main(){
    var n int
    fmt.Print("Enter a number: ")
    fmt.Scan(&n)
    isPalindrome(n)
}
package main
import "fmt"

func fib(ch chan int){
    n := <-ch
    a, b := 0, 1
    for i:=0; i<n; i++{
        fmt.Print(a, " ")
        a, b = b, a+b
    }
}

func main(){
    ch := make(chan int)
    go fib(ch)
    ch <- 10
    close(ch)
}

/*
package main

import "fmt"

func fib(n int, ch chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}
	close(ch)
}

func main() {
	ch := make(chan int)

	go fib(10, ch)

	for val := range ch {
		fmt.Print(val, " ")
	}
}
*/
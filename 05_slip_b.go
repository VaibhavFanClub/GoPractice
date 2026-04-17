package main
import "fmt"

type Employee struct {
    eno int
    ename string
    sal float64
}

func main(){
    var n int
    fmt.Print("Enter no of emps: ")
    fmt.Scan(&n)
    e := make([]Employee, n)
    min := 10000000000000000.0
    
    fmt.Println("\nEnter records of emps: ")
    for i:=0; i<n; i++{
        fmt.Print("eno: ")
        fmt.Scan(&e[i].eno)
        fmt.Print("name: ")
        fmt.Scan(&e[i].ename)
        fmt.Print("sal: ")
        fmt.Scan(&e[i].sal)
        if e[i].sal < min {
            min = e[i].sal
        }
    }
    
    fmt.Println("\nEmp with min sal: ")
    for i:=0; i<n; i++{
        if e[i].sal == min{
            fmt.Println("eno: ", e[i].eno)
            fmt.Println("name: ", e[i].ename)
            fmt.Println("sal: ", e[i].sal)
        }
    }
}
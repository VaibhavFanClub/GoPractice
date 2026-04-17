package main
import "fmt"

func main(){
    arr := [6]int{5, 3, 7, 4, 2, 9}
    fmt.Println("Orignal: ", arr)
    for p:=1; p<len(arr); p++{
        for i:=0; i<len(arr) - p; i++{
            if arr[i] > arr[i+1]{
                temp := arr[i]
                arr[i] = arr[i+1]
                arr[i+1] = temp
            }
        }
    }
    fmt.Print("Sorted: ", arr)
}
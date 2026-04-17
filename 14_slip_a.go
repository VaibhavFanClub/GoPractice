package main
import "fmt"

func main(){
    s := []int{1, 2, 3, 4}
    fmt.Println("Original: ", s)
    s = append(s, 5, 6)
    fmt.Println("After append: ", s)
    s = append(s[:2], s[3:]...)                 // *
    fmt.Println("After removing 3: ", s)
    newS := make([]int, len(s))                 // *
    copy(newS, s)                               // *
    fmt.Println("After copying: ", newS)
}
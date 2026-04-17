package main
import "fmt"

type Student struct {
    roll int
    name string
}

func (s *Student) show(){
    fmt.Println("roll: ", s.roll, "\nname: ", s.name)
}

func main(){
    s := Student{1, "Raj"}
    s.show()
}
package main
import "fmt"

type Arr [5]int

func (n Arr) copy() Arr{
    /* you can do this too or can return n directly
    var copied Arr
	for i := 0; i < len(n); i++ {
		copied[i] = n[i]
	}
	return copied
    */
    return n
}

func main(){
    n := Arr{1, 2, 3, 4, 5}
    new := n.copy()
    fmt.Print(new)
}
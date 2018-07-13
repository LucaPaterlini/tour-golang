package src

import (
	"golang.org/x/tour/tree"
	"fmt"
)

func _walk(t *tree.Tree, ch chan int){
	if t==nil {return}
	_walk(t.Left, ch)
	ch <- t.Value
	_walk(t.Right,ch)
}


func Walk(t *tree.Tree, ch chan int){
	_walk(t,ch)
	close(ch)
}

func Same(t1, t2 *tree.Tree)bool{
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1,ch1)
	go Walk(t2,ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if v1 != v2 || ok1 != ok2{ return false}
		if !ok1{break}
	}
	return true
}



func main()  {
	ch := make(chan int)
	go Walk(tree.New(1),ch)
	for i:= range ch {
		fmt.Print(i," ")
	}
	fmt.Println(Same(tree.New(1),tree.New(1)))
	fmt.Println(Same(tree.New(1),tree.New(2)))

}

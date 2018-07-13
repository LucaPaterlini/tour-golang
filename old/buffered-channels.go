package old

import "fmt"

func main(){
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(len(ch)==cap(ch))
}

package old

import "fmt"

func main(){
	var i interface{}
	describe(i)

	i = &i
	describe(i)

	i="hello"
	describe(i)

}

func describe(i interface{}){
	fmt.Printf("(%v, %T)\n",i,i)
}


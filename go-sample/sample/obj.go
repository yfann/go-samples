package sample

import "fmt"

type obj struct{
	a string
	b string
	c string
}

func change(o *obj){
	o.a="change:o.a"
}

func TestObj()  {
	o:=obj{
		a:"123",
		b:"123",
	}

	change(&o)

	fmt.Println(o)
}
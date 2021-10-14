package sample

import (
	"fmt"
	"time"
)

func TestTime()  {
	//t:=time.Now()
	//
	//s:=fmt.Sprintf("%d-",t.Year())
	//
	//fmt.Println(s)

	var b time.Time
	fmt.Println(b)
	fmt.Println(b.IsZero())
	b=time.Now()
	fmt.Println(b.IsZero())
	fmt.Println(b)
	var i int32=10

	c:=b.Add(time.Minute*time.Duration(i))
	fmt.Println(c)
	//fmt.Println(time.Now().Format("2006-01-02-15-04-05"))
	//fmt.Println(time.Now().Format("20060102150405"))

	fmt.Printf("%s before %s is %s\n", time.Now(),c,time.Now().Before(c))
	fmt.Println(time.Now().After(c))
}

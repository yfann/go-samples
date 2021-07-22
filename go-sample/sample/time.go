package sample

import (
	"fmt"
	"time"
)

func TestTime()  {
	t:=time.Now()

	s:=fmt.Sprintf("%d-",t.Year())

	fmt.Println(s)


	fmt.Println(time.Now().Format("2006-01-02-15-04-05"))
	fmt.Println(time.Now().Format("20060102150405"))
}

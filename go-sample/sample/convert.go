package sample

import (
	"fmt"
	"strconv"
)

func TestConvert(){
	i,_:=strconv.ParseInt("123",10,32)
	fmt.Println(i)
}

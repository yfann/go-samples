package sample

import (
	"fmt"
	"strings"
)

func  TestStr()  {
	fmt.Println(strings.TrimSpace(" s sdfsdfdd ssd  df    "))
	fmt.Println(strings.ReplaceAll(" s sdfsdfdd ssd  df"," ","_"))
	fmt.Println(fmt.Sprintf("xxx %s","123"))
}
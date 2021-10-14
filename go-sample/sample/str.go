package sample

import (
	"fmt"
	"strings"
)

func  TestStr()  {
	fmt.Println(strings.TrimSpace(" s sdfsdfdd ssd  df    "))
	fmt.Println(strings.ReplaceAll(" s sdfsdfdd ssd  df"," ","_"))
	fmt.Println(fmt.Sprintf("xxx %s","123"))

	m := make(map[string]string)
	m["aa"]="bb"
	m["aa1"]="bb1"
	fmt.Println("map:", m)
}
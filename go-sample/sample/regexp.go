package sample

import (
	"fmt"
	"regexp"
)

func TestReg()  {
	r, _ := regexp.Compile("\\$(\\w+)")
	fmt.Println(r.FindString("$namespace| SELECT  namespace , container_name as name,cpu,memory, count(1) as count group by name, cpu, memory, namespace"))
	l:=r.FindStringSubmatch("$namespace| SELECT  namespace , container_name as name,cpu,memory, count(1) as count group by name, cpu, memory, namespace")
	fmt.Println(l[0],l[1])
}

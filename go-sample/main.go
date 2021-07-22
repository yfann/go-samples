package main

import (
	"fmt"
	"go-sample/vault"
	"strconv"
	"strings"
)

func main() {
	//sample.CreateFolder()

	//sample.TestObj()

	//sample.TestTime()

	//sample.TestReg()

	//testMap()

	//sample.CreateFile()

	//testSlice()
	//testStr()
	//testVault()
	//delVault()
	//sample.TestStr()
	//fmt.Println(vault.ParseToken(`
	//	Unseal Key 1: .....
	//	Unseal Key 2: .....
	//	Unseal Key 3: .....
	//	Unseal Key 4: .....
	//	Unseal Key 5: .....
	//
	//	Initial Root Token:.....
	//`))
	//sample.TestHash()
	//vault.TestK8sAuth()
	//var x interface{} = "abc"
	//fmt.Println(fmt.Sprintf("%v", x))
	//sample.ReadFile()
	//fmt.Println(len("ry.cn-shanghai.aliyuncs.com/hsc/hsc-vault-operator-bundle:1.1.2"))
	vault.AddTransitKey("test")
}

func testVault() {
	c := vault.GetVaultClient()

	sec, err := c.Logical().Write("/kv/test/vault/", map[string]interface{}{
		"name": "yfann",
		"age":  "100",
		"pwd":  "asdf9s0fd90sf9s0f",
	})

	fmt.Println(sec)
	fmt.Println(err)
}

func delVault() {
	c := vault.GetVaultClient()

	c.Logical().Delete("/kv/test/vault")
}

func testStr() {
	//fmt.Println(strings.Replace("$sdfg sdf as $sdfg","$sdfg","@@@@",-1))
	a := "namespace:<$project>"
	fmt.Println(strings.Replace(a, "<$project>", "123", 1))
	fmt.Println(a)

	b := "asdfasdf| 12332"
	fmt.Println(strings.Split(b, "|"))

	fmt.Println(strings.HasPrefix("my string", "my  s"))
	fmt.Println(strings.HasPrefix("my string", "xxx"))

	fmt.Println(strconv.Quote("Hello, 世界"))
	fmt.Println("\"Hello, \"世界\"")
	fmt.Println(strconv.Unquote("\"Hello, \"世界\""))
	fmt.Println(strconv.Quote("Hello, 世界"))
	fmt.Println(strconv.QuoteToASCII("Hello, 世界"))
}

func testMap() {
	var a map[string]string
	a = make(map[string]string)
	a["a"] = "a"
	fmt.Println(a)
}

func testSlice() {
	var s []string
	var c []string

	s = append(s, "test")
	fmt.Println(c == nil)

	s = append(s, c...)
	fmt.Println(s)

	fmt.Println(strings.Join([]string{"sdfsdf", "sdfsdf", "apple"}, ","))
}

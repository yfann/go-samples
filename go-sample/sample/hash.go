package sample



import (
	"fmt"
	"hash/fnv"
)

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func TestHash() {
	fmt.Println(hash("HelloWorld"))
	fmt.Println(hash("HelloWorld "))
	fmt.Println(hash("HelloWorld  "))
	fmt.Println(hash("HelloWorld "))
	fmt.Println(hash("HelloWorld                                     "))
}
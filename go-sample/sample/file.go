package sample

import (
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
)

func CreateFolder(){
	path:="C:\\tmp\\tt\\123\\456/7889"
	if _,err:=os.Stat(path);os.IsNotExist(err){
		err := os.Mkdir(path, 0755)
		fmt.Println("err:",err)
	}
}


func CreateFile(){
	_, err := os.Create("C:\\tmp\\tt\\a.txt")
	if err!=nil{
		fmt.Println("err:",err)
	}
}


func ReadFile(){
	dat, _ := ioutil.ReadFile("C:\\tmp\\aa.txt")
	fmt.Print(string(dat))
}


func GetEnv() (map[string]string, error) {
	dat, _ := ioutil.ReadFile("C:\\tmp\\aa.txt")

	m := make(map[string]string)
	if err := json.Unmarshal(dat, &m); err != nil {
		return nil, err
	}

	return nil, nil
}
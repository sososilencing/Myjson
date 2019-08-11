package main

import (
	"Myjson/model"
	"encoding/json"
	"fmt"
	"unsafe"
)

func main() {
	//var card string
	//card="{\"Number\":\"12312\",\"Name\":\"what\"}"
	//b:=*(*[]byte)(unsafe.Pointer(&card))
	//ca:=model.Card{}
	//json.Unmarshal(b,&ca)
	//fmt.Println(ca.Name)
	//fmt.Println(ca)
	//fmt.Println(ca.Number)
	student := &model.Student{}
	s, _ := json.Marshal(student)
	fmt.Println(*(*string)(unsafe.Pointer(&s)))
}

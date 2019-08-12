package main

import (
	"Myjson/method"
	"Myjson/model"
	"encoding/json"
	"fmt"
	"unsafe"
)

func main() {
	var card string
	card = "{\"Number\":\"12312\",\"Name\":\"你死了\"}"
	b := *(*[]byte)(unsafe.Pointer(&card))

	fmt.Println(b)

	var ca model.Card

	method.Unmarshal(card, &ca)
	//if err := json.Unmarshal(b, &ca);err!=nil{
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(ca.Number)
	//fmt.Println(ca.Name)
	//fmt.Println(ca)

	var a [3]int
	a[0] = 1
	a[1] = 3
	a[2] = 4
	s, _ := json.Marshal(a)
	fmt.Println(*(*string)(unsafe.Pointer(&s)))

	var name [2]int
	json.Unmarshal(s, &name)
	fmt.Println(name)
}

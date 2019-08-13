package main

import (
	"MyJson/method"
	"MyJson/model"
	"fmt"
)

func main() {
	student := &model.Student{
		StuCode:  1234,
		StuName:  "hgn",
		CardName: model.Card{
			Number: "123",
			Name:   "321",
		},
		Thing:    [2]string{},
	}
	b:=method.Marshel(student)
	//b:=method.Marshel(&ca)
	// b:= *(*[]byte)(unsafe.Pointer(&card))
	//fmt.Println(*(*string)(unsafe.Pointer(&b)))
	ca :=&model.Student{}
	method.Unmarshal(string(b), &ca)
	fmt.Println(ca)
	//json.Unmarshal(b,&ca)
	//if err := json.Unmarshal(b, &ca);err!=nil{
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(ca)
	//fmt.Println(ca.Name)
	//fmt.Println(ca)

	//var a [3]int
	//a[0] = 1
	//a[1] = 3
	//a[2] = 4
	//s, _ := json.Marshal(a)
	//fmt.Println(*(*string)(unsafe.Pointer(&s)))
	//
	//var name [2]int
	//json.Unmarshal(s, &name)
	//fmt.Println(name)
}

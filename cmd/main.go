package main

import (
	"Myjson/model"
	"encoding/json"
	"fmt"
	"unsafe"
)

func main() {

	//student := &model.Student{
	//	StuCode: 2018212152,
	//	StuName: "hgn",
	//
	//	CardName: model.Card{
	//		Number: "16534",
	//		Name:   "card",
	//	},
	//}
	//
	//e := &student
	//b, _ := json.Marshal(&e)
	//
	//str := *(*string)(unsafe.Pointer(&b))
	//fmt.Println(str)
	//
	//m:=make(map[string][]int)
	//t:=make([]int,10)
	//t[0]=1
	//t[1]=3
	//m["1"]=t
	//fmt.Println(m)
	//s,_:=json.Marshal(&m)
	//str:=*(*string)(unsafe.Pointer(&s))
	//fmt.Println(str)
	//var a [10]bool
	//a[0]=true
	//a[1]=false

	a := make(map[string]map[string]string)
	q := make(map[string]string)
	q["xx"]="aa"
	a["1"]=q
	//var d interface{}
	student:=model.Student{}
	s,_:=json.Marshal(student)
	str:=*(*string)(unsafe.Pointer(&s))
	fmt.Println(str)
}

package main

import (
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

	q := make(map[string]map[string][]int)
	a := make(map[string]map[string]map[string][]int)
	g := make(map[string][]int)
	z := make([]int, 4)
	z[0] = 1
	z[1] = 2
	g["123"] = z
	q["xx"] = g
	a["1"] = q
	//var d interface{}
	//student:=model.Student{}
	dd := make([]int, 10)
	dd[1] = 1
	s, _ := json.Marshal(a)
	str := *(*string)(unsafe.Pointer(&s))
	fmt.Println(str)
}

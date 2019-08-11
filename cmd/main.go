package main

import "encoding/json"

func main(){

	//student:=&model.Student{
	//	StuCode:  2018212152,
	//	StuName:  "hgn",
	//	Thing: []string{
	//		"123",
	//		"23",
	//		"13",
	//	},
	//	CardName: model.Card{
	//		Number: "16534",
	//		Name:"card",
	//	},
	//}
	//t:=reflect.TypeOf(student)
	//e:=&student
	//json.Marshal(&e)
	//t=reflect.TypeOf(&e)
	//s:=t.Elem().Elem().Elem().Kind()
	//t=reflect.TypeOf(t)
	//t=reflect.TypeOf(t)
	//fmt.Println(s)
	//str:=*(*string)(unsafe.Pointer(&b))
	//fmt.Println(str)
	//
	//m:=make(map[string][]int)
	//t:=make([]int,10)
	//t[0]=1
	//t[1]=3
	//m["1"]=t
	//fmt.Println(m)
	//s,_:=json.Marshal(&m)
	//str=*(*string)(unsafe.Pointer(&s))
	//fmt.Println(str)
	var a [10]bool
	a[0]=true
	a[1]=false
	json.Marshal(a)
}

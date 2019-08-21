package main

import (
	"Myjson/quse"
	"Myjson/model"
	"fmt"
	_"time"
	"unsafe"
)

func main() {
	//t1 := time.Now()
	student:=&model.Student{
		StuCode: 231,
		Si: model.Card{
			Number: "123",
			Name: map[string]int{"13212":1223,"dasasd":213,"dsaj":123},
		},
		D:       []string{"13", "qad"},
		StuName: "sajf",
	}
	//t := make([]*model.Student,4)
	//t[0]=student
	//t[1]=&model.Student{}
	//a:= method.Marshel(t)
	s := quse.Marshel(student)
	fmt.Println(quickbytetostring(s))
	//fmt.Println(quickbytetostring(a))
	var test model.Student
	quse.Unmarshal(quickbytetostring(s),&test)
	//json.Unmarshal(s,&test)
	//json.Unmarshal(a,&test)
	fmt.Println(test)



	//var m map[string][]int
	//m =make(map[string][]int)
	//i := make([]int,10)
	//i[0]=1
	//i[1]=2
	//m["123"]=i
	//str:=method.Marshel(m)
	//fmt.Println(string(str))
	//
	//str1:=method.Marshel(i)
	//var shu []int
	//method.Unmarshal(string(str1),&shu)
	//fmt.Println(shu)
	//fmt.Println("总用时:",time.Now().Nanosecond()-t1.Nanosecond(),"纳秒")
}

func quickbytetostring(b[] byte)  string{
	return *(*string)(unsafe.Pointer(&b))
}
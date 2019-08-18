package main

import (
	"Myjson/method"
	"Myjson/model"
	"fmt"
	"time"
	"unsafe"
)

func main() {
	t1 := time.Now()
	student:=&model.Student{
		StuCode: 231,
		Si: model.Card{
			Number: "123",
			Name:   "2134",
			X: struct{ Ma string }{Ma: "213"},
		},
		D:       [3]string{"13", "qad"},
		StuName: "sajf",
	}
	//t := make([]*model.Student,4)
	//t[0]=student
	fmt.Println(student)
	a:= method.Marshel(student)
	fmt.Println(quickbytetostring(a))
	var test model.Student
	method.Unmarshal(quickbytetostring(a),&test)
	//json.Unmarshal(a,&test)
	fmt.Println(test)

	fmt.Println(time.Now().Nanosecond()-t1.Nanosecond())
}

func quickbytetostring(b[] byte)  string{
	return *(*string)(unsafe.Pointer(&b))
}

package main

import (
	"Myjson/method"
	"Myjson/model"
	"fmt"
	"unsafe"
)

func main() {
	student := &model.Student{
		StuCode:  1234,
		StuName:  "hgn",
		CardName: model.Card{
			Number: "123",
			Name:   "321",
			Man:    model.Peo{Ma:"13"},
		},
	}
	student1 :=make(map[string]model.Student)
	//student["123"]="123"
	//student["321"]="321"
	student1["123"]=*student
	b := method.Marshel(student)
	str := *(*string)(unsafe.Pointer(&b))

	var stu model.Student
	//s,_:=json.Marshal(student)
    method.Unmarshal(str,&stu)

	fmt.Println(stu)
	//fmt.Println(student1)
}

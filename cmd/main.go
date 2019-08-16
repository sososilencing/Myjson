package main

import (
	"Myjson/method"
	"fmt"
)

func main() {
	//student := &model.Student{
	//	StuCode:  1234,
	//	StuName:  "hgn",
	//	CardName: model.Card{
	//		Number: "123",
	//		Name:   "321",
	//		//Man:    model.Peo{Ma:"13"},
	//	},
	//}


	//b := method.Marshel(student1)
	//str := *(*string)(unsafe.Pointer(&b))
	sad := make(map[string]string)
	q := make(map[string]string)
	qqq := make(map[string]map[string]string)
	sad["x"]="xa"
	sad["xc"]="q"
	qqq["2131"]=sad
	q["xa"]="a"
	q["cv"]="xa"
	qqq["sdd"]=q
	qqq["xasd"] = map[string]string{
		"sda":"sda",
		"x1":"qa",
	}
	a:=method.Marshel(qqq)
	//a,_:=json.Marshal(qqq)
	fmt.Println(string(a))

	var stu map[int]string
	//stu = make(map[string]map[string]string)
	//stu = make([]string,3)
	//stu[1]="das"
	//b,_:=json.Marshal(student1)
	//b:=method.Marshel(stu)
	//fmt.Println(string(b))
    method.Unmarshal(string(a),&stu)
	//json.Unmarshal(a,&stu)
	fmt.Println(stu,"stu")
	//fmt.Println(student1)
}

package main

import (
	"bytes"
	"fmt"
	"reflect"
	"unsafe"
)

var buffer bytes.Buffer

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
	//
	//e:=&student
	//t:=&e
	//b:=Marshel(&t)
	//fmt.Printf(*(*string)(unsafe.Pointer(&b)))
	a :=make(map[string][]int)
	s:=make([]int,3)
	s[0]=1
	s[1]=2
	a["123"]=s
	b:=Marshel(a)
	str:=*(*string)(unsafe.Pointer(&b))
	fmt.Println(str)
}

type Name struct {
	V reflect.Value
	T reflect.Type
	viceV reflect.Value
	viceT reflect.Type
}
// 要写入一个结构体 或者全局变量
// 会有一个递归调用
// 用偏移量来增加 性能
// 递归调用 拿到指针 是莫得用的  需要得到他真实的类型 数值 才可以 操作 如果有多个&&&& 就循环拿到 相当于while  拿到实体
func Marshel(obj interface{}) bytes.Buffer{
	buffer.WriteString("{")
	marshel(obj)
	buffer.WriteString("}")
	return buffer
}

func marshel(obj interface{}) {
	v :=reflect.ValueOf(obj)
	t :=reflect.TypeOf(obj)
	name:=Name{
		V:     v,
		T:     t,
	}
	getkind1(name)
}

func getElem1(name Name){
	switch name.V.Kind() {
	case reflect.Ptr,reflect.Struct:
		getkind1(name)
	default:
		fmt.Println(name.V.Elem())
	}
}

func getkind1(name Name) func(){
	k:=name.V.Kind()
	switch k {
	case reflect.Bool:
		return getBool
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return getInt(name)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return getUint
	case reflect.Float32,reflect.Float64:
		return getFloat
	case reflect.String:
		fmt.Println(name.T.Kind())
		fmt.Println(name.V.String())
		return getString(name)
	case reflect.Interface:
		return getInterface
	case reflect.Struct:
		return getStruct(name)
	case reflect.Map:
		fmt.Println(name)
		return getMap(name)
	case reflect.Slice:
		return getSlice
	case reflect.Array:
		return getArray
	case reflect.Ptr:
		return getPtr(name)
	default:
		return nil
	}
}

func getBool()  {

}

func getInt(name Name)  func(){
	fmt.Println("我到达这里int")

	return nil
}

func getFloat(){

}
func getUint()  {

}
func getArray(){



}
func getInterface()  {

}
func getMap(name Name)  func(){
	for _,k:=range name.V.MapKeys(){
		//这里得到的是 key 值的 类型
		fmt.Println(k.Type())
		// this is about value
		value:=name.V.MapIndex(k)

		name:=Name{
			V: k,
			T: k.Type(),
			viceV:value,
			viceT:value.Type(),
		}
		fmt.Println(name)
		//  这个是 key 的值


		json(name)
		//name.V.MapIndex(k).Type() this is about value's type

	}
	return nil
}
func getPtr(name Name) func(){
	name.V=name.V.Elem()
	name.T=name.T.Elem()
	getElem1(name)
	return nil
}
func getSlice()  {

}
func getString(name Name)  func(){
	return nil
}
func getStruct(name Name) func(){
	num:=name.V.NumField()
	for i:=0;i<num;i++{
		f := name.V.Field(i)
		k := name.T.Field(i).Name
		fmt.Println(f,k)
	}
	return nil
}

func json(name Name)  func(){
	fmt.Println(name.viceT,name.viceV)
	//switch p.Kind() {
	//case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64,reflect.Float32,reflect.Float64:
	//	fmt.Println(p,value)
	//	return getInt()
	//case reflect.String:
	//	fmt.Println(p,value)
	//	return nil
	//}
	//name:=Name{
	//	V: value,
	//	T: p,
	//}
	//return getkind1(name)
	return nil
}
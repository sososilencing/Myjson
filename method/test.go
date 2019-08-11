package main

import (
	"Myjson/model"
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

var buffer bytes.Buffer

func main() {
	student := &model.Student{
		StuCode: 2018212152,
		StuName: "hgn",
		Thing: []string{
			"123",
			"23",
			"13",
		},
		CardName: model.Card{
			Number: "16534",
			Name:   "card",
		},
	}

	e := &student
	t := &e
	b := Marshel(&t)
	fmt.Printf(*(*string)(unsafe.Pointer(&b)))
}

type Queue struct {
	queue   []Queue
	name    []Name
	endname int
	endque  int
}

type Name struct {
	V reflect.Value
	T reflect.Type
}

// 要写入一个结构体 或者全局变量
// 会有一个递归调用
// 用偏移量来增加 性能
// 递归调用 拿到指针 是莫得用的  需要得到他真实的类型 数值 才可以 操作 如果有多个&&&& 就循环拿到 相当于while  拿到实体
func Marshel(obj interface{}) []byte {
	marshel(obj)
	return buffer.Bytes()
}

func marshel(obj interface{}) {
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)
	name := Name{
		V: v,
		T: t,
	}
	getkind1(name)
}

func getElem1(name Name) {
	switch name.V.Kind() {
	case reflect.Ptr, reflect.Struct:
		getkind1(name)
	default:
		fmt.Println(name.V.Elem())
	}
}

func getkind1(name Name) func() {
	k := name.V.Kind()
	switch k {
	case reflect.Bool:
		return getBool(name)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return getInt(name)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return getUint
	case reflect.Float32, reflect.Float64:
		return getFloat(name)
	case reflect.String:
		return getString(name)
	case reflect.Interface:
		return getInterface
	case reflect.Struct:
		return getStruct(name)
	case reflect.Map:
		return getMap(name)
	case reflect.Slice:
		return getSlice(name)
	case reflect.Array:
		return getArray(name)
	case reflect.Ptr:
		return getPtr(name)
	default:
		return nil
	}
}

func getBool(name Name) func() {
	b := name.V.Bool()
	buffer.WriteString(strconv.FormatBool(b))
	return nil
}

func getInt(name Name) func() {
	i := name.V.Int()
	buffer.WriteString(strconv.Itoa(int(i)))
	return nil
}

func getFloat(name Name) func() {
	v := name.V.Float()
	v1 := strconv.FormatFloat(v, 'f', -1, 64)
	buffer.WriteString(v1)
	return nil
}
func getUint() {

}
func getArray(name Name) func() {
	buffer.WriteString("[")
	for i := 0; i < name.V.Len(); i++ {
		json(Name{
			V: name.V.Index(i),
			T: name.V.Index(i).Type(),
		})
		if i != name.V.Len()-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("]")
	return nil
}
func getInterface() {

}
func getMap(name Name) func() {
	que := Queue{}
	que.name = make([]Name, 10)
	que.queue = make([]Queue, 10)
	for _, k := range name.V.MapKeys() {
		//这里得到的是 key 值的 类型
		//fmt.Println(k.Type())
		// this is about value
		value := name.V.MapIndex(k)
		name1 := Name{
			V: k,
			T: k.Type(),
		}

		name2 := Name{
			V: value,
			T: value.Type(),
		}
		//fmt.Println(name)
		//  这个是 key 的值

		que.name[que.endname] = name1
		que.endname++
		que.name[que.endname] = name2
		que.endname++
		//name.V.MapIndex(k).Type() this is about value's type
	}
	do(que)
	return nil
}
func getPtr(name Name) func() {
	name.V = name.V.Elem()
	name.T = name.T.Elem()
	getElem1(name)
	return nil
}
func getSlice(name Name) func() {
	buffer.WriteString("[")
	for i := 0; i < name.V.Len(); i++ {
		json(Name{
			V: name.V.Index(i),
			T: name.V.Index(i).Type(),
		})
		if i != name.V.Len()-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("]")
	return nil
}
func getString(name Name) func() {
	buffer.WriteString("\"" + name.V.String() + "\"")
	return nil
}
func getStruct(name Name) func() {
	que := Queue{}
	que.name = make([]Name, 10)
	que.queue = make([]Queue, 10)
	num := name.V.NumField()
	for i := 0; i < num; i++ {
		f := name.V.Field(i)

		k := name.T.Field(i)

		v := reflect.ValueOf(k.Name)
		//fmt.Println(v)
		name1 := Name{
			V: v,
			T: k.Type,
		}

		name2 := Name{
			V: f,
			T: f.Type(),
		}

		que.name[que.endname] = name1
		que.endname++
		que.name[que.endname] = name2
		que.endname++
	}
	do(que)
	return nil
}

func json(name Name) func() {
	t := name.T
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		getInt(name)
	case reflect.Float32, reflect.Float64:
		getFloat(name)
	case reflect.String:
		getString(name)
	default:
		getkind1(name)
	}
	return nil
}

func do(queue Queue) {
	buffer.WriteString("{")
	//分为每一个原子类型
	for i := 0; i < queue.endname; i++ {
		name := queue.name[i]
		if i%2 == 0 {
			//这里 永远都是一个string
			buffer.Write([]byte("\""))
			buffer.Write([]byte(name.V.String()))
			buffer.Write([]byte("\""))
			buffer.Write([]byte(":"))
		} else {
			json(name)
			if i != queue.endname-1 {
				buffer.WriteString(",")
			}
		}
	}
	buffer.WriteString("}")
}

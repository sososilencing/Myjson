package method

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

var buffer bytes.Buffer

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

//内部真正实现
func marshel(obj interface{}) {
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)
	name := Name{
		V: v,
		T: t,
	}
	getkind1(name)
}

//得到实体
func getElem1(name Name) {
	switch name.V.Kind() {
	case reflect.Ptr, reflect.Struct:
		getkind1(name)
	default:
		buffer.WriteString("null")
	}
}

//得到类型 进行判断 下面是各个类型的入口
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

//布尔类型
func getBool(name Name) func() {
	b := name.V.Bool()
	buffer.WriteString(strconv.FormatBool(b))
	return nil
}

//整型
func getInt(name Name) func() {
	i := name.V.Int()
	buffer.WriteString(strconv.Itoa(int(i)))
	return nil
}

// 浮点数
func getFloat(name Name) func() {
	v := name.V.Float()
	v1 := strconv.FormatFloat(v, 'f', -1, 32)
	buffer.WriteString(v1)
	return nil
}
func getUint() {

}

// 数组的序列化
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

//map 的序列化
func getMap(name Name) func() {
	que := Queue{}

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

		que.endname++
		que.name = append(que.name, name1)
		que.endname++
		que.name = append(que.name,name2)
		//name.V.MapIndex(k).Type() this is about value's type
	}
	do(que)
	return nil
}

//指针？？
func getPtr(name Name) func() {
	name.V = name.V.Elem()
	name.T = name.T.Elem()
	getElem1(name)
	return nil
}

//切片的序列化
func getSlice(name Name) func() {
	l := name.V.Len()
	if l == 0 {
		buffer.WriteString("null")
		return nil
	}
	buffer.WriteString("[")
	for i := 0; i < l; i++ {
		json(Name{
			V: name.V.Index(i),
			T: name.V.Index(i).Type(),
		})
		if i != l-1 {
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

// 结构体的 序列化
func getStruct(name Name) func() {
	que := Queue{}
	num := name.V.NumField()
	for i := 0; i < num; i++ {

		f := name.V.Field(i)
		k := name.T.Field(i)
		v := reflect.ValueOf(k.Name)

		name1 := Name{
			V: v,
			T: k.Type,
		}

		name2 := Name{
			V: f,
			T: f.Type(),
		}

		que.endname++
		que.name = append(que.name, name1)
		que.endname++
		que.name = append(que.name,name2)
	}
	do(que)
	return nil
}

// 这个又是一个判断 判断 第二次解析出来的该放在哪
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

//放在一个队列里面 然后 取出来 进行判断 然后写入
func do(queue Queue) {
	buffer.WriteString("{")
	//分为每一个原子类型
	for i := 0; i < queue.endname; i++ {
		name := queue.name[i]
		if i%2 == 0 {
			//这里 永远都是一个string
			buffer.Write([]byte("\""))
			switch name.V.Kind() {
			case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
				buffer.WriteString(strconv.Itoa(int(name.V.Int())))
				case reflect.String:
				buffer.WriteString(name.V.String())
			}
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

// 开始反序列化
// 反序列也需要递归样 而且 要入栈出栈
func Unmarshal(str string, i interface{}) {
 	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	if t.Kind() != reflect.Ptr {
		return
	}
	//fmt.Println(i)
	name := Name{
		V: v,
		T: t,
	}
	name.decide()
	//fmt.Println(str)
	fmt.Println(name.V.Kind())

	//把map 实例化  明天再想想
	name.V.Set(reflect.MakeMap(name.T))

	//这个是 map 里的key 类型判断 与 value 判断不同
	//fmt.Println(name.T.Key().Kind())
	//name.enter(str)
}

func(name *Name) enter(str string)  {
	switch name.V.Kind() {
	case reflect.Bool:
		name.unBool(str)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		name.unInt(str)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return
	case reflect.Float32, reflect.Float64:
		name.unFloat(str)
	case reflect.String:
		name.unString(str)
	case reflect.Interface:
		return
	case reflect.Struct:
		name.unmarshal(str)
	case reflect.Map:
		name.unMap(str)
	case reflect.Slice:
		name.unSlice(str)
	case reflect.Array:
		name.unArray(str)
	case reflect.Ptr:
		return
	default:
		return
	}
}

// 这是切片扩容 搞定了
func(name *Name) unSlice(str string)  {
	 if name.V.IsNil(){
	 	newv:=reflect.MakeSlice(name.T,name.V.Len(), len(str))
	 	reflect.Copy(newv,name.V)
	 	name.V.Set(newv)
	 }
	 name.V.SetLen(name.V.Len()+1)
	 name.V.Index(0).Set(reflect.ValueOf(1))
}


//没有写完  //不管了 虽然有点小bug
func(name *Name) unArray(str string)  {
	fmt.Println("array")
	reg := regexp.MustCompile("[0-9]+")
	num := reg.FindAllString(str,-1)
	fmt.Println(num)
	for i,k:=range num{
		if name.V.Len() > i{
			f,err := strconv.Atoi(k)
			if err != nil {
				return
			}
		name.V.Index(i).Set(reflect.ValueOf(f))
		}else {
			return
		}
	}
}

// 问题大得很
func(name *Name) unMap(str string)  int{
		begain := -1
		first := -1
		end := -1
		i := 0

		for i = 0;i < len(str);i++{
			if str[i]=='"' && begain == -1 && first ==-1{
				begain = i
			}else if str[i]=='"' && begain !=-1{
				end = i
				key := str[begain+1:end]

				k := name.T.Elem().Kind()
				//a := make(map[string]string)

				switch  k{
				case reflect.Map:
					name1:=&Name{
						V: name.V.MapIndex(reflect.ValueOf(key)),
						T: name.T.Elem(),
					}
					//fmt.Println(str[end+2:])
					i+=name1.unMap(str[end+2:])+1
					//fmt.Println(name.V.MapIndex(reflect.ValueOf(key)))
					//fmt.Println()
				}
				begain = -1
				end = -1
			}

			if str[i]==':' && first == -1 && begain == -1{
				first = i
			}else if str[i]==',' && first !=-1{
				end = i
				//value := str[first+1:end]
				//fmt.Println(value)
				//fmt.Println()
				end = -1
				first = -1
			}else if str[i]=='}' && first !=-1{
				return i
			}
		}
		return i
}

// 解析bool
func(name *Name) unBool(str string){
	reg := regexp.MustCompile(`{([\s\S]+)}`)
	if reg.MatchString(str){
		f :=reg.FindStringSubmatch(str)
		b,err := strconv.ParseBool(f[1])
		if err!=nil{
			return
		}
		name.V.SetBool(b)
	}else {
		b,err := strconv.ParseBool(str)
		if err != nil {
			return
		}
		name.V.SetBool(b)
	}
}

func(name *Name) unInt(str string)  {
	b,err:=strconv.Atoi(str)
	if err != nil {
		return
	}
	name.V.SetInt(int64(b))
}

func(name *Name) unFloat(str string)  {
	b,err :=strconv.ParseFloat(str,32)
	if err != nil {
		return
	}
	name.V.SetFloat(b)
}

func(name *Name) unString(str string)  {
	reg :=regexp.MustCompile("\"[\\s\\S]+\"[\\s\\S]{1}")
	if reg.MatchString(str){
		fmt.Println(str,"?")
		return
		//name.V.SetString(str)
	}
	name.V.SetString(str)
}
// 来来来面向对象 编程  把这个传入的 interface 变为对象 然后 解析 这个对象 给这个对象 赋值 赋值 然后 因为是指针 就可以改变了 嘻嘻嘻
// 写垃圾了 不能复用
func(name *Name) unmarshal(str string)  int{
	defer func() {
		if r:=recover();r!=nil{

		}
	}()

	var key string
	begin := -1
	end := -1
	first := -1
	var i int

	for i = 0;i<len(str);i++ {
		if first == -1 {
			if str[i] == '"' && begin != -1 {

				end = i
				b := str[begin+1 : end]
				key = b

				switch name.V.FieldByName(key).Kind() {
				case reflect.Struct:

					name1 := &Name{
						V: name.V.FieldByName(key),
					}

					i += name1.unmarshal(str[i+2:]) + 1

				case reflect.Slice:
				case reflect.Array:
				case reflect.Map:
				default:
				}
				begin = -1
				end = -1
			}else if str[i] == '"' && begin == -1 {
				begin = i
			}
		}
		// 这个 value 这里还要做处理才可以
		if begin == -1 {
			if str[i] == ':' {
				first = i
			} else if str[i] == ',' && first != -1 {
				end = i
				b := str[first+1 : end]
				name.set(key, b)
				end = -1
				first = -1
			} else if str[i] == '}' && first != -1 {
				end = i
				b := str[first+1 : end]
				name.set(key, b)
				end = -1
				first = -1
				return i
			}
		}
	}
	return i
}

func(name *Name) decide() {
	e:=name.T.Elem()
	name.V=name.V.Elem()
	name.T=name.T.Elem()
	switch e.Kind() {
	case reflect.Struct:

	case reflect.Ptr:
		name.getElem()
	}
}


func(name *Name)  getElem(){
	name.T = name.T.Elem()
	name.V = name.V.Elem()
	name.decide()
}

func(name *Name) set(s string,v string) (str string){
	defer func() {
		if r:=recover();r!=nil{
			str = "参数错误"
		}
	}()

	fmt.Println(s,":",v)
	c := name.V.FieldByName(s)
	switch c.Kind() {
	case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
		i,err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err.Error())
		}
		c.SetInt(int64(i))
	case reflect.String:
		c.SetString(v[1:len(v)-1])
	case reflect.Bool:
		b,err := strconv.ParseBool(v)
		if err != nil {
			fmt.Println(err.Error())
		}
		c.SetBool(b)
	case reflect.Float32,reflect.Float64:
		f,err := strconv.ParseFloat(v,64)
		if err != nil {
			fmt.Println(err.Error())
		}
		c.SetFloat(f)
	}
	return str
}
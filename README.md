# Quse
### 使用方法:



直接拉取这个包:https://github.com/sososilencing/Myjson

你就可以快乐的使用了

#### 函数: 

目前有 

1.marshel  :序列化

2.unmarshel : 反序列化

### 代码示例:

```go
package main
import  "Myjson/quse"
type Student struct {
	StuCode  int
	Si  Card
	D   []string
	StuName  string
}

func main() {
	student:=&model.Student{
		StuCode: 231,
		Si: model.Card{
			Number: "123",
			Name: map[string]int{"13212":1223,"dasasd":213,"dsaj":123},
		},
		D:       []string{"13", "qad"},
		StuName: "sajf",
		s := quse.Marshel(student)
		var test model.Student
	   quse.Unmarshal(string(s),&test)
	   fmt.Println(test)
	}
```

### About

后面如果还有机会,继续完善。
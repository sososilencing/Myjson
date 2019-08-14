package model

type Student struct {
	StuCode  int
	StuName  string
	CardName Card
}

type Card struct {
	Number string
	Name   string
	Man Peo
}
type peoplee struct {
	White string
	Year int
	Stu  student
	LangStr []string
	LangInt []int
	Langstu []student
}
type student struct {
	Name string
	Age int
}

type Peo struct {
	Ma string
}
package model

type Student struct {
	StuCode  int
	Si  Card
	D   [3]string
	StuName  string
}

type Card struct {
	Number string
	Name   string
	X Peo
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
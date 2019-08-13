package model

type Student struct {
	StuCode  int
	StuName  string
	CardName Card
	Thing    [2]string
}

type Card struct {
	Number string
	Name   string
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
	ManArray []man
}

type man struct {
	Man string
}
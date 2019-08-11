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

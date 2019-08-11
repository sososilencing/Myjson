package model

type Student struct {
	StuCode  int
	StuName  string
	CardName Card
}

type Card struct {
	Number string
	Name   string
}

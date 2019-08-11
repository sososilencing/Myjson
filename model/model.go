package model

type Student struct {
	StuCode int
	StuName string
	CardName Card
	Thing []string
}

type Card struct {
	Number string
	Name string
}
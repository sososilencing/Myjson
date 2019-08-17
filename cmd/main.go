package main

import (
	"Myjson/method"
	_"encoding/json"
	"fmt"
)

func main() {
	a := make(map[string]string)
	a["dsa"]="adsmfa"
	a["dsas"]="msankfdn"
	a["cmnak"]="amskfbn"
	fmt.Println(a)
	v:=method.Marshel(a)
	fmt.Println(string(v))
	var k map[string]string

	method.Unmarshal(string(v),&k)
	//json.Unmarshal()
	fmt.Println(k)

}

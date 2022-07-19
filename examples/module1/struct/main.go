package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type MyType struct {
	Name string `json:"name"`
	Address
}
type Address struct {
	City string `json:"city"`
}

func main() {
	mt := MyType{Name: "test", Address: Address{City: "shanghai"}}
	//通过&来取某个值的内存地址
	b, _ := json.Marshal(&mt)
	fmt.Println(string(b))
	myType := reflect.TypeOf(mt)
	name := myType.Field(0)
	tag := name.Tag.Get("json")
	println(tag)

	//addressName := myType.Field(1)
	//addressTag := addressName.Tag.Get("json")
	//println(addressTag)

	fmt.Printf("字段名称:%v, 字段类型:%v\n", myType.Field(1), myType.Field(1))
	myAddress := reflect.TypeOf(myType.Field(1).Type)
	//myAddress := reflect.TypeOf(mt.Address)

	fmt.Printf("字段名称:%v, 字段类型:%v\n", myAddress.Field(0), myAddress.Field(0))
	addressName := myAddress.Field(0)
	addressTag := addressName.Tag.Get("json")
	println(addressTag)

	tb := TypeB{P2: "p2", TypeA: TypeA{P1: "p1"}}
	//可以直接访问 TypeA.P1
	println(tb.P1)
}

type TypeA struct {
	P1 string
}

type TypeB struct {
	P2 string
	TypeA
}

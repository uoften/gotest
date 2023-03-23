package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type A struct {
	Id bson.ObjectId `bson:"redis"`
}

type B struct{
	Aid bson.ObjectId
}
func main() {
	//var a= new(A)
	//a.Id = bson.NewObjectId()
	//fmt.Println(a.Id.Hex())
	//fmt.Println("Hello, 世界", len("世界123asd!@#"), utf8.RuneCountInString("世123asd!@#\n\n\n\n界"))
	s      := "123456789"
	first3 := s[0:4]
	last3  := s[len(s)-4:]
	fmt.Println(first3)
	fmt.Println(last3)
}
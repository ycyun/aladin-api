package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	s := Hello("name")
	fmt.Println(s)
	if s != "Hello, name" {
		t.Error("Wrong result")
	}
}

//func TestGetBook(t *testing.T) {
//	s:=GetBook("8950991721")
//	a, err:=json.MarshalIndent(s, "", " ")
//
//	if err!=nil{
//		_=fmt.Errorf("%v", err)
//	}
//	fmt.Println(string(a))
//
//}

func TestSearchBook(t *testing.T) {
	//s := SearchBook("소드 아트")
	//fmt.Println(t.Name())
	//for _, book := range s {
	//	fmt.Printf("%8v원, %v\n", book.PriceSales, book.Title)
	//}
	//fmt.Printf("%v Books Searched\n", len(s))
	////b, _ := json.MarshalIndent(s[len(s)-1], "", "  ")
	////fmt.Printf("%v\n", string(b))
	//fmt.Printf("%v\n", s[len(s)-1].Isbn)
	//book := GetBook(s[len(s)-1].Isbn)
	book := GetBook("K462639633")
	b, _ := json.MarshalIndent(book, "", "  ")
	//b, _ := json.Marshal(book)
	fmt.Printf("%v\n", string(b))
	fmt.Println("")
	return
}

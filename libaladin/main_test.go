package libaladin

import (
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
	s := SearchBook("소드 아트")
	fmt.Println(t.Name())
	for _, book := range s {
		fmt.Printf("%8v원, %v\n", book.PriceSales, book.Title)
	}
	fmt.Printf("%v Books Searched", len(s))
	fmt.Println("")
	return
}

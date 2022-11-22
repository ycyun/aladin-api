package libaladin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	_book := SearchBookAuthor("카와하라 레키")
	//book, _ := json.MarshalIndent(_book, "", "  ")
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(_book)
	if err != nil {
		_ = fmt.Errorf("err!!!: %v\n", err)
	}
	fmt.Printf("%s\n", buffer.Bytes())
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

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	_book := GetBook("K622730603")
	//book, _ := json.MarshalIndent(_book, "", "  ")
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(_book)
	if err != nil {
		fmt.Errorf("err!!!: %v\n", err)
	}
	fmt.Printf("%s\n", buffer.Bytes())
}

package parser

import (
	"testing"
	"fmt"
	"io/ioutil"
)

func TestParserCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_data_test.html")
	if err != nil {
		panic(err)
	}
	const resultSize = 470
	fmt.Printf("%s\n", contents)
	result := ParserCityList(contents)
	if len(result.Requests) == resultSize {
		fmt.Println("ok")
	}
}

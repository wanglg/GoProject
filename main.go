package main

import (
	"fmt"
	"io"
	"golang.org/x/text/encoding"
	"golang.org/x/net/html/charset"
	"bufio"
	"regexp"
	"crawler/engine"
	"crawler/zhenai/parser"
)


func main() {
	//data, err := fetcher.Fetcher("http://www.zhenai.com/zhenghun")
	//if err != nil {
	//	panic(err)
	//}
	//printCityList(data)
	engine.SimpleEngine.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParesFunc: parser.ParserCityList,
	})




	//math := re.FindAllString(text, -1)
	//fmt.Println(math)
}
func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	maths := re.FindAllSubmatch(contents, -1)
	for _, m := range maths {
		fmt.Printf("city:%s url:%s\n", m[2], m[1])
	}
	fmt.Printf("city size %d\n", len(maths))
}
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

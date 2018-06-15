package parser

import (
	"crawler/engine"
	"regexp"
)

const citylistRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParserCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(citylistRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	limit:=3
	for _, m := range matches {
		result.Items = append(result.Items, "city:"+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParesFunc: ParseCity,
		})
		limit--
		if limit==0 {
			break
		}
		//fmt.Printf("city:%s url:%s\n", m[2], m[1])
	}
	//fmt.Printf("city size %d\n", len(matches))
	return result
}

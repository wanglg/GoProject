package parser

import (
	"crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		userName := string(m[2])
		result.Items = append(result.Items, "user:"+userName)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			//ParesFunc: engine.NilParse,
			ParesFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes, userName)
			},
		})
	}
	return result
}

package parser

import (
	"crawler/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)
//const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	//re := regexp.MustCompile(profileRe)
	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		userName := string(m[2])
		//result.Items = append(result.Items, "user:"+userName)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			//ParesFunc: engine.NilParse,
			ParesFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes, userName)
			},
		})
	}
	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			//ParesFunc: engine.NilParse,
			ParesFunc: ParseCity,
		})
	}
	return result
}

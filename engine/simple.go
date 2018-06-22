package engine

import (
	"crawler/fetcher"
	"log"
)

type SimpleEngine struct {
}

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parserResult, err := worker(r)
		if err != nil {
			log.Printf("error urr ：%s : %v", r.Url, err)
			continue
		}
		requests = append(requests, parserResult.Requests...)
		for _, item := range parserResult.Items {
			log.Printf("item:%v", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetcher(r.Url)
	if err != nil {
		log.Printf("error urr ：%s : %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParesFunc(body), nil

}

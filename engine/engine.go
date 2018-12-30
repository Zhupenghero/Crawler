package engine

import (
	"crawler/fetch"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, s := range seeds {
		requests = append(requests, s)
	}
	for len(requests) > 0 {
		request := requests[0]
		log.Printf("Fetching %v", request.Url)
		requests = requests[1:]

		body, err := fetch.Fetch(request.Url)
		if err != nil {
			log.Printf("fetching %s err,Fetch err:%v", request.Url, err)
			continue
		}
		parseResult := request.ParseFunc(body)
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got Item :%v", item)
		}
	}
}

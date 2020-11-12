package main

import (
	"gin/crawler/fetcher"
	"gin/crawler/zhenai/structs"
	"log"
)

/*
执行爬虫引擎 可以接收多个种子
*/
func run(seeds ...structs.Request) {
	var requests []structs.Request
	for _, request := range seeds {
		requests = append(requests, request)
	}
	for i := 0; i < len(requests); i++ {
		request := requests[i]
		byte, err := fetcher.Fetch(request.Url)
		if err != nil {
			log.Printf("Fetcher error url:%s, err:%s", request.Url, err)
			continue
		}
		result := request.ParserFunc(byte)
		requests = append(
			requests,
			result.Requests...,
		)
		items := result.Items
		for _, item := range items {
			log.Println(item)
		}
	}
}

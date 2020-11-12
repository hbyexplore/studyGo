package parser

import (
	"fmt"
	"gin/crawler/zhenai/structs"
	"regexp"
)

var movieList = regexp.MustCompile(`<a class="v_playBtn" href="(/vod-read-id-[0-9]+.html)" target="_blank" title="([^\"]*)"><i></i></a>`)

func MovieListParser(body *[]byte) *structs.ParseResult {
	content := movieList.FindAllSubmatch(*body, -1)
	result := structs.ParseResult{}
	for _, v := range content {
		result.Items = append(result.Items, string(v[2]))
		result.Requests = append(result.Requests, structs.Request{
			Url:        "http://www.tv8box.com" + string(v[1]),
			ParserFunc: MovieParser,
		})
	}
	fmt.Printf("Matches found:%d\n", len(content))
	return &result
}

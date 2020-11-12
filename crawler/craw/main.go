package main

import (
	"gin/crawler/dianyinghezi/parser"
	"gin/crawler/zhenai/structs"
)

const zhenaiUrl = "http://www.zhenai.com/zhenghun/"
const maoyanUrl = "https://dianyinghezi.com/films"
const dianyinhezi = "http://www.tv8box.com/list-select-id-1-type--area--year--star--state--order-addtime.html"

func main() {
	run(structs.Request{
		Url:        dianyinhezi,
		ParserFunc: parser.MovieListParser,
	})
	/*if err != nil {
		fmt.Println("爬虫遭遇异常:",err)
	}
	for k,v:=range *cityAndUrl{
		fmt.Printf("key:%s,value:%s\n",k,v)
	}*/

}

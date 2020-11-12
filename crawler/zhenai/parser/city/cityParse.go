package city

import (
	"fmt"
	"gin/crawler/zhenai/parser/people"
	"gin/crawler/zhenai/structs"
	"regexp"
	"strings"
)

const cityReg = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]*)</a>`

func CityParser(body *[]byte) *structs.ParseResult {
	reg := regexp.MustCompile(cityReg)
	content := reg.FindAllSubmatch(*body, -1)
	result := structs.ParseResult{}
	for _, v := range content {
		str := string(v[1])
		str = strings.Replace(str, "p", "ps", 1)
		result.Items = append(result.Items, string(v[2]))
		result.Requests = append(result.Requests, structs.Request{
			Url:        str,
			ParserFunc: people.PeopleInfoParse,
		})
	}
	fmt.Printf("Matches found:%d\n", len(content))
	return &result
}

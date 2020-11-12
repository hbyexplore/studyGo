package city

import (
	"fmt"
	"gin/crawler/zhenai/structs"
	"regexp"
)

const cityListreg = `<a * href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]*)</a>`

func CityListParser(body *[]byte) *structs.ParseResult {
	reg := regexp.MustCompile(cityListreg)
	content := reg.FindAllSubmatch(*body, -1)
	result := structs.ParseResult{}
	for i, v := range content {
		if i < 1 {
			result.Items = append(result.Items, string(v[2]))
			result.Requests = append(result.Requests, structs.Request{
				Url:        string(v[1]),
				ParserFunc: CityParser,
			})
		}
		break
	}
	fmt.Printf("Matches found:%d\n", len(content))
	return &result
}

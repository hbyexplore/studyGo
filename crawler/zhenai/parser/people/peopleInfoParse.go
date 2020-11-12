package people

import (
	"fmt"
	"gin/crawler/zhenai/structs"
	"regexp"
)

var (
	basicInfoReg     = regexp.MustCompile(`<div * class="m-btn purple">([^<]*)</div>`)
	detailInfoReg    = regexp.MustCompile(`<div * class="m-btn pink">([^<]*)</div>`)
	objectInfoReg    = regexp.MustCompile(`<div * class="m-btn">([^<]*)</div>`)
	innerOsReg       = regexp.MustCompile(`<div * class="m-btn">([^<]*)</div>`)
	peopleHobbiesReg = regexp.MustCompile(`<div * class="m-btn">([^<]*)</div>`)
)

func PeopleInfoParse(body *[]byte) *structs.ParseResult {
	basicInfoByte := basicInfoReg.FindAllSubmatch(*body, -1)
	for _, value := range basicInfoByte {
		fmt.Println(value)
	}
	return &structs.ParseResult{
		Requests: nil,
		Items:    nil,
	}
}

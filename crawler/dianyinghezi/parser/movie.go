package parser

import (
	"fmt"
	"gin/crawler/dianyinghezi/entity"
	"gin/crawler/zhenai/structs"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	//图片
	imageReg = regexp.MustCompile(`<img src="(http://pic.tcc-interiors.com/vod/[0-9]+-[0-9]+-[0-9]+/[0-9a-z]+.jpg)" data-src="*" alt="*" style="display: block;">`)
	//主演
	actorReg = regexp.MustCompile(`<a href="/vod-search-actor-*.html">([^<])*</a>`)
	//导演
	directorReg = regexp.MustCompile(`<a href="/vod-search-director-*.html">([^<])*</a>`)
	//类型
	movieTypeReg = regexp.MustCompile(`<a href="/list-read-id-1.html">([^<]*)</a>`)
	//别名
	aliasReg = regexp.MustCompile(`<li class="clearfix"><em class="em_tit">又名：</em>([^<]*)</li>`)
	//简介
	synopsisReg = regexp.MustCompile(`<span>([^<])*</span>`)
	//评分
	scoreReg = regexp.MustCompile(`<em class="em_score ff-score-val">([^<])*</em>`)
)

func MovieParser(body *[]byte) *structs.ParseResult {
	result := structs.ParseResult{}
	if len(*body) <= 1400 {
		return &result
	}
	yearsAndTile := getYearsAndTitle(body)
	image := toString(imageReg.FindSubmatch(*body))
	actor := toString(actorReg.FindSubmatch(*body))
	director := toString(directorReg.FindSubmatch(*body))
	movieType := toString(movieTypeReg.FindSubmatch(*body))
	alias := toString(aliasReg.FindSubmatch(*body))
	synopsis := toString(synopsisReg.FindSubmatch(*body))
	score := toString(scoreReg.FindSubmatch(*body))
	movie := entity.Movies{
		Title:     yearsAndTile[0],
		ImageUrl:  image,
		Actor:     actor,
		Director:  director,
		MovieType: movieType,
		Alias:     alias,
		Synopsis:  synopsis,
		Score:     score,
		Years:     yearsAndTile[1],
	}
	result.Items = append(result.Items, nil)
	result.Requests = append(result.Requests, structs.Request{
		Url:        "",
		ParserFunc: nil,
	})
	fmt.Printf("MovieInfo:%v", movie)
	return &result
}
func toString(bytes [][]byte) string {
	return string(bytes[1])
}
func getYearsAndTitle(body *[]byte) [2]string {
	readerBody := *body
	dou, err := goquery.NewDocumentFromReader(strings.NewReader(string(readerBody)))
	if err != nil {
		fmt.Println("生成Document对象失败!!!")
		return [2]string{}
	}
	years := dou.Find("em:contains(年代：)+em").Text()
	title := dou.Find("div.txt_intro_con>div>h1").Text()
	return [2]string{title, years}
}

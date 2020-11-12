package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

func main() {
	client := http.Client{Timeout: time.Second * 3}
	req, err := http.NewRequest("GET", "http://www.zhenai.com/zhenghun", nil)
	if err != nil {
		fmt.Println("创建请求失败,", err)
		return
	}
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("http get error", err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read error", err)
		return
	}
	//fmt.Printf("%s\n",body)
	//用正则表达式匹配获取所有的网址
	cityAndUrl := reg(body)
	fmt.Println(cityAndUrl)
	/*links := collectlinks.All(res.Body)
	for _, link := range links {
		fmt.Println("parse url", link)
	}*/

}

func reg(body []byte) (cityAndUrl map[string]string) {
	reg := regexp.MustCompile(`<a * href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]*)</a>`)
	content := reg.FindAllSubmatch(body, -1)
	cityAndUrl = make(map[string]string, 500)
	for _, v := range content {
		//fmt.Printf("url:%s, city:%s\n",v[1],v[2])
		url := string(v[1])
		city := string(v[2])
		cityAndUrl[city] = url
	}
	fmt.Printf("Matches found:%d\n", len(content))
	return
}

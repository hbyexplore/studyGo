package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Fetch(url string) (*[]byte, error) {
	client := http.Client{Timeout: time.Second * 15}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	//cookie:="FSSBBIl1UgzbN7NO=5Ae4MhNByMRPcHA1g3y0Qd.IErkjtNe4Qg6IAx4RG6qiQyF6sh271k50BtpVdlu_6CWVvg.uSzO4yzjSONCqgRq; FSSBBIl1UgzbN7NP=5U_u7U2574qlqqqmTX6z1Pq1AxyQxqKSFcHKHtOVrzpaTz7It3gOx3vlMmlSyT4c6Ps7sZ2UzhRzAjPtBMlYqrWTPIpDygXiY_RKv.DNUFUzLQJz8wZpP3t7Vdpy6ElUOaF66uU9_x1Dlb0HqK.yLv85bsuMCSW0u_mFVzxZwLEykzLWB8Gst0eg_mVbmlGSvrzhPWDV4JhBeOEl6KQiMzU7WzA.n6ORbUbLOASSF8vwSnDm7vc2bGHSy7tTxMvz6uv0Gs.Ir9cwjFrJdmSXHOS; sid=b7be8106-9028-47fb-aa06-ae44ed7f7df1; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1605167808; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1605168130; _exid=TnCqk73BYeiCh6ZZlRy6C9HGu9JINgUKPHqa868wH0Xv4aaGelvruKyiYS0w%2BUEgGIvw9Z01JYVjpB4qbK8rLg%3D%3D; ec=9CCw0RNK-1605167809243-b7d8d4618a7dc1779127441; _efmdata=pNpLKTxzBa5jiEr%2FAHxKIAn6bYgtEmbw1KqQtDIHMUhtAwVRWadA5imat3EAYqBcvBdrBqG6b8oGyxaafsAoJW2trI2ZFlZcmAsnJ8AC1y8%3D"
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:82.0) Gecko/20100101 Firefox/82.0")
	//req.Header.Set("Cookie",cookie)
	//req.Header.Set("Host","album.zhenai.com")
	//req.Header.Set("Referer","https://album.zhenai.com/u/1412872831")
	//req.Header.Set("TE","Trailers")
	//req.Header.Set("Connection","keep-alive")
	//req.Header.Set("Content-Type","application/x-www-form-urlencoded;charset=utf-8")
	//req.Header.Set("Accept","application/json, text/plain, */*")
	//req.Header.Set("Accept-Encoding","gzip, deflate, br")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	for res.StatusCode != http.StatusOK {
		fmt.Errorf("请求响应错误,信息为:%s", res.Status)
		res, err = client.Do(req)
		fmt.Println("状态码为:", res.Status)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return &body, nil
}

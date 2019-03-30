package fetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(10 * time.Millisecond)

func Fetcher(url string) (contents []byte, err error) {
	<-rateLimiter
	resp, err := http.Get(url)
	if err != nil || resp == nil {
		log.Print("出错了,正在重试", err)
		//如果出错了，每秒默认重试一次
		time.Sleep(time.Millisecond * 300)
		_, _ = Fetcher(url)
	}
	//defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code:%d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}

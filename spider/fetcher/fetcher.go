package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//500毫秒请求一次
var rateLimiteTime = time.Tick(1*time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<- rateLimiteTime

	log.Printf("Fetcher url is %s", url)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")
	resp,err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetch error code %d", resp.StatusCode)
	}

	reader := bufio.NewReader(resp.Body)
	//返回正确的编码类型
	e := determineEncoding(reader)

	//将网页编码转换为utf-8
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

//检测内容的编码类型
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	//获取前面1024个字节
	b, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	//检测内容的编码类型，需传1024个字节进去检测
	e, _, _ := charset.DetermineEncoding(b, "")
	return e
}
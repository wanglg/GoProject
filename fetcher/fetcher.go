package fetcher

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/unicode"
	"log"
)

func Fetcher(url string) ([]byte, error) {
	fmt.Println("Fetcher:url->" + url)
	//resp, err := http.Get(url)
	//if err != nil {
	//	return nil, err;
	//}
	//defer resp.Body.Close()
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodeyReader := bufio.NewReader(resp.Body)
		e := determineEncoding(bodeyReader)
		utf8Reader := transform.NewReader(bodeyReader, e.NewDecoder())
		return ioutil.ReadAll(utf8Reader)
	} else {
		return nil, fmt.Errorf("wrong status code : %d", resp.StatusCode)
	}
}
func determineEncoding(reader *bufio.Reader) encoding.Encoding {
	bytes, err := reader.Peek(1024)
	if err != nil {
		log.Println("determineEncoding err")
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

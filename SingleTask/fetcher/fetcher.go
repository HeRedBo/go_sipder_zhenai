package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	// e := determineEncodeing(resp.Body)
	e := determineEncodeing(bodyReader)
	//utf8Reader := transform.NewReader(resp.Body,e.NewDecoder())
	utf8Reader := transform.NewReader(bodyReader,e.NewDecoder())
	return ioutil.ReadAll(utf8Reader);
}

func determineEncodeing(r *bufio.Reader) encoding.Encoding {
	//bytes, err := bufio.NewReader(r).Peek(1024)
	bytes, err := r.Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e,_,_ := charset.DetermineEncoding(bytes,"")
	return e
}

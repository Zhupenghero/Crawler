package fetch

import (
	"net/http"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/unicode"
	"log"
)

func Fetch(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		err := fmt.Errorf("ERR STATUS CODE:%d", res.StatusCode)
		return nil, err
	}
	bodyReader := bufio.NewReader(res.Body)
	e := determineEncoding(bodyReader)
	gbkReader := transform.NewReader(bodyReader, e.NewEncoder())
	return ioutil.ReadAll(gbkReader)
}
func determineEncoding(r io.Reader) encoding.Encoding{
	bytes,err:= bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Println(err)
		return unicode.UTF8
	}
	e,_,_:=charset.DetermineEncoding(bytes,"")
	return e
}

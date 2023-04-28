package client

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) io.Reader {
	res, err := http.Get(url)
	if err != nil {
		fmt.Print("请求错误", err)
	}
	//data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Print("数据错误", err)
	}
	return res.Body
}

func FetchString(url string) string {
	res, err := http.Get(url)
	if err != nil {
		fmt.Print("请求错误", err)
		return ""
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Print("数据错误", err)
		return ""
	}
	return string(data)
}

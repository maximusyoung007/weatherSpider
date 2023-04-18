package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Connect(url string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Print("请求错误", err)
		return
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Print("数据错误", err)
		return
	}
	fmt.Print(string(data))
}

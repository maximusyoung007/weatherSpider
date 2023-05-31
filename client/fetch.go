package client

import (
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"weatherSpider/logu"
)

var log = &logu.Logger

func Fetch(url string) io.Reader {
	res, err := http.Get(url)
	if err != nil {
		(*log).WithFields(logrus.Fields{"error": err}).Error("请求错误")
	}
	if err != nil {
		(*log).WithFields(logrus.Fields{"error": err}).Error("数据错误")
	}
	return res.Body
}

func FetchString(url string) string {
	res, err := http.Get(url)
	if err != nil {
		(*log).WithFields(logrus.Fields{"error": err}).Error("请求错误")
		return ""
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		(*log).WithFields(logrus.Fields{"error": err}).Error("数据错误")
		return ""
	}
	return string(data)
}

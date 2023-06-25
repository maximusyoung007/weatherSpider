package mail

import (
	"bytes"
	"fmt"
	"github.com/jordan-wright/email"
	"html/template"
	"net/smtp"
	"strconv"
	"time"
	"weatherSpider/database"
	"weatherSpider/logu"
	"weatherSpider/structs"
)

var log = &logu.Logger

func SendEmail() {
	em := email.NewEmail()

	em.From = "1091261998@qq.com"
	em.To = []string{"1091261998@qq.com"}

	em.Subject = "发送邮件测试"
	//em.Text = []byte("发个邮件试试")
	areaList := database.GetRow()
	ec := &structs.EmailContext{}
	ec.AreaList = areaList
	ec.Date = strconv.Itoa(time.Now().Year()) + "年" + strconv.Itoa(int(time.Now().Month())) + "月" + strconv.Itoa(time.Now().Day()) + "日"
	body, _ := template.ParseFiles("templates/email.html")
	fmt.Println(body)
	buf := new(bytes.Buffer)
	body.Execute(buf, ec)
	em.HTML = []byte(buf.Bytes())
	err := em.Send("smtp.qq.com:25", smtp.PlainAuth("", "1091261998@qq.com", "njvicqgfdqevjehi", "smtp.qq.com"))
	if err != nil {
		(*log).Info(err)
	}
	(*log).Info("send successfully...")

}

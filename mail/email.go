package mail

import (
	"github.com/jordan-wright/email"
	"net/smtp"
	"weatherSpider/logu"
)

var log = &logu.Logger

func SendEmail() {
	em := email.NewEmail()

	em.From = "1091261998@qq.com"
	em.To = []string{"1091261998@qq.com"}

	em.Subject = "发送邮件测试"
	em.Text = []byte("发个邮件试试")
	err := em.Send("smtp.qq.com:25", smtp.PlainAuth("", "1091261998@qq.com", "njvicqgfdqevjehi", "smtp.qq.com"))
	if err != nil {
		(*log).Info(err)
	}
	(*log).Info("send successfully...")
}

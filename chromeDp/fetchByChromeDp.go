package chromeDp

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/sirupsen/logrus"
	"log"
	"time"
	"weatherSpider/logu"
)

// 处理单个请求
func GetHttpHtmlContent(url string, selector string, sel interface{}) (string, error) {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true),
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.UserAgent(`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36`),
	}

	//初始化参数，先传一个空的数据
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	c, _ := chromedp.NewExecAllocator(context.Background(), options...)

	// create context
	chromeCtx, cancel := chromedp.NewContext(c, chromedp.WithLogf(log.Printf))
	defer cancel()
	// 执行一个空task, 用提前创建Chrome实例
	chromedp.Run(chromeCtx, make([]chromedp.Action, 0, 1)...)

	//创建一个上下文，超时时间为40s
	timeoutCtx, cancel := context.WithTimeout(chromeCtx, 40*time.Second)
	defer cancel()

	var htmlContent string
	err := chromedp.Run(timeoutCtx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(selector),
		chromedp.OuterHTML(sel, &htmlContent, chromedp.ByJSPath),
	)

	//select {
	//case <-timeoutCtx.Done(): // 因为ctx带超时参数，当时间期限到了之后就会走到这里退出协程
	//	//if err != nil {
	//	//fmt.Errorf("请求异常", err)
	//	//e := chromedp.Cancel(timeoutCtx)
	//	//t, _ := timeoutCtx.Deadline()
	//	//logu.Logger.WithFields(logrus.Fields{"time": t})
	//	//cancel()
	//	//logu.Logger.WithFields(logrus.Fields{"info": "关闭context异常"}).Error(e)
	//
	//	logu.Logger.WithFields(logrus.Fields{"定位": "chromeDp请求异常"}).Error(err)
	//
	//	//}
	//	fmt.Println("Done. ", time.Now().Format("2006-01-02 15:04:05"))
	//	//chromedp.Cancel(timeoutCtx)
	//	cancel()
	//	return "", err
	//
	//default: // 协程循环执行for，当ctx.Done()无信号时总是走到Default分支
	//	//fmt.Println("case default ", time.Now().Format("2006-01-02 15:04:05"))
	//	//time.Sleep(time.Second)
	//
	//}

	if err != nil {
		cancel()
		e := chromedp.Cancel(timeoutCtx)
		logu.Logger.WithFields(logrus.Fields{"定位": "chromeDp请求异常"}).Error(err)
		logu.Logger.Error(e)
		return "", err
	}

	//每次请求完成后关闭context
	chromedp.Cancel(timeoutCtx)
	//cancel()

	return htmlContent, nil
}

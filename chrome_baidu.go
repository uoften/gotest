package main


import (

	"context"

	"fmt"

	"io/ioutil"

	"time"


	"github.com/chromedp/chromedp"

)


func main() {

	// 参数设置

	options := []chromedp.ExecAllocatorOption{

		chromedp.Flag("headless", false),

		chromedp.UserAgent("Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36"),

	}

	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), options...)

	defer cancel()


	// 创建chrome示例

	ctx, cancel := chromedp.NewContext(allocCtx)

	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)

	defer cancel()


	var (

		buf   []byte

		value string

	)

	err := chromedp.Run(ctx,

		chromedp.Tasks{

			// 打开导航

			chromedp.Navigate("https://baidu.com/"),

			// 等待元素加载完成

			chromedp.WaitVisible("body", chromedp.ByQuery),

			// 输入chromedp

			chromedp.SendKeys(".gLFyf.gsfi", "chromedp", chromedp.NodeVisible),

			// 打印输入框的值

			chromedp.Value(".gLFyf.gsfi", &value),

			// 提交

			chromedp.Submit(".gLFyf.gsfi", chromedp.ByQuery),

			chromedp.Sleep(3 * time.Second),

			// 截图

			chromedp.CaptureScreenshot(&buf),

		},

	)

	if err != nil {

		fmt.Println(err)

	}

	fmt.Println("value: ", value)

	if err := ioutil.WriteFile("fullScreenshot.png", buf, 0644); err != nil {

		fmt.Println(err)

	}

}
package main

import (
	"fmt"
	"gotest/library"
)

func main()  {
	urlStr:="https://imgconvert.csdnimg.cn/ahr0chm6ly9tbwjpei5xcgljlmnul21tyml6x3buzy9bme8yuw1pznhcbu9ielprb2ditny0qudgchbsvkxxmmjrcglibxk2s21bawnrtjdlmlfptulcyuxxylnnulpucexpyjr6nvhomuvnmgvmntlkdgrwneuwus82nda?x-oss-process=image/format,png"
	res, err := library.GetCloudflare(urlStr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

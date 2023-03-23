package main

func main()  {
	spider.NewSpider(NewMyPageProcesser(), "TaskName").
		AddUrl("https://github.com/hu17889?tab=repositories", "html"). // Start url, html is the responce type ("html" or "json")
		AddPipeline(pipeline.NewPipelineConsole()).                    // Print result on screen
		SetThreadnum(3).                                               // Crawl request by three Coroutines
		Run()
}
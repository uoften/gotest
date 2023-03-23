package s_spider

import (
	"gotest/common/utils"
	"gotest/library"
	"net"
	"net/http"
	"strconv"

	"gotest/config"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"go.uber.org/zap"
)

var SpiderService = &spiderService{}

type spiderService struct {
	novelListCollector   *colly.Collector
	chapterListCollector *colly.Collector
	chapterCollector     *colly.Collector
}

type novel struct {
	Title                    string
	Author                   string
	Category                 string
	Summary                  string
	ChapterCount             int
	WordCount                string
	CoverSrcUrl              string
	NovelSrcUrl              string
	CurrentCrawChapterPageNo int
}
type chapter struct {
	Novel         *novel
	Title         string
	ChapterSrcUrl string
	Content       string
	Sort          int
}

/**
生成一个collector对象
*/
func (this *spiderService) NewCollector() *colly.Collector {
	collector := colly.NewCollector()
	collector.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   90 * time.Second,
			KeepAlive: 90 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   90 * time.Second,
		ExpectContinueTimeout: 90 * time.Second,
	})

	//是否允许相同url重复请求
	collector.AllowURLRevisit = config.GlobalConfig.SpiderAllowUrlRevisit

	//默认是同步,配置为异步,这样会提高抓取效率
	collector.Async = config.GlobalConfig.SpiderAsync

	collector.DetectCharset = true

	// 对于匹配的域名(当前配置为任何域名),将请求并发数配置为2
	//通过测试发现,RandomDelay参数对于同步模式也生效
	if err := collector.Limit(&colly.LimitRule{
		// glob模式匹配域名
		DomainGlob: config.GlobalConfig.SpiderLimitRuleDomainGlob,
		// 匹配到的域名的并发请求数
		Parallelism: config.GlobalConfig.SpiderLimitRuleParallelism,
		// 在发起一个新请求时的随机等待时间
		RandomDelay: time.Duration(config.GlobalConfig.SpiderLimitRuleRandomDelay) * time.Second,
	}); err != nil {
		utils.Logger.Error("生成一个collector对象, 限速配置失败", zap.Error(err))
	}

	//配置反爬策略(设置ua和refer扩展)
	extensions.RandomUserAgent(collector)
	extensions.Referer(collector)

	return collector
}

/**
初始化collector
*/
func (this *spiderService) initCollector() {
	this.configNovelListCollector()
	this.configChapterListCollector()
	this.configChapterCollector()
}

/**
配置NovelListCollector
*/
func (this *spiderService) configNovelListCollector() {
	//避免对collector对象的每个回调注册多次, 否则回调内逻辑重复执行多次, 会引发逻辑错误
	if this.novelListCollector != nil {
		return
	}
	this.novelListCollector = this.NewCollector()

	this.novelListCollector.OnHTML("div.list_main li", func(element *colly.HTMLElement) {
		// 抽取某小说的入口页面地址和章节列表页的入口地址
		novelUrl, exist := element.DOM.Find("div.book-img-box a").Attr("href")
		if !exist {
			utils.Logger.Error("爬取小说列表页, 抽取当前小说的入口url, 异常", zap.Any("novelUrl", novelUrl))
			return
		}
		chapterListUrl := strings.ReplaceAll(novelUrl, "book", "chapter")
		utils.Logger.Info("爬取小说列表页, 抽取章节列表的入口url, 完成", zap.Any("chapterListUrl", chapterListUrl))

		//抽取小说剩余信息，并组装novel对象
		novel := &novel{}
		novel.Title = strings.TrimSpace(element.DOM.Find("div.book-mid-info p.t").Text())
		novel.NovelSrcUrl = chapterListUrl
		novel.CoverSrcUrl = element.DOM.Find("div.book-img-box img").AttrOr("src", "")
		novel.Author = strings.TrimSpace(element.DOM.Find("div.book-mid-info p.author span").First().Text())
		novel.Category = strings.TrimSpace(element.DOM.Find("div.book-mid-info p.author a").Text())
		novel.Summary = strings.TrimSpace(element.DOM.Find("div.book-mid-info p.intro").Text())
		novel.WordCount = strings.TrimSpace(element.DOM.Find("div.book-mid-info p.update").Text())

		// 创建上下文对象
		ctx := colly.NewContext()
		ctx.Put("novel", novel)

		// 爬取章节列表页
		utils.Logger.Info("爬取小说列表页, 开始", zap.Any("novelTitle", novel.Title), zap.Any("chapterListUrl", chapterListUrl))
		if err := this.chapterListCollector.Request("GET", chapterListUrl, nil, ctx, nil); err != nil {
			utils.Logger.Error("爬取小说列表页, 爬取章节列表页, 异常", zap.Any("chapterListUrl", chapterListUrl))
			return
		}
	})

	/**
	爬取当前列表页的下一页
	*/
	this.novelListCollector.OnHTML("div.tspage a.next", func(element *colly.HTMLElement) {
		nextUrl := element.Request.AbsoluteURL(element.Attr("href"))
		utils.Logger.Info("爬取小说列表页的下一页, 开始", zap.Any("nextUrl", nextUrl))

		if err := this.novelListCollector.Visit(nextUrl); err != nil {
			utils.Logger.Error("爬取小说列表页的下一页, 异常", zap.Any("nextUrl", nextUrl), zap.Error(err))
			return
		}

		utils.Logger.Info("爬取小说列表页的下一页, 完成", zap.Any("nextUrl", nextUrl))
	})

	this.novelListCollector.OnError(func(response *colly.Response, e error) {
		utils.Logger.Error("爬取小说列表页, OnError", zap.Any("url", response.Request.URL.String()), zap.Error(e))

		//请求重试
		response.Request.Retry()
	})

	utils.Logger.Info("配置NovelListCollector, 完成")
}

/**
配置ChapterListCollector
*/
func (this *spiderService) configChapterListCollector() {
	if this.chapterListCollector != nil {
		return
	}
	this.chapterListCollector = this.NewCollector()

	this.chapterListCollector.OnRequest(func(r *colly.Request) {
		utils.Logger.Info("爬取章节列表页, OnRequest", zap.Any("url", r.URL.String()))
	})
	// 从章节列表页抓取第一章节的入口地址
	this.chapterListCollector.OnHTML("div.catalog_b li:nth-child(1) a", func(h *colly.HTMLElement) {
		// 抽取某章节的地址
		chapterUrl, exist := h.DOM.Attr("href")
		if !exist {
			utils.Logger.Error("爬取章节列表页, 爬取第1章, 抽取chapterUrl, 异常", zap.Any("srcUrl", h.Request.URL))
			return
		}
		chapterUrl = h.Request.AbsoluteURL(chapterUrl)
		chapterTitle := h.DOM.Text()
		utils.Logger.Info("爬取章节列表页, 爬取第1章, 抽取chapterUrl, 完成", zap.Any("chapterUrl", chapterUrl), zap.Any("chapterTitle", chapterTitle))

		// 获取上下文信息
		novel := h.Response.Ctx.GetAny("novel").(*novel)
		novel.ChapterCount = h.DOM.Parent().Parent().Find("li").Length()
		novel.CurrentCrawChapterPageNo = 0

		// 爬取章节
		utils.Logger.Info("爬取章节列表页, 开始爬取第1章", zap.Any("novelTitle", novel.Title), zap.Any("chapterTitle", chapterTitle))
		if err := this.chapterCollector.Request("GET", chapterUrl, nil, h.Response.Ctx, nil); err != nil {
			utils.Logger.Error("爬取章节列表页, 爬取第1章, 异常", zap.Any("chapterUrl", chapterUrl), zap.Error(err))
			return
		}
	})
	this.chapterListCollector.OnError(func(response *colly.Response, e error) {
		utils.Logger.Error("爬取章节列表页, OnError", zap.Any("url", response.Request.URL.String()), zap.Error(e))

		//请求重试
		response.Request.Retry()
	})
}

/**
配置configChapterCollector
*/
func (this *spiderService) configChapterCollector() {
	if this.chapterCollector != nil {
		return
	}
	this.chapterCollector = this.NewCollector()

	// 爬取章节
	this.chapterCollector.OnHTML("div.mlfy_main", func(h *colly.HTMLElement) {
		chapterTitle := strings.TrimSpace(h.DOM.Find("h3.zhangj").Text())
		content, err := h.DOM.Find("div.read-content").Html()
		if err != nil {
			utils.Logger.Error("爬取章节, 解析内容, 异常", zap.Error(err))
			return
		}

		// 获取上下文信息
		novel := h.Response.Ctx.GetAny("novel").(*novel)
		// 累加爬取的章节页码
		novel.CurrentCrawChapterPageNo++

		chapter := &chapter{}
		chapter.Content = content
		chapter.Novel = novel
		chapter.Title = chapterTitle
		chapter.ChapterSrcUrl = h.Request.URL.String()
		chapter.Sort = novel.CurrentCrawChapterPageNo

		utils.Logger.Info("爬取章节, 完成", zap.Any("novelTitle", chapter.Novel.Title), zap.Any("chapterTitle", chapter.Title), zap.Any("novelSrcUrl", chapter.Novel.NovelSrcUrl), zap.Any("chapterSrcUrl", chapter.ChapterSrcUrl), zap.Any("chapter", chapter))
	})
	//通过翻页按钮爬取下一章
	this.chapterCollector.OnHTML("p.mlfy_page a:contains(下一章)", func(h *colly.HTMLElement) {
		nextChapterUrl, exist := h.DOM.Attr("href")
		if !exist {
			utils.Logger.Error("爬取下一章, 抽取下一页url， 异常", zap.Any("currentPage", h.Request.URL.String()))
			return
		}

		utils.Logger.Info("爬取下一章, 开始爬取", zap.Any("currentPage", h.Request.URL.String()), zap.Any("nextChapterUrl", nextChapterUrl))
		if err := this.chapterCollector.Request("GET", nextChapterUrl, nil, h.Response.Ctx, nil); err != nil {
			utils.Logger.Error("爬取下一章, 异常", zap.Any("currentPage", h.Request.URL.String()), zap.Any("nextChapterUrl", nextChapterUrl))
			return
		}
	})
	this.chapterCollector.OnError(func(response *colly.Response, e error) {
		utils.Logger.Error("爬取章节, OnError", zap.Any("url", response.Request.URL.String()), zap.Error(e))

		//请求重试
		response.Request.Retry()
	})
	this.chapterCollector.OnResponse(func(r *colly.Response) {
		filename := spiderGetFilename(r.Request.URL.String())
		filePath := library.DownloadFile(r.Request.URL.String(), filename)
		utils.Logger.Info("爬取章节, OnResponse, 保存文件", zap.Any("url", r.Request.URL.String()), zap.Any("filePath", filePath))
	})
}

func spiderGetFilename(url string) (filename string) {
	// 返回最后一个/的位置
	lastIndex := strings.LastIndex(url, "/")
	// 切出来
	filename = url[lastIndex+1:]

	// 时间戳解决重名
	if !strings.Contains(filename,".") {
		filename = strconv.Itoa(int(time.Now().UnixNano()))
	}
	return
}

/**
启动小说列表页爬取任务
*/
func (this *spiderService) StartCrawNovelListTask() error {
	// 初始化collector
	this.initCollector()

	if err := this.novelListCollector.Visit("https://www.51shu.com/sort_2"); err != nil {
		utils.Logger.Info("启动小说列表页爬取任务, 异常", zap.Error(err))
		return err
	}

	//若开启异步爬取模式, 则等待爬取线程执行完成
	if config.GlobalConfig.SpiderAsync {
		utils.Logger.Info("启动小说列表页爬取任务, 等待线程执行完成")
		this.novelListCollector.Wait()
	}

	utils.Logger.Info("启动小说列表页爬取任务, 完成")
	return nil
}
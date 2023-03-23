package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
)

func urllist(url string) ([]string, error) {
	var backurl []string
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	defer resp.Body.Close()
	fmt.Println(url)
	dom, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	dom.Find("a").Each(func(i int, selection *goquery.Selection) {
		href, ok := selection.Attr("href")
		if ok {
			link, err := resp.Request.URL.Parse(href)
			if err==nil{
				backurl = append(backurl, link.String())
			}
		}
	})
	return backurl, nil
}

func main() {
	firsturl := make(chan []string)
	defer close(firsturl)
	uniqurl := make(chan string)
	defer close(uniqurl)
	go func() { firsturl <- os.Args[1:] }()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range uniqurl {
				foundlink, err := urllist(link)
				if err != nil {
					fmt.Println(err)
				}
				if foundlink!=nil {
					go func() { firsturl <- foundlink }()
				}
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range firsturl {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				uniqurl <- link
			}
		}
	}
}

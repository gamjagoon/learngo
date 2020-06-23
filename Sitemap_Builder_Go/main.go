package main

import (
	"flag"
	"fmt"
	"io"
	"net/url"
	"sitemapbuilder/lib"
	"strings"
)

/*
 TODO:
 1. GET the webpage
 2. parse all the links on the page
 3. build proper urls with our links
 4. filter out any links w/ a diff domain
 5. Find allpages (BFS)
 6. print out XML
*/

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a stie map for")
	maxDepth := flag.Int("depth", 10,"the maximum number of links deep to traverse")
	flag.Parse()

	fmt.Println(*urlFlag)

	pages := bfs(*urlFlag,*maxDepth)
	for _, page := range pages{
		fmt.Println(page)
	}
}

// struct{} => 빈 구조체를 가리킨다
type empty struct{}
func bfs(urlStr string, maxDepth int)[]string {
	seen := make(map[string]empty)
	var q map[string]empty
	nq := map[string]empty{
		urlStr : empty{},
	}
	for i := 0; i < maxDepth ; i++ {
		q, nq = nq, make(map[string]empty)
		for url, _ := range q{
			if _, ok := seen[url]; ok {
				continue
			} 
			seen[url] = struct{}{}
			for _, link := range get(url){
				nq[link] = struct{}{}
			}
		}
	}
	ret := make([]string, 0, len(seen))
	for url, _ := range seen {
		ret = append(ret, url)
	}
	return ret
}

func get(urlFlag string)[]string {
	page := lib.GetWebpage(&urlFlag)
	defer page.Body.Close()
	
	reqUrl := page.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host: reqUrl.Host,
	}
	base := baseUrl.String()
	
	return filter(hrefs(page.Body,base),withPrefix(base))
}


func hrefs(r io.Reader,base string )[]string{
	links, _ := lib.Parse(r)
	var ret []string
	for _,i := range links {
		switch{
		case strings.HasPrefix(i.Href,"/"):
			ret = append(ret, base + i.Href)
		case strings.HasPrefix(i.Href,"http"):
			ret = append(ret, i.Href)
		}
	}

	return ret
}

func filter(links []string, keepFn func(string)bool)[]string {
	var ret []string
	for _, link := range links {
		// https://gopercise.com
		if keepFn(link){
			ret = append(ret, link)
		}
	}
	return ret
}

func withPrefix(pfx string) func(string)bool {
	return func (link string) bool {
		return strings.HasPrefix(link,pfx)
	}
}
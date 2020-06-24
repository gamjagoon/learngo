package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
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
const xmlns = "http//www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a stie map for")
	maxDepth := flag.Int("depth", 10, "the maximum number of links deep to traverse")
	flag.Parse()

	pages := bfs(*urlFlag, *maxDepth)

	toXML := urlset{
		Xmlns: xmlns,
	}
	for _, page := range pages {
		toXML.Urls = append(toXML.Urls, loc{page})
	}

	fmt.Print(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(toXML); err != nil {
		panic(err)
	}
}


func bfs(urlStr string, maxDepth int) []string {
	seen := make(map[string]struct{})
	var q map[string]struct{}
	nq := map[string]struct{}{
		urlStr: {},
	}
	for i := 0; i < maxDepth; i++ {
		q, nq = nq, make(map[string]struct{})
		for url := range q {
			if _, ok := seen[url]; ok {
				continue
			}
			seen[url] = struct{}{}
			for _, link := range get(url) {
				nq[link] = struct{}{}
			}
		}
	}
	ret := make([]string, 0, len(seen))
	for url := range seen {
		ret = append(ret, url)
	}
	return ret
}

func get(urlFlag string) []string {
	page := lib.GetWebpage(&urlFlag)
	defer page.Body.Close()

	reqURL := page.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	base := baseURL.String()

	return filter(hrefs(page.Body, base), withPrefix(base))
}

func hrefs(r io.Reader, base string) []string {
	links, _ := lib.Parse(r)
	var ret []string
	for _, i := range links {
		switch {
		case strings.HasPrefix(i.Href, "/"):
			ret = append(ret, base+i.Href)
		case strings.HasPrefix(i.Href, "http"):
			ret = append(ret, i.Href)
		}
	}

	return ret
}

func filter(links []string, keepFn func(string) bool) []string {
	var ret []string
	for _, link := range links {
		// https://gopercise.com
		if keepFn(link) {
			ret = append(ret, link)
		}
	}
	return ret
}

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}

package main

/*
 TODO:
 1. GET the webpage
 2. parse all the links on the page
 3. build proper urls with our links
 4. filter out any links w/ a diff domain
 5. Find allpages (BFS)
 6. print out XML
*/

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sitemapbuilder/lib"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a stie map for")
	flag.Parse()

	fmt.Println(*urlFlag)

	page := lib.GetWebpage(urlFlag)
	defer page.Body.Close()

	io.Copy(os.Stdout, page.Body)

}

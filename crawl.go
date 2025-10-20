package main

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"time"
)

type Downloaded struct {
	url  string
	body io.ReadCloser
}

type Extracted struct {
	url   string
	words []string
}

func manager(urlCh chan string, downloadableUrlCh chan string, visitedSet map[string]struct{}, rules *RobotsRules, fastMode bool) {
	fmt.Println("Manager started")

	for url := range urlCh {
		_, visited := visitedSet[url]
		if visited {
			continue
		}
		visitedSet[url] = struct{}{}
		if rules.Disallowed(url) {
			continue
		}
		if !fastMode {
			fmt.Printf("crawling %s\n", url)
		}
		downloadableUrlCh <- url
	}

	fmt.Println("Manager ended")
}

func sleeper(requestCh chan bool, readyCh chan bool, rules *RobotsRules) {
	fmt.Println("Sleeper started")

	for range requestCh {
		time.Sleep(rules.crawlDelay)
		readyCh <- true
	}

	fmt.Println("Sleeper ended")
}

func downloader(downloadableUrlCh chan string, bodyCh chan Downloaded, requestCh chan bool, readyCh chan bool, inProgCh chan bool) {
	fmt.Println("Downloader started")

	for url := range downloadableUrlCh {
		<-readyCh
		inProgCh <- true
		// fmt.Printf("Downloader: Downloading %s\n", url)
		body := download(url)
		requestCh <- true

		if body == nil {
			<-inProgCh
			continue
		}
		bodyCh <- Downloaded{url, body}
		<-inProgCh
	}

	fmt.Println("Downloader ended")
}

func extractor(bodyCh chan Downloaded, pageCh chan Extracted, dirtyUrlCh chan string) {
	fmt.Println("Extractor started")

	stopper := NewStopper()

	for downloaded := range bodyCh {
		url := downloaded.url
		body := downloaded.body

		reader := bufio.NewReader(body)
		words, hrefs := extract(reader, stopper)
		body.Close()

		if words != nil {
			page := Extracted{url, words}
			pageCh <- page
		}

		for _, href := range hrefs {
			dirtyUrlCh <- href
		}
	}

	fmt.Println("Extractor ended")
}

func builder(pageCh chan Extracted, idx *Index) {
	fmt.Println("Builder started")

	for page := range pageCh {
		if idx == nil {
			continue
		}
		url := page.url
		words := page.words
		for _, word := range words {
			(*idx).Increment(word, url)
		}
	}

	fmt.Println("Builder ended")
}

func adder(dirtyUrlCh chan string, urlCh chan string, host string) {
	fmt.Println("Adder started")

	for dirtyUrl := range dirtyUrlCh {
		url := cleanHref(host, dirtyUrl)
		if extractHost(url) == host {
			urlCh <- url
		}
	}

	fmt.Println("Adder ended")
}

func blockUntilChannelsEmpty(ch1 chan string, ch2 chan string, ch3 chan Downloaded, ch4 chan Extracted, ch5 chan string, print bool) {
	// time.Sleep(100 * time.Millisecond)
	for {
		if print {
			fmt.Printf("%d %d %d %d %d\n", len(ch1), len(ch2), len(ch3), len(ch4), len(ch5))
		}
		if len(ch1) == 0 && len(ch2) == 0 && len(ch3) == 0 && len(ch4) == 0 && len(ch5) == 0 {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func telemetry(visitedSet map[string]struct{}, ch1 chan string, ch2 chan string, ch3 chan Downloaded, ch4 chan Extracted, ch5 chan string, inProfCh chan bool) {
	for {
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("\n%d visited, channels: %d %d %d %d %d. in progress downloads: %d", len(visitedSet), len(ch1), len(ch2), len(ch3), len(ch4), len(ch5), len(inProfCh))
	}
}

/* Crawls the website starting from the given seed URL, then crawling all links found on that page, and so on for links found on those pages. fastMode ignores crawl-delay and prints less. If idx is not nil, crawl will also build the index using the index's Increment method */
func crawl(seed string, fastMode bool, idx *Index) {
	startTime := time.Now()

	visitedSet := make(map[string]struct{})
	host := extractHost(seed)
	if host == "" {
		return
	}

	robotsTxtUrl, err := url.JoinPath(host, "robots.txt")
	if err != nil {
		panic(err)
	}
	rules := parseRobotsTxt(robotsTxtUrl)
	if fastMode {
		rules.SetCrawlDelay(0)
	}

	urlCh := make(chan string, 10000)
	downloadableUrlCh := make(chan string, 100000)
	requestCh := make(chan bool, 10000)
	readyCh := make(chan bool, 10000)
	bodyCh := make(chan Downloaded, 10000)
	pageCh := make(chan Extracted, 10000)
	dirtyUrlCh := make(chan string, 10000)

	inProgCh := make(chan bool, 100000)

	// go telemetry(visitedSet, urlCh, downloadableUrlCh, bodyCh, pageCh, dirtyUrlCh, inProgCh)

	go manager(urlCh, downloadableUrlCh, visitedSet, rules, fastMode)
	go sleeper(requestCh, readyCh, rules)
	for i := 0; i < 10000; i++ {
		go downloader(downloadableUrlCh, bodyCh, requestCh, readyCh, inProgCh)
	}
	go extractor(bodyCh, pageCh, dirtyUrlCh)
	go builder(pageCh, idx)
	go adder(dirtyUrlCh, urlCh, host)

	urlCh <- seed
	readyCh <- true
	// blockUntilChannelsEmpty(urlCh, downloadableUrlCh, bodyCh, pageCh, dirtyUrlCh, true)
	time.Sleep(2 * time.Second)

	endTime := time.Now()
	duration := endTime.Sub(startTime).Seconds()
	numUrls := len(visitedSet)
	fmt.Printf("crawled %d urls in %.2f seconds (%.2f per second)\n", numUrls, duration, (float64)(numUrls)/duration)
}

/* Extracts the host from the given href */
func extractHost(href string) string {
	u, err := url.Parse(href)
	if err != nil {
		fmt.Printf("Error parsing href %q: %v\n", href, err)
		return ""
	}
	u.Path = "/"
	return u.String()
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"sync"
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

func manager(urlCh chan string, downloadableUrlCh chan string, visitedSet map[string]struct{}, rules *RobotsRules, fastMode bool, wg *sync.WaitGroup, inProgCh chan bool) {
	fmt.Println("Manager started")

	for url := range urlCh {
		_, visited := visitedSet[url]
		if visited {
			wg.Done()
			continue
		}
		visitedSet[url] = struct{}{}
		if rules.Disallowed(url) {
			wg.Done()
			continue
		}
		inProgCh <- true
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

func downloader(downloadableUrlCh chan string, bodyCh chan Downloaded, requestCh chan bool, readyCh chan bool) {
	fmt.Println("Downloader started")

	for url := range downloadableUrlCh {
		<-readyCh
		// fmt.Printf("Downloader: Downloading %s\n", url)
		body := download(url)
		requestCh <- true

		if body == nil {
			continue
		}
		bodyCh <- Downloaded{url, body}
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

func builder(pageCh chan Extracted, idx *Index, wg *sync.WaitGroup, inProgCh chan bool) {
	fmt.Println("Builder started")

	for page := range pageCh {
		if idx == nil {
			wg.Done()
			<-inProgCh
			continue
		}
		url := page.url
		words := page.words
		for _, word := range words {
			(*idx).Increment(word, url)
		}
		wg.Done()
		<-inProgCh
	}

	fmt.Println("Builder ended")
}

func adder(dirtyUrlCh chan string, urlCh chan string, host string, wg *sync.WaitGroup, inProgCh chan bool) {
	fmt.Println("Adder started")

	for dirtyUrl := range dirtyUrlCh {
		url := cleanHref(host, dirtyUrl)
		if extractHost(url) == host {
			urlCh <- url
			wg.Add(1)
		}
	}

	fmt.Println("Adder ended")
}

func telemetry(telemetryDoneCh chan bool, visitedSet map[string]struct{}, ch1 chan string, ch2 chan string, ch3 chan Downloaded, ch4 chan Extracted, ch5 chan string, inProfCh chan bool) {
	fmt.Println("Telemetry started")

	for {
		fmt.Printf("TELEMETRY: %d visited, channels: %d %d %d %d %d. in progress: %d\n", len(visitedSet), len(ch1), len(ch2), len(ch3), len(ch4), len(ch5), len(inProfCh))
		if len(telemetryDoneCh) > 0 {
			break
		}
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Telemetry ended")
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

	var wg sync.WaitGroup
	wg.Add(1)

	telemetryDoneCh := make(chan bool, 1)
	urlCh := make(chan string, 10000)
	downloadableUrlCh := make(chan string, 100000)
	requestCh := make(chan bool, 10000)
	readyCh := make(chan bool)
	bodyCh := make(chan Downloaded, 10000)
	pageCh := make(chan Extracted, 10000)
	dirtyUrlCh := make(chan string, 10000)
	inProgCh := make(chan bool, 100000)

	go telemetry(telemetryDoneCh, visitedSet, urlCh, downloadableUrlCh, bodyCh, pageCh, dirtyUrlCh, inProgCh)
	go manager(urlCh, downloadableUrlCh, visitedSet, rules, fastMode, &wg, inProgCh)
	go sleeper(requestCh, readyCh, rules)
	for i := 0; i < 1000; i++ {
		go downloader(downloadableUrlCh, bodyCh, requestCh, readyCh)
	}
	go extractor(bodyCh, pageCh, dirtyUrlCh)
	go builder(pageCh, idx, &wg, inProgCh)
	go adder(dirtyUrlCh, urlCh, host, &wg, inProgCh)

	urlCh <- seed
	readyCh <- true
	wg.Wait()

	endTime := time.Now()
	duration := endTime.Sub(startTime).Seconds()
	numUrls := len(visitedSet)
	fmt.Printf("crawled %d urls in %.2f seconds (%.2f per second)\n", numUrls, duration, (float64)(numUrls)/duration)

	telemetryDoneCh <- true
	close(telemetryDoneCh)
	close(urlCh)
	close(downloadableUrlCh)
	close(requestCh)
	close(bodyCh)
	close(pageCh)
	close(dirtyUrlCh)
	close(inProgCh)
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

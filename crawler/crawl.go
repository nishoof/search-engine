package crawler

import (
	"bufio"
	"io"
	"log/slog"
	"net/url"
	"sync"
	"time"

	"github.com/nishoof/search-engine/index"
)

type Downloaded struct {
	url  string
	body io.ReadCloser
}

type Extracted struct {
	url   string
	words map[string]int
	title string
}

func manager(urlCh chan string, downloadableUrlCh chan string, visitedSet map[string]struct{}, rules *RobotsRules, wg *sync.WaitGroup) {
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
		downloadableUrlCh <- url
	}
}

func sleeper(requestCh chan bool, readyCh chan bool, rules *RobotsRules) {
	for range requestCh {
		time.Sleep(rules.crawlDelay)
		readyCh <- true
	}
}

func downloader(downloadableUrlCh chan string, bodyCh chan Downloaded, requestCh chan bool, readyCh chan bool, wg *sync.WaitGroup) {
	for url := range downloadableUrlCh {
		<-readyCh
		slog.Debug("Downloading", "url", url)
		body := download(url)
		requestCh <- true

		if body == nil {
			wg.Done()
			continue
		}
		bodyCh <- Downloaded{url, body}
	}
}

func extractor(bodyCh chan Downloaded, pageCh chan Extracted, dirtyUrlCh chan string) {
	stopper := NewStopper()

	for downloaded := range bodyCh {
		url := downloaded.url
		body := downloaded.body

		reader := bufio.NewReader(body)
		words, hrefs, title := extract(reader, stopper)
		body.Close()

		if words != nil {
			page := Extracted{url, words, title}
			pageCh <- page
		}

		for _, href := range hrefs {
			dirtyUrlCh <- href
		}
	}
}

func builder(pageCh chan Extracted, idx *index.Index, wg *sync.WaitGroup) {
	for page := range pageCh {
		if idx == nil {
			wg.Done()
			continue
		}
		(*idx).AddDoc(page.url, page.title)
		for word, count := range page.words {
			(*idx).Increment(word, page.url, count)
		}
		wg.Done()
	}
}

func adder(dirtyUrlCh chan string, urlCh chan string, host string, wg *sync.WaitGroup) {
	for dirtyUrl := range dirtyUrlCh {
		url := cleanHref(host, dirtyUrl)
		if extractHost(url) == host {
			urlCh <- url
			wg.Add(1)
		}
	}
}

/* Crawls the website starting from the given seed URL, then crawling all links found on that page, and so on for links found on those pages. fastMode ignores Crawl-delay. If idx is not nil, Crawl will also build the index using the index's Increment method */
func Crawl(seed string, fastMode bool, idx *index.Index) {
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

	go manager(urlCh, downloadableUrlCh, visitedSet, rules, &wg)
	go sleeper(requestCh, readyCh, rules)
	for i := 0; i < 1000; i++ {
		go downloader(downloadableUrlCh, bodyCh, requestCh, readyCh, &wg)
	}
	go extractor(bodyCh, pageCh, dirtyUrlCh)
	go builder(pageCh, idx, &wg)
	go adder(dirtyUrlCh, urlCh, host, &wg)

	urlCh <- seed
	readyCh <- true
	wg.Wait()

	if idx != nil {
		(*idx).Flush()
	}

	telemetryDoneCh <- true
	close(telemetryDoneCh)
	close(urlCh)
	close(downloadableUrlCh)
	close(requestCh)
	close(bodyCh)
	close(pageCh)
	close(dirtyUrlCh)
}

/* Extracts the host from the given href */
func extractHost(href string) string {
	u, err := url.Parse(href)
	if err != nil {
		slog.Error("Error parsing href", "href", href, "error", err)
		return ""
	}
	u.Path = "/"
	return u.String()
}

package main

// handles the robots.txt file

import (
	"bufio"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type RobotsRules struct {
	disallowList []string
	crawlDelay   time.Duration
}

func (rules RobotsRules) Disallowed(url string) bool {
	for _, pattern := range rules.disallowList {
		matched, err := regexp.MatchString(pattern, url)
		if err != nil {
			panic(err)
		}
		if matched {
			return true
		}
	}
	return false
}

func (rules *RobotsRules) SetCrawlDelay(delaySeconds float64) {
	rules.crawlDelay = time.Duration(delaySeconds * (float64)(time.Second))
}

func parseRobotsTxt(robotsTxtUrl string) *RobotsRules {
	// make the RobotRules struct
	rules := new(RobotsRules)
	rules.SetCrawlDelay(0.5) // default crawl delay of 0.5 seconds (at most 2 requests per second)

	// try do download robots.txt
	resp, err := http.Get(robotsTxtUrl)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == 404 {
		fmt.Printf("%s not found\n", robotsTxtUrl)
		return rules
	}
	body := resp.Body
	if body == nil {
		panic("unable to download " + robotsTxtUrl)
	}
	defer body.Close()

	// call helper function on each line
	applies := false // if the current line applies to me (based on user agent)
	scanner := bufio.NewScanner(body)
	for scanner.Scan() {
		line := scanner.Text()
		parseRobotsTxtLine(line, rules, &applies)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return rules
}

func parseRobotsTxtLine(line string, rules *RobotsRules, applies *bool) {
	if len(line) == 0 {
		return // skip empty lines
	}
	if line[0] == '#' {
		return // skip comments
	}

	// extract key and value
	split := strings.Split(line, ": ")
	if len(split) != 2 {
		panic("unknown line format: " + line)
	}
	key := split[0]
	value := split[1]

	// handle user agent
	if key == "User-agent" {
		pattern := fixPattern(value)
		*applies = isMyUserAgent(pattern)
		return
	}

	// if the current line doesn't apply to use (applies to a different user agent), then skip
	if !*applies {
		return
	}

	// handle everything else
	switch key {
	case "Disallow":
		pattern := fixPattern(value)
		rules.disallowList = append(rules.disallowList, pattern)
	case "Crawl-delay":
		delaySeconds, err := strconv.ParseFloat(value, 64)
		if err != nil {
			panic("Invalid robots.txt. Crawl-delay must be a number, but got " + value + ". " + err.Error())
		}
		rules.SetCrawlDelay(delaySeconds)
	default:
		fmt.Printf("Skipping unsupported robots.txt line: %s\n", line)
	}
}

func fixPattern(str string) string {
	return strings.ReplaceAll(str, "*", ".*")
}

func isMyUserAgent(pattern string) bool {
	const MY_USER_AGENT = "Go-http-client/1.1"
	matched, err := regexp.MatchString(pattern, MY_USER_AGENT)
	if err != nil {
		panic(err)
	}
	return matched
}

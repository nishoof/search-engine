package main

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/nishoof/search-engine/crawler"
	"github.com/nishoof/search-engine/index"
	"github.com/nishoof/search-engine/testutils"
)

var idx index.Index
var simpleTestServer *httptest.Server
var top10TestServer *httptest.Server

func TestMain(m *testing.M) {
	simpleTestServer = testutils.NewSimpleTestServer()
	defer simpleTestServer.Close()

	top10TestServer = testutils.NewTop10TestServer()
	defer top10TestServer.Close()

	idx = index.NewIndexInMemory()
	crawler.Crawl(top10TestServer.URL+"/top10", true, &idx)

	os.Exit(m.Run())
}

package main

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/nishoof/search-engine/index"
)

var idx index.Index
var simpleTestServer *httptest.Server
var top10TestServer *httptest.Server

func TestMain(m *testing.M) {
	simpleTestServer = getSimpleTestServer()
	defer simpleTestServer.Close()

	top10TestServer = getTop10TestServer()
	defer top10TestServer.Close()

	idx = index.NewIndexInMemory()
	crawl(top10TestServer.URL+"/top10", true, &idx)

	os.Exit(m.Run())
}

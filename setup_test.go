package main

import (
	"net/http/httptest"
	"os"
	"testing"
)

var idx Index
var simpleTestServer *httptest.Server
var top10TestServer *httptest.Server

func TestMain(m *testing.M) {
	simpleTestServer = getSimpleTestServer()
	defer simpleTestServer.Close()

	top10TestServer = getTop10TestServer()
	defer top10TestServer.Close()

	idx = NewIndexInMemory()
	crawl(top10TestServer.URL+"/top10", true, &idx)

	os.Exit(m.Run())
}

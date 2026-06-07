package crawler

import (
	"io"
	"log/slog"
	"net/http"
)

/* Downloads the contents of the website of the given url and returns a ReadCloser on it */
func download(url string) io.ReadCloser {
	resp, err := http.Get(url)
	if err != nil {
		slog.Error("Error downloading", "url", url, "error", err)
		return nil
	}
	return resp.Body
}

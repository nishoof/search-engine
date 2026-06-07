package crawler

import (
	"io"
	"log/slog"
	"net/http"
)

// Downloads the contents of the website of the given url and returns a
// ReadCloser on it. Handles errors by logging and returning nil.
func download(url string) io.ReadCloser {
	resp, err := http.Get(url)
	if err != nil {
		slog.Error("Error downloading", "url", url, "error", err)
		return nil
	}
	if resp.StatusCode == 404 {
		resp.Body.Close()
		slog.Warn("Page not found (404)", "url", url)
		return nil
	}
	if resp.StatusCode != 200 {
		resp.Body.Close()
		slog.Error("Error downloading (non-200 status code)", "url", url, "status_code", resp.StatusCode)
		return nil
	}
	return resp.Body
}

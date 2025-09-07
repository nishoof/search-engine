package main

import "testing"

func TestCleanHref(t *testing.T) {
	tests := []struct {
		host, href, want string
	}{
		{"https://usf-cs272-f25.github.io/", "/", "https://usf-cs272-f25.github.io/"},
		{"https://usf-cs272-f25.github.io/", "/help/", "https://usf-cs272-f25.github.io/help/"},
		{"https://usf-cs272-f25.github.io/", "/syllabus/", "https://usf-cs272-f25.github.io/syllabus/"},
		{"https://usf-cs272-f25.github.io/", "https://gobyexample.com/", "https://gobyexample.com/"},
	}

	for _, test := range tests {
		got := cleanHref(test.host, test.href)
		if got != test.want {
			t.Errorf("For base %q and href %q, got %q but wanted %q\n", test.host, test.href, got, test.want)
		}
	}
}

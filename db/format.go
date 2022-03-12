package db

import (
	"net/url"
	"strings"
)

const (
	TimestampFormat = "2006-01-02 15:04:05"
)

// begin prelude/url.ts

func Query(q map[string]string) (query string) {
	for key, value := range q {
		query += key + "=" + url.QueryEscape(value) + "&"
	}

	if n := len(query); n > 0 {
		query = query[:n-1]
	}
	return
}

func QueryArray(elems []string) string {
	return strings.Join(elems, ",")
}

func AppendQuery(url, query string) string {
	if !strings.ContainsRune(url, '?') {
		return url + "?" + query
	}
	if url[len(url)-1] != '?' {
		url += "&"
	}
	return url + query
}

// end prelude/url.ts

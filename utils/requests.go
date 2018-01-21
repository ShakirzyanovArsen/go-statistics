package utils

import (
	"net/http"
	"time"
)

const dateTimeLayout = "2006-01-02 15:04:05"

func MakeTimesFromQuery(r *http.Request) (from time.Time, to time.Time) {
	query := r.URL.Query()
	from, _ = time.Parse(dateTimeLayout, query.Get("from"))
	to, _ = time.Parse(dateTimeLayout, query.Get("to"))
	return from, to
}

func CheckTimesExistInQuery(r *http.Request) bool {
	query := r.URL.Query()
	if query.Get("from") != `` && query.Get("to") != `` {
		return true
	}
	return false
}

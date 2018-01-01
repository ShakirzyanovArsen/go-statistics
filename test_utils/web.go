package test_utils

import (
	"net/http"
	"slurm_statistics/utils"
	"net/http/httptest"
)

func TestRequest(method string, url string, handl http.HandlerFunc) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, url, nil)
	utils.CheckFatal(err)
	rr := httptest.NewRecorder()
	handl.ServeHTTP(rr, req)
	return rr
}

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckCount(t *testing.T) {
	// 创建一个请求
	req, err := http.NewRequest("GET", "/checkCount", nil)
	if err != nil {
		t.Fatal(err)
	}

	// 创建一个响应recorder(满足http.responsewriter)去记录响应
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(checkCount)

	// 我们的handlers满足http.Handler，所以我们能调用serveHTTP 方法
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("code: got %v want %v", status, 200)
	}
}


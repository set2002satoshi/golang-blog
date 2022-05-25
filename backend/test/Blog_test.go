package main


import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/set2002satoshi/golang-blog/router"
)


func TestMain(t *testing.T) {
	content := `{
		"name": "testくん",
		"title": "testTitle"
	}`
	ts := httptest.NewServer(SetUpRouter())
	defer ts.Close()
	resp, _ := http.PUT("api/", content)
	fmt.Println(resp)
}
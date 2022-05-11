package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	REQ_URL = "http://localhost:8080"
)

func main() {

	// fs_handler := http.FileServer(http.Dir("static"))

	// http.Handle("/", fs_handler)

	values := url.Values{}
	values.Add("fruit", "apple")

	resp, err := http.PostForm(REQ_URL+"/onlyPost", values)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

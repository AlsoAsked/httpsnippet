package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	url := "http://mockbin.com/har"

	payload := strings.NewReader("Hello World")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "text/plain")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
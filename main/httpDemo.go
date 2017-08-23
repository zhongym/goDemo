package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	var response, _ = http.Get("http://127.0.0.1:8080/hello?id=asdda")
	if response.StatusCode==200 {
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	}
}

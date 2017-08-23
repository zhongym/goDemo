package main

import (
	"net/http"
	"log"
	"encoding/json"
	"strconv"
	"io"
	"fmt"
)

type User struct {
	ID   int  `json:"id"`
	Name string `json:"name"`
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var idSt = req.URL.Query()["id"][0]
	var id, _ = strconv.Atoi(idSt)

	var user = User{id, "zym"}
	var body, _ = json.Marshal(user)
	fmt.Println(string(body))
	io.WriteString(w, string(body))
}
func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

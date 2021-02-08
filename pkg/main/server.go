package main

import (
	"encoding/json"
	"net/http"
)

func main()  {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		json.NewEncoder(writer).Encode("hello")
	})
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LineChart struct {
	Categories []string
	Sales      []int
	RD         []int
}

func lineHandler(w http.ResponseWriter, req *http.Request) {

	getProject := req.URL.Path[len("/line/"):]
	fmt.Println(getProject)
	switch req.Method {
	case "GET":

		line := LineChart{
			Sales:      []int{1, 2, 3, 4, 5, 6, 7, 8},
			Categories: []string{"One", "Two", "Three", "Foru", "Five", "Six", "Seven", "Eight"},
			RD:         []int{5, 6, 7, 8, 1, 2, 3, 4},
		}

		buf, err := json.Marshal(line)
		fmt.Println(line)
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(buf)
	default:
		w.WriteHeader(400)
	}
}

func main() {

	fmt.Println("Start Server :8003/line ")
	http.HandleFunc("/line/", lineHandler)
	http.ListenAndServe(":8003", nil)

}

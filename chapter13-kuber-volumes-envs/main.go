package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const PORT = 3000
const TXT_PATH = "./permanent-data/text.txt"

func main() {
	http.HandleFunc("/text", func(res http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			getHandler(res)
		case http.MethodPost:
			postHandler(res, req)
		default:
			res.WriteHeader(http.StatusBadRequest)
			res.Header().Add("Content-Type", "text/plain")
			res.Write([]byte("Unknown resource"))
		}
	})
	log.Printf("Server running on port :%d\n", PORT)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}

type RequestBody struct {
	Text string `json:"text"`
}

func getHandler(res http.ResponseWriter) {
	content, err := os.ReadFile(TXT_PATH)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(content)
}

func postHandler(res http.ResponseWriter, req *http.Request) {
	var reqBody RequestBody
	if err := json.NewDecoder(req.Body).Decode(&reqBody); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Header().Add("Content-Type", "text/plain")
		res.Write([]byte(err.Error()))
		return
	}

	file, err := os.OpenFile(TXT_PATH, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Header().Add("Content-Type", "text/plain")
		res.Write([]byte(err.Error()))
		return
	}

	defer file.Close()

	if _, err := file.Write([]byte(fmt.Sprintf("%s\n", reqBody.Text))); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Header().Add("Content-Type", "text/plain")
		res.Write([]byte(err.Error()))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write([]byte("ok"))
}

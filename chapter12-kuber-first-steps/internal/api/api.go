package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Api struct{}

func New() *Api {
	return &Api{}
}

func (*Api) Start() {
	port := 8080

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Content-Type", "text/html")
		html := `
			<h1>Hello from this Go app !!!</h1>
			<p>Try to send a request to /error and see what would happen</p>`
		_, _ = res.Write([]byte(html))
	})

	http.HandleFunc("/error", func(res http.ResponseWriter, req *http.Request) {
		os.Exit(1)
	})

	log.Printf("Server is running on port %d\n", port)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

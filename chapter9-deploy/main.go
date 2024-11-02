package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", http.StripPrefix("/", fs))

	port := ":8080"
	fmt.Printf("Server running on port  %s\n", port)
	fmt.Println(http.ListenAndServe(port, nil))
}

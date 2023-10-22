package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	log.Println("watcher is awake")

	resp, err := http.Get(`http://genie:8081`)
	if err != nil {
		log.Println("genie is not live")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	log.Printf("genie responded with %d : %s", resp.StatusCode, string(body))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})
	http.ListenAndServe(":8080", nil)
}

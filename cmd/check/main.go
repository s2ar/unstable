package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{Timeout: 30 * time.Second}

	for {
		time.Sleep(1 * time.Second)

		log.Println("request to unstable")
		r, err := client.Get("http://unstable:8083/api/team/top")
		if err != nil {
			fmt.Println(err)
			continue
		}
		_ = r.Body.Close()
	}
}

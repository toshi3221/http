package main

import (
	"fmt"
	"github.com/toshi3221/http"
)

func main() {
	client, _ := http.NewClient()
	body, error := client.Get("https://example.com/")

	if error != nil {
		fmt.Printf("not found.")
		return
	}

	fmt.Printf("body:\n\n%s\n", body)
}

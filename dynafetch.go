package main

import (
	"dynafetch/credentials"
	"dynafetch/fetchxml"
	"os"
	"strings"
)

func main() {
	println("Waiting for credentials...")
	ch := make(chan credentials.RequestData)
	go credentials.Recieve(ch)
	credentials := <-ch

	println("Credentials recieved!")

	path := os.Args[1]
	fetchBytes, _ := os.ReadFile(path)
	fetch := string(fetchBytes)

	println("Fetching entities...")

	entities := fetchxml.Execute(credentials, fetch)

	println("Entities fetched!")
	values := entities["value"].([]any)

	for _, entity := range values {
		parsed := entity.(map[string]any)

		for key, value := range parsed {
			if(strings.HasPrefix(key, "@")) {
				continue
			}

			switch v := value.(type) {
			case string:
				println(key, v)
			case float64:
				println(key, v)
			case bool:
				println(key, v)
			default:
				println(key, v)
			}
		}

		println("\n")
	}
}

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/matejkoncal/dynafetch/credentials"
	"github.com/matejkoncal/dynafetch/fetchxml"
	"github.com/matejkoncal/dynafetch/terminal"
	"github.com/matejkoncal/dynafetch/watch"
	webui "github.com/matejkoncal/dynafetch/web-ui"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	assets, _ := webui.Assets()

	fs := http.FileServer(http.FS(assets))
	http.Handle("/", http.StripPrefix("/", fs))
	go http.ListenAndServe(":2222", nil)
	openBrowser("http://localhost:2222")

	select {}

	filePath := os.Args[1]

	fileProvided := isFileProvided()

	if !fileProvided {
		fmt.Println("Please provide a valid file path")
		return
	}

	println("Waiting for credentials...")
	credentials := waitForCredentials()
	println("Credentials recieved!")

	channel := make(chan watch.FileEvent)
	go watch.WatchFile(filePath, channel)

	for {
		terminal.Clear()

		fetchBytes, err := os.ReadFile(filePath)
		if err != nil {
			println(err)
		}

		fetch := string(fetchBytes)

		entities, err := fetchxml.Execute(credentials, fetch)
		if err != nil {
			println(err)
		}

		parsed, err := parseEntities(entities)
		if err != nil {
			fmt.Println(err)
		}

		terminal.PrintEntities(parsed)

		<-channel
	}

}

func waitForCredentials() credentials.RequestData {
	ch := make(chan credentials.RequestData)
	go credentials.Recieve(ch)
	credentials := <-ch
	return credentials
}

func isFileProvided() bool {
	if len(os.Args) < 2 {
		return false
	}

	if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
		return false
	}

	return true
}

func parseEntities(entities []byte) ([]any, error) {
	var jsonData map[string]any

	json.Unmarshal(entities, &jsonData)

	if errMap, ok := jsonData["error"].(map[string]any); ok {
		if msg, ok := errMap["message"].(string); ok {
			return nil, errors.New(msg)
		}
	}

	return jsonData["value"].([]any), nil
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		println(err)
	}

}

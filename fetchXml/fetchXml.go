package fetchxml

import (
	"dynafetch/credentials"
	"dynafetch/metadata"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func Execute(credentials credentials.RequestData, fetchXml string) ([]byte, error) {
	params := url.Values{}
	params.Add("fetchXml", fetchXml)

	parsedURL, err := url.Parse(credentials.URL)

	entityName, err := getEntityName(fetchXml)

	if err != nil {
		return nil, err
	}

	parsedURL.RawQuery = params.Encode()

	colectionName, err := metadata.GetCollectionName(credentials, entityName)

	if err != nil {
		return nil, err
	}

	parsedURL.Path = fmt.Sprintf("/api/data/v9.1/%s", colectionName)

	req, err := http.NewRequest("GET", parsedURL.String(), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Cookie", credentials.Cookie)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func getEntityName(fetch string) (string, error) {
	decoder := xml.NewDecoder(strings.NewReader(fetch))

	for {
		tok, err := decoder.Token()

		if err != nil {
			return "", err
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

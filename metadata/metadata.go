package metadata

import (
	"dynafetch/credentials"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetCollectionName(credentials credentials.RequestData, entityName string) (string, error) {

	parsedURL, err := url.Parse(credentials.URL)

	if err != nil {
		return "", err
	}

	parsedURL.RawQuery = "$select=EntitySetName"

	parsedURL.Path = fmt.Sprintf("/api/data/v9.0/EntityDefinitions(LogicalName='%s')", entityName)

	req, _ := http.NewRequest("GET", parsedURL.String(), nil)
	req.Header.Set("Cookie", credentials.Cookie)

	resp, _ := http.DefaultClient.Do(req)
	respBody, _ := io.ReadAll(resp.Body)

	var jsonData map[string]any

	json.Unmarshal(respBody, &jsonData)

	if jsonData["EntitySetName"] == nil {
		return "", errors.New("EntitySetName not found")
	}

	return jsonData["EntitySetName"].(string), nil
}

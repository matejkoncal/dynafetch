package fetchxml

import (
	"dynafetch/credentials"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetCollectionName(credentials credentials.RequestData, entityName string) string {

	parsedURL, err := url.Parse(credentials.URL)
	parsedURL.RawQuery = "$select=EntitySetName"

	if err != nil {
		panic(err)
	}

	// /api/data/v9.2/EntityDefinitions(LogicalName='account')

	parsedURL.Path = fmt.Sprintf("/api/data/v9.0/EntityDefinitions(LogicalName='%s')", entityName)

	println((parsedURL.String()))

	req, _ := http.NewRequest("GET", parsedURL.String(), nil)
	req.Header.Set("Cookie", credentials.Cookie)

	println("Fetching data...")
	resp, _ := http.DefaultClient.Do(req)
	respBody, _ := io.ReadAll(resp.Body)

	var jsonData map[string]any

	json.Unmarshal(respBody, &jsonData)

	return jsonData["EntitySetName"].(string)
}

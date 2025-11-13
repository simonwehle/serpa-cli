package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"serpa-cli/internal/types"
)

func CreatePlaces(baseUrl, apiVersion string, dbCategories []types.Category, csvPlaces []types.Place) ([]types.Place, error) {
	createdPlaces := make([]types.Place, 0, len(csvPlaces))

	for _, place := range csvPlaces {
		for _, category := range dbCategories {
			if category.Name == place.CategoryName {
				place.CategoryID = category.CategoryID
				break
			}
		}
		jsonData, err := json.Marshal(place)
		if err != nil {
			return nil, fmt.Errorf("error marshalling place %v: %w", place.Name, err)
		}

		url := fmt.Sprintf("%s%s/place", baseUrl, apiVersion)

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			return nil, fmt.Errorf("Error creating request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("Error sending request: %v", err)
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("Error reading response body: %v", err)
		}

		if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("API returned non-success status: %d, body: %s", resp.StatusCode, string(body))
		}

		var createdPlace types.Place
		err = json.Unmarshal(body, &createdPlace)
		if err != nil {
			return nil, fmt.Errorf("Error unmarshalling response: %v", err)
		}

		createdPlaces = append(createdPlaces, createdPlace)
	}
	

	return createdPlaces, nil
}


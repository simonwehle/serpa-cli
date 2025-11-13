package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"serpa-cli/internal/types"
)

func CreateCategories(baseUrl, apiVersion string, csvCategories []types.Category) ([]types.Category, error) {
	createdCategories := make([]types.Category, 0, len(csvCategories))

	for _, category := range csvCategories {
		jsonData, err := json.Marshal(category)
		if err != nil {
			return nil, fmt.Errorf("error marshalling category %v: %w", category.Name, err)
		}

		url := fmt.Sprintf("%s%s/category", baseUrl, apiVersion)

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

		var createdCategory types.Category
		err = json.Unmarshal(body, &createdCategory)
		if err != nil {
			return nil, fmt.Errorf("Error unmarshalling response: %v", err)
		}

		createdCategories = append(createdCategories, createdCategory)
	}

	return createdCategories, nil
}

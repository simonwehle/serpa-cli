package api

import (
	"fmt"

	"serpa-cli/internal/types"
	"serpa-cli/internal/utils"
)

func CreateCategories(baseUrl, apiVersion string, csvCategories []types.Category) ([]types.Category, error) {
    createdCategories := make([]types.Category, 0, len(csvCategories))
    url := fmt.Sprintf("%s%s/category", baseUrl, apiVersion)
    for _, category := range csvCategories {
        createdCategory, err := utils.DoPostRequest[types.Category](url, category)
        if err != nil {
            return nil, fmt.Errorf("Error during post request: %w", err)
        }
        createdCategories = append(createdCategories, *createdCategory)
    }
    return createdCategories, nil
}

package api

import (
	"fmt"

	"serpa-cli/internal/types"
	"serpa-cli/internal/utils"
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
		url := fmt.Sprintf("%s%s/place", baseUrl, apiVersion)

		createdPlace, err := utils.DoPostRequest[types.Place](url, place)
		if err != nil {
			return nil, fmt.Errorf("Error during post request: %w", err)
		}

		createdPlaces = append(createdPlaces, *createdPlace)
	}

	return createdPlaces, nil
}


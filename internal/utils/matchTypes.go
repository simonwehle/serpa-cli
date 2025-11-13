package utils

import (
	"serpa-cli/internal/types"
)

func MatchPlaces(apiCategories []types.Category, csvPlaces []types.Place) []types.Place {
    for i, place := range csvPlaces {
        for _, category := range apiCategories {
            if category.Name == place.CategoryName {
                csvPlaces[i].CategoryID = category.CategoryID
                break
            }
        }
    }
    return csvPlaces
}

func MatchAssets(apiPlaces []types.Place, placeAssets []types.PlaceAssets) []types.PlaceAssets {
    var matchedAssets []types.PlaceAssets
    for _, asset := range placeAssets {
        for _, place := range apiPlaces {
            if place.Name == asset.PlaceName {
                asset.PlaceID = place.PlaceID
                matchedAssets = append(matchedAssets, asset)
                break
            }
        }
    }
    return matchedAssets
}
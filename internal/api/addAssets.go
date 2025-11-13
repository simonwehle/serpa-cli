package api

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"serpa-cli/internal/types"
)

func AddAssets(fullUrl string, matchedAssets []types.PlaceAssets) error {
	client := &http.Client{}

	for _, placeAssets := range matchedAssets {
		uploadUrl := fmt.Sprintf("%s/place/%d/assets", fullUrl, placeAssets.PlaceID)

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		for _, assetPath := range placeAssets.Assets {
			file, err := os.Open(assetPath)
			if err != nil {
				return fmt.Errorf("failed to open file %s: %w", assetPath, err)
			}

			part, err := writer.CreateFormFile("assets", filepath.Base(assetPath))
			if err != nil {
				file.Close()
				return fmt.Errorf("failed to create form file for %s: %w", assetPath, err)
			}

			_, err = io.Copy(part, file)
			file.Close()
			if err != nil {
				return fmt.Errorf("failed to copy file data for %s: %w", assetPath, err)
			}
		}

		err := writer.Close()
		if err != nil {
			return fmt.Errorf("failed to close multipart writer: %w", err)
		}

		req, err := http.NewRequest("POST", uploadUrl, body)
		if err != nil {
			return fmt.Errorf("failed to create POST request: %w", err)
		}
		req.Header.Set("Content-Type", writer.FormDataContentType())

		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("POST request failed: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("upload failed for place %d: status %s", placeAssets.PlaceID, resp.Status)
		}

		fmt.Printf("Uploaded %d assets for place ID %d successfully\n", len(placeAssets.Assets), placeAssets.PlaceID)
	}

	return nil
}

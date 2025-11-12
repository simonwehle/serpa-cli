package files

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"serpa-cli/internal/types"
)

func ReadCategoriesCSV(root, categoriesFile string) ([]types.Category, error) {
	path := filepath.Join(root, categoriesFile)
	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening categories file:", err)
		os.Exit(1)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading categories file:", err)
		os.Exit(1)
	}

	var categories []types.Category
	for _, rec := range records {
    if len(rec) < 3 {
		return nil, fmt.Errorf("File category.csv requires columns 'name,icon,color'")
    }
    category := types.Category{
        Name: strings.TrimSpace(rec[0]),
        Icon: strings.TrimSpace(rec[1]),
        Color: strings.TrimSpace(rec[2]),
    }
    categories = append(categories, category)
}

	return categories, nil
}
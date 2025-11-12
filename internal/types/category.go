package types

type Category struct {
	CategoryID int    `json:"category_id"` 
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	Color      string `json:"color"`
}
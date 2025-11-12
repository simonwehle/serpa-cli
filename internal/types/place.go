package types

type Place struct {
	PlaceID     int     `json:"place_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	CategoryID  int     `json:"category_id"`
}
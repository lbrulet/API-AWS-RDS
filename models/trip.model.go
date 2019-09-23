package models

// Trip struct model
type Trip struct {
	ID       int     `json:"id"`
	StartLat float32 `json:"start_lat"`
	StartLng float32 `json:"start_lng"`
	EndLat   float32 `json:"end_lat"`
	EndLng   float32 `json:"end_lng"`
	IDUser   int     `json:"id_user"`
}

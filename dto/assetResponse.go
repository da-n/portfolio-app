package dto

type AssetResponse struct {
	Id      int    `json:"id"`
	Isin    string `json:"isin"`
	Name    string `json:"name"`
	Percent int    `json:"percent"`
}

package dto

type AssetResponse struct {
	Id      int64  `json:"id"`
	Isin    string `json:"isin"`
	Name    string `json:"name"`
	Percent int64  `json:"percent"`
}

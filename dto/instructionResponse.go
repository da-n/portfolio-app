package dto

type InstructionResponse struct {
	Id              int    `json:"id"`
	OrderSheetId    int    `json:"orderSheetId"`
	InstructionType string `json:"instructionType"`
	Isin            string `json:"isin"`
	Amount          int    `json:"amount"`
	CurrencyCode    string `json:"currencyCode"`
	CreatedAt       string `json:"createdAt"`
}

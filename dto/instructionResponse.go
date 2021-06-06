package dto

type InstructionResponse struct {
	Id              int64  `json:"id"`
	OrderSheetId    int64  `json:"orderSheetId"`
	InstructionType string `json:"instructionType"`
	Amount          int64  `json:"amount"`
	CurrencyCode    string `json:"currencyCode"`
}

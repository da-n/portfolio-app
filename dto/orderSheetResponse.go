package dto

type OrderSheetResponse struct {
	Id                  int64                  `json:"id"`
	WithdrawalRequestId int64                  `json:"withdrawalRequestId"`
	Status              string                 `json:"status"`
	CreatedAt           string                 `json:"createdAt"`
	Instructions        *[]InstructionResponse `json:"instructions"`
}

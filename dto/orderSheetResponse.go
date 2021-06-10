package dto

type OrderSheetResponse struct {
	Id                  int                    `json:"id"`
	WithdrawalRequestId int                    `json:"withdrawalRequestId"`
	Status              string                 `json:"status"`
	CreatedAt           string                 `json:"createdAt"`
	Instructions        *[]InstructionResponse `json:"instructions"`
}

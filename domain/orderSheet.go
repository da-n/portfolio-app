package domain

import (
	"github.com/da-n/portfolio-app/dto"
)

const OrderSheetComplete = "complete"

type OrderSheet struct {
	Id                  int64  `db:"id"`
	AccountId           int64  `db:"account_id"`
	WithdrawalRequestId int64  `db:"withdrawal_request_id"`
	Status              string `db:"status"`
	CreatedAt           string `db:"created_at"`
}

// ToDto takes a OrderSheet and casts it to dto.OrderSheetResponse
func (w OrderSheet) ToDto() dto.OrderSheetResponse {
	return dto.OrderSheetResponse{
		Id:                  w.Id,
		AccountId:           w.AccountId,
		WithdrawalRequestId: w.WithdrawalRequestId,
		Status:              w.Status,
		CreatedAt:           w.CreatedAt,
	}
}

package domain

import (
	"github.com/da-n/portfolio-app/dto"
)

type WithdrawalRequest struct {
	Id        int64  `db:"id"`
	AccountId int64  `db:"account_id"`
	Amount    int64  `db:"amount"`
	CreatedAt string `db:"created_at"`
}

// ToDto takes a WithdrawalRequest and casts it to dto.WithdrawalRequestResponse
func (w WithdrawalRequest) ToDto() dto.WithdrawalRequestResponse {
	return dto.WithdrawalRequestResponse{
		Id:        w.Id,
		AccountId: w.AccountId,
		Amount:    w.Amount,
		CreatedAt: w.CreatedAt,
	}
}

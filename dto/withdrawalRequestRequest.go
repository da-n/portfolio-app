package dto

import (
	"fmt"
	"github.com/da-n/portfolio-app/errs"
	"strings"
)

type WithdrawalRequestRequest struct {
	AccountId int64 `json:"accountId"`
	Amount    int64 `json:"amount"`
}

func (r WithdrawalRequestRequest) Validate() *errs.AppError {
	errors := make([]string, 0)

	if r.AccountId < 0 {
		errors = append(errors, fmt.Sprintf("accountId is required"))
	}

	if r.Amount <= 0 {
		errors = append(errors, fmt.Sprintf("amount must be greater than 0"))
	}

	if len(errors) > 0 {
		return errs.NewValidationError(strings.Join(errors, ", "))
	}

	return nil
}

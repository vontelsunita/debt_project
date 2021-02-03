package models

import (
	"time"
)

//All structs rated to Debt
type Debts struct {
	Id  int  `json:"id"`
	Amount float32 `json:"amount"`
	IsInPaymentPlan bool `json:"is_in_payment_plan"`
	RemainingAmount float32 `json:"remaining_amount"`
	NextPaymentDueDate *time.Time `json:"next_payment_due_date"`
}


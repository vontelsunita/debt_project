package models

import (
	"time"
)

type PaymentsPlan struct {
	Id  int  `json:"id"`
	DebtId int `json:"debt_id"`
	AmountToPay  float32  `json:"amount_to_pay"`
	InstallmentFrequency string `json:"installment_frequency"`
	InstallmentAmount float32 `json:"installment_amount"`
	StartDateStr       string    `json:"start_date"`
	StartDate  time.Time
}


type Payments struct {
	PaymentPlanId  int  `json:"payment_plan_id"`
	Amount  float32  `json:"amount"`
	Date       string    `json:"date"`
}


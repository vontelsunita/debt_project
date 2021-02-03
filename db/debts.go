package db

import (
	"debt_project/models"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

func GetDebts() ([]models.Debts, error) {
	url := "https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api/debts"
	var payload []models.Debts
	res, err := http.Get(url)
	if err != nil {
		logrus.Errorf("Error getting Debts %v", err)
		return payload, err
	}
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		logrus.Errorf("Error reading response body for Debts %v", err)
		return payload, err
	}

	err = json.Unmarshal(body, &payload)
	if err != nil {
		logrus.Errorf("Error Unmarshalling response for Debts %v", err)
		return payload, err
	}
	return payload, err
}

func ProcessDebts() ([]models.Debts, error) {
	//Get all Debts
	debts, err := GetDebts()

	//Get all Payment Plans
	payPlans, err := GetPaymentPlans()

	//Accumulate all payments made for a PaymentPlan in map
	paymentMap, err := GetPayments()

	//Loop through all Payment Plans
	for _, pp := range payPlans {
		for p, d := range debts {
			if pp.DebtId == d.Id {
				//Debt is in payment plan
				debts[p].IsInPaymentPlan = true

				debtsCleared := false
				// Calculate Remaining amount after preexisting payments are applied
				amountPaid, ok := paymentMap[pp.Id]
				if ok {
					if pp.AmountToPay >= amountPaid {
						debts[p].RemainingAmount = pp.AmountToPay - amountPaid
						if debts[p].RemainingAmount <= 0 {
							logrus.Printf("All Debts cleared for DebtID %d", pp.DebtId)
							debtsCleared = true
						}
					}
				} else {
					debts[p].RemainingAmount = pp.AmountToPay
				}

				//Set the next payment date by using the payment plan start_date,
				var timeToAdd time.Duration
				if debtsCleared == false {
					//Set NextPaymentDueDate only if a debt is due
					if pp.InstallmentFrequency == "WEEKLY" {
						timeToAdd = time.Hour * 24 * 7
					} else if pp.InstallmentFrequency == "BI_WEEKLY" {
						timeToAdd = time.Hour * 24 * 14
					}
					//Timestamp is already in ISO 8601 (time.RFC3339) format. No need to convert
					//t, err :=time.Parse(pp.StartDate.Add(time.Hour*24*14).Format(time.RFC3339), pp.StartDate.String())
					t := pp.StartDate.Add(timeToAdd)
					debts[p].NextPaymentDueDate = &t
				}

			}
		}
	}
	for i, d := range debts {
		//If a debt is not in PaymentPlan set RemainingAmount to Amount Owed.
		if debts[i].IsInPaymentPlan == false {
			debts[i].RemainingAmount = d.Amount
		}
		b, err := json.Marshal(d)
		if err != nil {
			logrus.Errorf("Error Marshalling Debts %v", err)
		}
		logrus.Println(string(b))
	}
	return debts, err
}


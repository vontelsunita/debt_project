package tests

import (
     "debt_project/db"
     "testing"
)

func TestProcessDebts(t *testing.T) {
     //We know Debt ID 4 is not in payment plan and Debt ID 0 is paid off
     debts, err := db.ProcessDebts()
     if err != nil {
          t.Errorf("%v", err)
     } else {
          //Loop through debts
          for _,d := range debts {
               if d.Id == 4 && d.IsInPaymentPlan == true {
                    t.Errorf("DebtID %d is not supposed to be in Payment Plan", d.Id)
               }
               if d.Id == 0 && d.RemainingAmount > 0 {
                    t.Errorf("DebtID %d is cleared. RemainingAmount is invalid %f", d.Id, d.RemainingAmount)
               }
          }
     }
}

func TestGetDebts(t *testing.T) {
     debts, _ := db.GetDebts()
     numDebts := len(debts)
     if numDebts != 5 {
          t.Errorf("Num debts expecting 5 got %d", numDebts)
     }
}

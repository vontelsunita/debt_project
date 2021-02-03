package db

import (
	"debt_project/models"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

func GetPaymentPlans() ([]models.PaymentsPlan, error) {
	url := "https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api/payment_plans"
	var payload []models.PaymentsPlan
	res, err := http.Get(url)
	if err != nil {
		logrus.Errorf("Error getting PaymentsPlan %v", err)
		return payload, err
	}
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		logrus.Errorf("Error reading response body for PaymentsPlan %v", err)
		return payload, err
	}
	err = json.Unmarshal(body, &payload)
	if err != nil {
		logrus.Errorf("Error unmarshalling PaymentsPlan %v", err)
		return payload, err
	}
	const layout = "2006-01-02"
	for i, p := range payload {
		startDate, err := time.Parse(layout, p.StartDateStr)
		if err != nil {
			logrus.Errorf("Error Parsing startDate in PaymentsPlan ID %d %v", p.Id, err)
		} else {
			payload[i].StartDate = startDate
		}
	}
	return payload, err
}

func GetPayments() (map[int]float32, error) {
	url := "https://my-json-server.typicode.com/druska/trueaccord-mock-payments-api/payments"
	var payload []models.Payments
	//Map if PaymentPlanID and Total amount of payments made
	paymentMap := make(map[int]float32)

	res, err := http.Get(url)
	if err != nil {
		logrus.Errorf("Error getting Payments %v", err)
		return paymentMap, err
	}
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		logrus.Errorf("Error reading response body for Payments %v", err)
		return paymentMap, err
	}
	err = json.Unmarshal(body, &payload)

	//Create map of PaymentPlanId and total sum of payments
	for _, i := range payload {
		//Check if a payment exists for the PaymentPlanId
		amt, ok := paymentMap[i.PaymentPlanId]
		if ok {
			//If there is already a payment for a PaymentPlanID add the amounts
			paymentMap[i.PaymentPlanId] = amt + i.Amount
		} else {
			paymentMap[i.PaymentPlanId] = i.Amount
		}
	}
	return paymentMap, err
}


package api

import (
	"encoding/json"
	"fmt"
	"monitoring/pkg/http"
)

type Billing struct {
	Balance int    `json:"balance"`
	Bonus   int    `json:"bonus"`
	Status  string `json:"status"`
	Summ    int    `json:"summ"`
	Unpaid  int    `json:"unpaid"`
	UserId  int    `json:"user_id"`
}

func (api *VscaleAPI) GetBalance() float32 {
	result, err := http.PostRequest("GET", JoinURL(V1URL, "billing/balance"), []byte(""), map[string]string{
		"X-Token": api.ApiToken,
	})
	if err != nil {
		fmt.Printf("Exception: %s", err)
	}

	var billingResult Billing
	err = json.Unmarshal(result, &billingResult)
	if err != nil {
		fmt.Printf("Exception: %s", err)
	}

	return float32(billingResult.Balance) / 100.0
}

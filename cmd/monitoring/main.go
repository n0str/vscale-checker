package main

import (
	"fmt"
	"log"
	"monitoring/cmd"
	"monitoring/cmd/api"
	codex_notify "monitoring/pkg/codex-notify"
)

func main() {
	config := cmd.LoadConfig("config.yml")
	vscaleAPI := api.NewVscaleAPI(config.ApiToken)
	balance := vscaleAPI.GetBalance()
	log.Printf("%v", balance)

	if balance < config.Threshold {
		codex_notify.SendMessage(config.Notify.NotifyCodex.URL, fmt.Sprintf("⚠️💸 Running out of money on <b>vscale</b>: %d₽", int(balance)))
	}
}

package main

import (
	"flag"
	"fmt"
	"log"
	"monitoring/cmd"
	"monitoring/cmd/api"
	codex_notify "monitoring/pkg/codex-notify"
)

func main() {
	apiToken := flag.String("token", "", "Vscale API token")
	threshold := flag.Float64("threshold", 100.0, "Minimum sum to notify")
	url := flag.String("url", "", "CodeX Bot Notify URL")
	configFilename := flag.String("config", "", "Configuration file")

	var config = &cmd.Config{}
	flag.Parse()
	if *configFilename != "" {
		config = cmd.LoadConfig(*configFilename)
	} else {
		config.ApiToken = *apiToken
		config.Threshold = float32(*threshold)
		config.Notify.NotifyCodex = cmd.NotifyCodex{Enabled: true, URL: *url}
	}

	if config.ApiToken == "" {
		log.Fatalf("Usage: ./monitoring --token <API_TOKEN>")
	}

	vscaleAPI := api.NewVscaleAPI(config.ApiToken)
	balance := vscaleAPI.GetBalance()
	log.Printf("%v", balance)

	if balance < config.Threshold {
		codex_notify.SendMessage(config.Notify.NotifyCodex.URL, fmt.Sprintf("⚠️💸 Running out of money on <b>vscale</b>: %d₽", int(balance)))
	}
}

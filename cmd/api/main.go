package api

import (
	"log"
	"net/url"
	"path"
)

type VscaleAPI struct {
	ApiToken string
	Url      string
}

const V1URL = "https://api.vscale.io/v1"

func NewVscaleAPI(apiToken string) *VscaleAPI {
	var api VscaleAPI
	api.ApiToken = apiToken
	api.Url = V1URL
	return &api
}

func JoinURL(baseURL, methodURL string) string {
	u, err := url.Parse(baseURL)
	if err != nil {
		log.Fatalf("Incorrect URL: %s", err)
	}
	u.Path = path.Join(u.Path, methodURL)
	return u.String()
}

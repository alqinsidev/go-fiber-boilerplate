package model

type PingStatusResponse struct {
	Status string `json:"status"`
	Active bool   `json:"active"`
}

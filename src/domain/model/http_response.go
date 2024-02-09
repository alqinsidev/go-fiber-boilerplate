package model

type HttpResponse[T any] struct {
	Data    T      `json:"data"`
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

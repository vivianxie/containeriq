package handler

import (
	"encoding/json"
	"net/http"

	"gopkg.in/macaron.v1"
)

func GetPingV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{})

	return http.StatusOK, result
}

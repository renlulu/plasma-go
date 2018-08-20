package handler

import (
	"net/http"
	"encoding/json"
	"github.com/renlulu/plasma-go/cli"
)

func GetDepositBlock(w http.ResponseWriter, r *http.Request) {
	b := cli.GetDepositBlock()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(b)
}

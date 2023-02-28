package http

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
	"log"
	"os"
)

type PingResponse struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
	Parameters  struct {
		MigrateToChatId int `json:"migrate_to_chat_id"`
		RetryAfter      int `json:"retry_after"`
	} `json:"parameters"`
}

func CheckServerHealth(token string) bool {
	client := req.C()
	url := fmt.Sprintf("http://localhost:8081/%s/ping", token)
	res, err := client.R().Post(url)
	var response PingResponse
	_ = json.Unmarshal([]byte(res.String()), &response)
	if err != nil || !response.Ok {
		log.Println("Error during contacting TG Servers!")
		log.Println(response.Description)
		os.Exit(1)
		return false
	}
	return response.Ok
}

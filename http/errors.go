package http

import "encoding/json"

type ErrorResponse struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
	Parameters  struct {
		MigrateToChatID int `json:"migrate_to_chat_id"`
		RetryAfter      int `json:"retry_after"`
	} `json:"parameters"`
}

func unmarshalErrorResponse(res string) *ErrorResponse {
	var unmarshalled ErrorResponse
	err := json.Unmarshal([]byte(res), &unmarshalled)
	if err != nil {
		return nil
	}
	return &unmarshalled
}

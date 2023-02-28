package http

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
	"log"
	"strconv"
)

func GetChats(userToken string) []string {
	var chats []string
	client := req.C()
	url := fmt.Sprintf("http://localhost:8081/user%s/getChats", userToken)
	res, err := client.R().Post(url)
	if err != nil {
		log.Println("Error during GetChats Call:")
		log.Println(err)
		return nil
	}
	var getChatsResponse GetChatsResponse
	err = json.Unmarshal([]byte(res.String()), &getChatsResponse)
	if err != nil || !getChatsResponse.Ok {
		log.Println("Error during ConfirmLogin response unmarshaling or call ko")
		return nil
	}
	for _, chat := range getChatsResponse.Result {
		name := chat.FirstName + chat.LastName
		if len(name) < 1 {
			name = chat.Title
		}
		chats = append(chats, fmt.Sprintf("%s (%s) [%s]", name, chat.Type, strconv.Itoa(chat.Id)))
	}
	return chats
}

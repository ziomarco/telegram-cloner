package http

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
	"log"
)

type ChatListItem struct {
	Name string
	Type string
	Id   int
}

func GetChats(userToken string) []ChatListItem {
	var chats []ChatListItem
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
		chats = append(chats, ChatListItem{
			Name: name,
			Type: chat.Type,
			Id:   chat.Id,
		})
	}
	return chats
}

package database

import "gorm.io/gorm"

type ForwardedMessage struct {
	gorm.Model
	OriginalMessageId  int
	ForwardedMessageId int
	SourceChatId       int
	DestinationChatId  int
}

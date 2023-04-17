package database

import "gorm.io/gorm"

type HandledMessage struct {
	gorm.Model
	OriginalMessageId int
	HandledAt         int64
	MarshaledMessage  string
}

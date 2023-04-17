package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	database "telegramcloner/db/entities"
)

var dbGlobalInstance gorm.DB

const SqliteDbFilename = "telegramcloner.db"

func init() {
	log.Println("Init DB started")
	db, err := gorm.Open(sqlite.Open(SqliteDbFilename), &gorm.Config{})
	if err != nil {
		log.Panicf("Unable to connect to database: %s", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&database.ForwardedMessage{})
	if err != nil {
		log.Panicf("Unable to migrate database: %s", err)
	}
	dbGlobalInstance = *db
	log.Println("Init DB completed")
}

func RegisterNewForward(message database.ForwardedMessage) bool {
	result := dbGlobalInstance.Create(&message)
	if result.Error != nil {
		log.Printf("An error occurred while trying to track forwarded message %d in database. Message was forwarded but note that any replies to this message won't be forwarded correctly.", message.OriginalMessageId)
		log.Println(result.Error)
		return false
	}
	return result.RowsAffected > 0
}

func FindMessage(originalMessageId int) *database.ForwardedMessage {
	var message = &database.ForwardedMessage{
		OriginalMessageId: originalMessageId,
	}

	result := dbGlobalInstance.First(&message)

	if result.Error != nil {
		log.Println(result.Error)
		return nil
	}

	return message
}

func VerifyDatabase() {
	mock := database.ForwardedMessage{
		OriginalMessageId:  1,
		ForwardedMessageId: 2,
		SourceChatId:       3,
		DestinationChatId:  4,
	}

	res := RegisterNewForward(mock)

	if !res {
		log.Panicln("[InitDB] DB save failed")
	}

	found := FindMessage(mock.OriginalMessageId)

	if found.ForwardedMessageId != mock.ForwardedMessageId {
		log.Panicln("[InitDB] ForwardedMessageId mismatch")
	}
	if found.DestinationChatId != mock.DestinationChatId {
		log.Panicln("[InitDB] DestinationChatId mismatch")
	}
	if found.OriginalMessageId != mock.OriginalMessageId {
		log.Panicln("[InitDB] OriginalMessageId mismatch")
	}
	if found.SourceChatId != mock.SourceChatId {
		log.Panicln("[InitDB] SourceChatId mismatch")
	}
}

func RegisterHandledMessage(messageId int, marshaledMessage string) bool {
	now := time.Now()
	var message = &database.HandledMessage{
		OriginalMessageId: messageId,
		HandledAt:         now.Unix(),
		MarshaledMessage:  marshaledMessage,
	}
	result := dbGlobalInstance.Create(&message)
	if result.Error != nil {
		log.Printf("An error occurred while trying to track handled message %d in database.", messageId)
		log.Println(result.Error)
		return false
	}
	return result.RowsAffected > 0
}
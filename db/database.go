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

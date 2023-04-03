package database

import (
	database "telegramcloner/db/entities"
	"testing"
)

func TestFunction(t *testing.T) {
	mock := database.ForwardedMessage{
		OriginalMessageId:  1,
		ForwardedMessageId: 2,
		SourceChatId:       3,
		DestinationChatId:  4,
	}

	res := RegisterNewForward(mock)

	if !res {
		t.Error("DB save failed")
	}

	found := FindMessage(mock.OriginalMessageId)

	if found.ForwardedMessageId != mock.ForwardedMessageId {
		t.Error("ForwardedMessageId mismatch")
	}
	if found.DestinationChatId != mock.DestinationChatId {
		t.Error("DestinationChatId mismatch")
	}
	if found.OriginalMessageId != mock.OriginalMessageId {
		t.Error("OriginalMessageId mismatch")
	}
	if found.SourceChatId != mock.SourceChatId {
		t.Error("SourceChatId mismatch")
	}
}

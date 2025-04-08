package initializers

import "github.com/DevAthhh/xvibe-chat/internal/database"

func ConnectToDB() error {
	return database.Connect()
}

func SyncDB() error {
	return database.SyncDB()
}

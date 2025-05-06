package initializers

import "src/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})

}

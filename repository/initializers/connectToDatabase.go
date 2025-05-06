package initializers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"src/models"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	//dsn := os.Getenv("DATABASE_URL")
	dsn := "root:password@tcp(localhost:3306)/lexend_database?parseTime=true"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	//database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.User{})

	// ✅ Add the unique index for daily profit only once source = 'new daily profit'

	DB = database

}

package initializers

import "github.com/marcostota/gojwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}

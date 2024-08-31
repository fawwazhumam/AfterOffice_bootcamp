package migration

import (
	"fmt"
	"go_crud/database"
	"go_crud/model/user/entity"
	"log"
)

func RunMigration() {

	err := database.DB.AutoMigrate(&entity.User{})

	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated")
}
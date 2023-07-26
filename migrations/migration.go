package main

import (
	"github.com/osvaldodmvs/api/initializers"
	"github.com/osvaldodmvs/api/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Product{})
	initializers.DB.AutoMigrate(&models.User{})
}

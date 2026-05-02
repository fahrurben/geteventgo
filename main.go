package main

import (
	"log"

	"github.com/fahrurben/geteventgo/common"
	"github.com/fahrurben/geteventgo/events"
	"github.com/fahrurben/geteventgo/users"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&users.UserModel{})
	db.AutoMigrate(&events.Event{})
	db.AutoMigrate(&events.Pricing{})
}

func main() {
	db := common.Init()
	Migrate(db)

	sqlDB, err := db.DB()
	if err != nil {
		log.Println("failed to get sql.DB", err)
	} else {
		defer sqlDB.Close()
	}

	router := gin.Default()
	router.RedirectTrailingSlash = false

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("/api")
	users.UserEndpoints(v1)
	events.EventEndpoints(v1)

	router.Run()
}

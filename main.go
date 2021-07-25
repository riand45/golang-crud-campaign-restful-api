package main

import (
	"github.com/gin-gonic/gin"
	"github.com/riand45/golang-crud-campaign-restful-api/campaign"
	"github.com/riand45/golang-crud-campaign-restful-api/handler"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/campaign-service?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	campaignRepository := campaign.NewRepository(db)

	campaignService := campaign.NewService(campaignRepository)

	campaignHandler := handler.NewCampaignHandler(campaignService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.GET("/campaigns", campaignHandler.GetCampaigns)

	router.Run()
}

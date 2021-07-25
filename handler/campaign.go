package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/riand45/golang-crud-campaign-restful-api/campaign"
	"github.com/riand45/golang-crud-campaign-restful-api/helper"
	"net/http"
	"strconv"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, _ := h.service.FindCampaigns(userID)
	//if err != nil {
	//	response := helper.APIResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
	//	c.JSON(http.StatusBadRequest, response)
	//	return
	//}

	response := helper.APIResponse("List of campaigns", http.StatusOK, "success", campaigns)
	c.JSON(http.StatusOK, response)
}
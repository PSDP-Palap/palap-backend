package v1

import (
	"encoding/json"
	"net/http"
	"palap_backend/config"
	"palap_backend/models"

	"github.com/gin-gonic/gin"
)

// var services []models.Service = []models.Service{
// 	{ID: 1, Title: "โหลๆ 1234"},
// 	{ID: 2, Title: "โหลๆ 5678"},
// }

var supabaseClient = config.InitializeSupabase()

// @Summary Get all service
// @Produce  json
// @Success 200 {object} models.Service
// @Failure 404 {object} models.Error
// @Router /api/v1/service [get]
func GetServices(c *gin.Context) {
	var services []models.Service

	data, _, err := supabaseClient.From("services").Select("*", "", false).Execute()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to load data"})
		return
	}

	if err := json.Unmarshal(data, &services); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Unable to convert data"})
		return
	}
	c.JSON(http.StatusOK, services)
}

// func GetService(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))

// 	if err != nil {
// 		r := models.Error{Error: "Bad request"}
// 		c.JSON(http.StatusBadRequest, r)
// 		return
// 	}

// 	for _, service := range services {
// 		if service.ID == id {
// 			c.JSON(http.StatusOK, service)
// 			return
// 		}
// 	}

// 	r := models.Error{Error: "Not found"}
// 	c.JSON(http.StatusNotFound, r)
// }

// func CreateBook(c *gin.Context) {
// 	service := new(models.Service)

// 	if err := c.BindJSON(service); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter"})
// 		return
// 	}

// 	service.ID = len(services) + 1
// 	services = append(services, *service)

// 	c.JSON(http.StatusOK, service)
// }

// func UpdateService(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
// 		return
// 	}

// 	updatedService := new(models.Service)

// 	if err := c.BindJSON(updatedService); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter"})
// 		return
// 	}

// 	for i, service := range services {
// 		if service.ID == id {
// 			service.Title = updatedService.Title
// 			services[i] = service
// 			c.JSON(http.StatusOK, service)
// 			return
// 		}
// 	}

// 	c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
// }

// func DeleteService(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
// 	}

// 	for i, service := range services {
// 		if service.ID == id {
// 			services = append(services[:i], services[i+1:]...)
// 			c.JSON(http.StatusNoContent, gin.H{"message": "No content"})
// 			return
// 		}
// 	}

//c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
// }

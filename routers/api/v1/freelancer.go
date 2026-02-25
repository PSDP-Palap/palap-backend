package v1

import (
	"net/http"
	"strconv"

	"palap_backend/models"

	"github.com/gin-gonic/gin"
)

var freelancers []models.Freelancer = []models.Freelancer{
	{ID: 1, Name: "John Doe", Email: "john@example.com", Bio: "Expert in web development", Skills: "Go, React, Node.js", Rating: 5, Status: "active", CreatedAt: "2026-01-15"},
	{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Bio: "UI/UX Designer", Skills: "Figma, Adobe XD", Rating: 4, Status: "active", CreatedAt: "2026-02-01"},
}

var freelanceJobs []models.FreelanceJob = []models.FreelanceJob{
	{ID: 1, Title: "Build mobile app", Description: "Need to build an iOS app", Budget: 50000, Status: "open", CreatedAt: "2026-02-20"},
	{ID: 2, Title: "Web design", Description: "Design new website", Budget: 30000, Status: "open", CreatedAt: "2026-02-22"},
}

var productServices []models.ProductService = []models.ProductService{
	{ID: 1, Name: "Web Development", Description: "Custom web development services", Price: 10000, Category: "Software", Status: "active"},
	{ID: 2, Name: "Mobile App Development", Description: "iOS and Android app development", Price: 50000, Category: "Software", Status: "active"},
}

var shopLocations []models.ShopLocation = []models.ShopLocation{
	{ID: 1, Name: "Shop 1", Address: "123 Main St", Phone: "08x-xxx-xxxx", City: "Bangkok", Country: "Thailand", Status: "active"},
	{ID: 2, Name: "Shop 2", Address: "456 Sukhumvit Rd", Phone: "08x-xxx-xxxx", City: "Bangkok", Country: "Thailand", Status: "active"},
}

// Freelancer Registration
// @Summary Freelancer Registration
// @Description Register new freelancer
// @Accept json
// @Produce json
// @Param freelancer body models.Freelancer true "Freelancer details"
// @Success 201 {object} models.Freelancer
// @Failure 400 {object} models.Error
// @Router /api/v1/registration [post]
func RegisterFreelancer(c *gin.Context) {
	freelancer := new(models.Freelancer)
	if err := c.BindJSON(freelancer); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: "Invalid request body"})
		return
	}

	freelancer.ID = len(freelancers) + 1
	freelancer.Status = "active"
	freelancers = append(freelancers, *freelancer)

	c.JSON(http.StatusCreated, freelancer)
}

// @Summary Get Freelance Jobs
// @Description Get list of available freelance jobs
// @Produce json
// @Success 200 {array} models.FreelanceJob
// @Router /api/v1/get_freelance_job [get]
func GetFreelanceJobs(c *gin.Context) {
	c.JSON(http.StatusOK, freelanceJobs)
}

// @Summary Get Freelancers
// @Description Get list of all freelancers
// @Produce json
// @Success 200 {array} models.Freelancer
// @Router /api/v1/manage_freelance [get]
func GetFreelancers(c *gin.Context) {
	c.JSON(http.StatusOK, freelancers)
}

// @Summary Get Freelancer by ID
// @Description Get freelancer details by ID
// @Produce json
// @Param id path int true "Freelancer ID"
// @Success 200 {object} models.Freelancer
// @Failure 404 {object} models.Error
// @Router /api/v1/manage_freelance/:id [get]
func GetFreelancer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: "Invalid freelancer ID"})
		return
	}

	for _, freelancer := range freelancers {
		if freelancer.ID == id {
			c.JSON(http.StatusOK, freelancer)
			return
		}
	}

	c.JSON(http.StatusNotFound, models.Error{Error: "Freelancer not found"})
}

// @Summary Create Freelancer
// @Description Create new freelancer (Admin only)
// @Accept json
// @Produce json
// @Param freelancer body models.Freelancer true "Freelancer details"
// @Success 201 {object} models.Freelancer
// @Failure 400 {object} models.Error
// @Router /api/v1/manage_freelance [post]
func CreateFreelancer(c *gin.Context) {
	freelancer := new(models.Freelancer)
	if err := c.BindJSON(freelancer); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: "Invalid request body"})
		return
	}

	freelancer.ID = len(freelancers) + 1
	freelancers = append(freelancers, *freelancer)

	c.JSON(http.StatusCreated, freelancer)
}

// @Summary Update Freelancer
// @Description Update existing freelancer
// @Accept json
// @Produce json
// @Param id path int true "Freelancer ID"
// @Param freelancer body models.Freelancer true "Updated freelancer details"
// @Success 200 {object} models.Freelancer
// @Failure 404 {object} models.Error
// @Router /api/v1/manage_freelance/:id [put]
func UpdateFreelancer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: "Invalid freelancer ID"})
		return
	}

	updatedFreelancer := new(models.Freelancer)
	if err := c.BindJSON(updatedFreelancer); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: "Invalid request body"})
		return
	}

	for i, freelancer := range freelancers {
		if freelancer.ID == id {
			updatedFreelancer.ID = id
			freelancers[i] = *updatedFreelancer
			c.JSON(http.StatusOK, updatedFreelancer)
			return
		}
	}

	c.JSON(http.StatusNotFound, models.Error{Error: "Freelancer not found"})
}

// @Summary Delete Freelancer
// @Description Delete freelancer by ID
// @Param id path int true "Freelancer ID"
// @Success 204
// @Failure 404 {object} models.Error
// @Router /api/v1/manage_freelance/:id [delete]
func DeleteFreelancer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: "Invalid freelancer ID"})
		return
	}

	for i, freelancer := range freelancers {
		if freelancer.ID == id {
			freelancers = append(freelancers[:i], freelancers[i+1:]...)
			c.JSON(http.StatusNoContent, nil)
			return
		}
	}

	c.JSON(http.StatusNotFound, models.Error{Error: "Freelancer not found"})
}

// Product and Service Management

// @Summary Get Products/Services
// @Description Get list of all products and services
// @Produce json
// @Success 200 {array} models.ProductService
// @Router /api/v1/manage_product_and_service [get]
func GetProductServices(c *gin.Context) {
	c.JSON(http.StatusOK, productServices)
}

// @Summary Get Product/Service by ID
// @Description Get product/service details by ID
// @Produce json
// @Param id path int true "Product/Service ID"
// @Success 200 {object} models.ProductService
// @Failure 404 {object} models.Error
// @Router /api/v1/manage_product_and_service/:id [get]
func GetProductService(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: "Invalid product/service ID"})
		return
	}

	for _, ps := range productServices {
		if ps.ID == id {
			c.JSON(http.StatusOK, ps)
			return
		}
	}

	c.JSON(http.StatusNotFound, models.Error{Error: "Product/Service not found"})
}

// @Summary Create Product/Service
// @Description Create new product or service
// @Accept json
// @Produce json
// @Param product body models.ProductService true "Product/Service details"
// @Success 201 {object} models.ProductService
// @Failure 400 {object} models.Error
// @Router /api/v1/manage_product_and_service [post]
func CreateProductService(c *gin.Context) {
	ps := new(models.ProductService)
	if err := c.BindJSON(ps); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: "Invalid request body"})
		return
	}

	ps.ID = len(productServices) + 1
	productServices = append(productServices, *ps)

	c.JSON(http.StatusCreated, ps)
}

// @Summary Update Product/Service
// @Description Update existing product/service
// @Accept json
// @Produce json
// @Param id path int true "Product/Service ID"
// @Param product body models.ProductService true "Updated product/service details"
// @Success 200 {object} models.ProductService
// @Failure 404 {object} models.Error
// @Router /api/v1/manage_product_and_service/:id [put]
func UpdateProductService(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: "Invalid product/service ID"})
		return
	}

	updatedPS := new(models.ProductService)
	if err := c.BindJSON(updatedPS); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: "Invalid request body"})
		return
	}

	for i, ps := range productServices {
		if ps.ID == id {
			updatedPS.ID = id
			productServices[i] = *updatedPS
			c.JSON(http.StatusOK, updatedPS)
			return
		}
	}

	c.JSON(http.StatusNotFound, models.Error{Error: "Product/Service not found"})
}

// @Summary Delete Product/Service
// @Description Delete product/service by ID
// @Param id path int true "Product/Service ID"
// @Success 204
// @Failure 404 {object} models.Error
// @Router /api/v1/manage_product_and_service/:id [delete]
func DeleteProductService(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: "Invalid product/service ID"})
		return
	}

	for i, ps := range productServices {
		if ps.ID == id {
			productServices = append(productServices[:i], productServices[i+1:]...)
			c.JSON(http.StatusNoContent, nil)
			return
		}
	}

	c.JSON(http.StatusNotFound, models.Error{Error: "Product/Service not found"})
}

// Shop Location Management

// @Summary Get Shop Locations
// @Description Get list of all shop locations
// @Produce json
// @Success 200 {array} models.ShopLocation
// @Router /api/v1/manage_shop_location [get]
func GetShopLocations(c *gin.Context) {
	c.JSON(http.StatusOK, shopLocations)
}

// @Summary Get Shop Location by ID
// @Description Get shop location details by ID
// @Produce json
// @Param id path int true "Shop Location ID"
// @Success 200 {object} models.ShopLocation
// @Failure 404 {object} models.Error
// @Router /api/v1/manage_shop_location/:id [get]
func GetShopLocation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: "Invalid shop location ID"})
		return
	}

	for _, location := range shopLocations {
		if location.ID == id {
			c.JSON(http.StatusOK, location)
			return
		}
	}

	c.JSON(http.StatusNotFound, models.Error{Error: "Shop location not found"})
}

// @Summary Create Shop Location
// @Description Create new shop location
// @Accept json
// @Produce json
// @Param location body models.ShopLocation true "Shop location details"
// @Success 201 {object} models.ShopLocation
// @Failure 400 {object} models.Error
// @Router /api/v1/manage_shop_location [post]
func CreateShopLocation(c *gin.Context) {
	location := new(models.ShopLocation)
	if err := c.BindJSON(location); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: "Invalid request body"})
		return
	}

	location.ID = len(shopLocations) + 1
	shopLocations = append(shopLocations, *location)

	c.JSON(http.StatusCreated, location)
}

// @Summary Update Shop Location
// @Description Update existing shop location
// @Accept json
// @Produce json
// @Param id path int true "Shop Location ID"
// @Param location body models.ShopLocation true "Updated shop location details"
// @Success 200 {object} models.ShopLocation
// @Failure 404 {object} models.Error
// @Router /api/v1/manage_shop_location/:id [put]
func UpdateShopLocation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: "Invalid shop location ID"})
		return
	}

	updatedLocation := new(models.ShopLocation)
	if err := c.BindJSON(updatedLocation); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: "Invalid request body"})
		return
	}

	for i, location := range shopLocations {
		if location.ID == id {
			updatedLocation.ID = id
			shopLocations[i] = *updatedLocation
			c.JSON(http.StatusOK, updatedLocation)
			return
		}
	}

	c.JSON(http.StatusNotFound, models.Error{Error: "Shop location not found"})
}

// @Summary Delete Shop Location
// @Description Delete shop location by ID
// @Param id path int true "Shop Location ID"
// @Success 204
// @Failure 404 {object} models.Error
// @Router /api/v1/manage_shop_location/:id [delete]
func DeleteShopLocation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: "Invalid shop location ID"})
		return
	}

	for i, location := range shopLocations {
		if location.ID == id {
			shopLocations = append(shopLocations[:i], shopLocations[i+1:]...)
			c.JSON(http.StatusNoContent, nil)
			return
		}
	}

	c.JSON(http.StatusNotFound, models.Error{Error: "Shop location not found"})
}

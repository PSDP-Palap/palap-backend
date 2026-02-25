package v1

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"palap_backend/models"

	"github.com/gin-gonic/gin"
)

var orders []models.Order = []models.Order{
	{ID: 1, UserID: 1, ServiceID: 1, AddressID: 1, Status: "Completed", PaymentMethod: "card", CreatedAt: "2026-02-25"},
	{ID: 2, UserID: 1, ServiceID: 2, AddressID: 2, Status: "Pending", PaymentMethod: "qr_code", QRCode: "https://example.com/qr/2", CreatedAt: "2026-02-25"},
}

var orderDetails []models.OrderDetail = []models.OrderDetail{
	{ID: 1, OrderID: 1, Description: "Service Description 1", Amount: 1000},
	{ID: 2, OrderID: 1, Description: "Service Description 2", Amount: 500},
	{ID: 3, OrderID: 2, Description: "Service Description 3", Amount: 1500},
}

// @Summary Payment
// @Description Process payment with card or QR code
// @Accept json
// @Produce json
// @Param Authorization header string true "User Token"
// @Param payment body models.PaymentRequest true "Payment details"
// @Success 200 {object} models.PaymentResponse
// @Failure 400 {object} models.Error
// @Router /api/v1/payment [post]
func Payment(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.Error{Error: "Missing authorization token"})
		return
	}

	var payment models.PaymentRequest
	if err := c.BindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: "Invalid request body"})
		return
	}

	// Validate payment method
	if payment.PaymentMethod != "card" && payment.PaymentMethod != "qr_code" {
		c.JSON(http.StatusBadRequest, models.Error{Error: "Invalid payment method"})
		return
	}

	// Process card payment
	if payment.PaymentMethod == "card" {
		if payment.PaymentCard == nil {
			c.JSON(http.StatusBadRequest, models.Error{Error: "Card details required"})
			return
		}
		response := models.PaymentResponse{Status: "OK"}
		c.JSON(http.StatusOK, response)
		return
	}

	// Generate QR code for qr_code payment
	qrCode := fmt.Sprintf("https://api.qr-server.com/v1/qr?size=200x200&data=order_%d_%d", payment.ServiceID, time.Now().UnixMilli())
	response := models.PaymentResponse{Status: "OK", QRCode: qrCode}
	c.JSON(http.StatusOK, response)
}

// @Summary Get Orders
// @Description Get all orders for authenticated user
// @Produce json
// @Param Authorization header string true "User Token"
// @Success 200 {array} models.Order
// @Failure 401 {object} models.Error
// @Router /api/v1/get_orders [get]
func GetOrders(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, models.Error{Error: "Missing authorization token"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// @Summary Get Order Details
// @Description Get details for a specific order
// @Produce json
// @Param order_id query int true "Order ID"
// @Success 200 {array} models.OrderDetail
// @Failure 404 {object} models.Error
// @Router /api/v1/get_order_details [get]
func GetOrderDetails(c *gin.Context) {
	orderIDStr := c.Query("order_id")
	if orderIDStr == "" {
		c.JSON(http.StatusBadRequest, models.Error{Error: "order_id is required"})
		return
	}

	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: "Invalid order_id"})
		return
	}

	var details []models.OrderDetail
	for _, detail := range orderDetails {
		if detail.OrderID == orderID {
			details = append(details, detail)
		}
	}

	if len(details) == 0 {
		c.JSON(http.StatusNotFound, models.Error{Error: "Order not found"})
		return
	}

	c.JSON(http.StatusOK, details)
}

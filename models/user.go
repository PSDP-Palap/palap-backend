package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type Order struct {
	ID              int    `json:"id"`
	UserID          int    `json:"user_id"`
	ServiceID       int    `json:"service_id"`
	AddressID       int    `json:"address_id"`
	Status          string `json:"status"`
	PaymentMethod   string `json:"payment_method"`
	QRCode          string `json:"qr_code,omitempty"`
	CreatedAt       string `json:"created_at"`
}

type OrderDetail struct {
	ID          int    `json:"id"`
	OrderID     int    `json:"order_id"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}

type PaymentRequest struct {
	ServiceID     int         `json:"service_id" binding:"required"`
	AddressID     int         `json:"address_id" binding:"required"`
	PaymentMethod string      `json:"payment_method" binding:"required"`
	PaymentCard   *PaymentCard `json:"payment_card"`
}

type PaymentCard struct {
	CardNumber string `json:"card_number"`
	CardHolder string `json:"card_holder"`
	ExpiryDate string `json:"expiry_date"`
	CVV        string `json:"cvv"`
}

type PaymentResponse struct {
	Status string `json:"status"`
	QRCode string `json:"qr_code,omitempty"`
}

package models

type Admin struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Status   string `json:"status"`
}

type Freelancer struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Bio         string `json:"bio"`
	Skills      string `json:"skills"`
	Rating      int    `json:"rating"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
}

type FreelanceJob struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Budget      int    `json:"budget"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
}

type ProductService struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
	Status      string `json:"status"`
}

type ShopLocation struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	City     string `json:"city"`
	Country  string `json:"country"`
	Status   string `json:"status"`
}

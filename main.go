// @title PALAP Backend API
// @version 1.0
// @description A comprehensive API for managing services, admins, freelancers, and products
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @basePath /
// @schemes http https
// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
package main

import (
	"log"
	"palap_backend/routers"
)

func main() {
	r := routers.InitializeRouter()

	if err := r.Run(); err != nil {
		log.Printf("failed to run server: %v", err)
	}
}

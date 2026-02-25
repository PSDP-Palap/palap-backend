package v1

import (
"net/http"
"strconv"

"palap_backend/models"

"github.com/gin-gonic/gin"
)

var admins []models.Admin = []models.Admin{
{ID: 1, Username: "admin1", Email: "admin1@example.com", Role: "Super Admin", Status: "active"},
{ID: 2, Username: "admin2", Email: "admin2@example.com", Role: "Admin", Status: "active"},
}

// @Summary Get All Admins
// @Description Get list of all admins
// @Produce json
// @Success 200 {array} models.Admin
// @Router /api/v1/manage_admin [get]
func GetAdmins(c *gin.Context) {
c.JSON(http.StatusOK, admins)
}

// @Summary Get Admin by ID
// @Description Get admin details by ID
// @Produce json
// @Param id path int true "Admin ID"
// @Success 200 {object} models.Admin
// @Failure 404 {object} models.Error
// @Router /api/v1/manage_admin/:id [get]
func GetAdmin(c *gin.Context) {
id, err := strconv.Atoi(c.Param("id"))
if err != nil {
c.JSON(http.StatusBadRequest, models.Error{Error: "Invalid admin ID"})
return
}

for _, admin := range admins {
if admin.ID == id {
c.JSON(http.StatusOK, admin)
return
}
}

c.JSON(http.StatusNotFound, models.Error{Error: "Admin not found"})
}

// @Summary Create Admin
// @Description Create new admin
// @Accept json
// @Produce json
// @Param admin body models.Admin true "Admin details"
// @Success 201 {object} models.Admin
// @Failure 400 {object} models.Error
// @Router /api/v1/manage_admin [post]
func CreateAdmin(c *gin.Context) {
admin := new(models.Admin)
if err := c.BindJSON(admin); err != nil {
c.JSON(http.StatusBadRequest, models.Error{Error: "Invalid request body"})
return
}

admin.ID = len(admins) + 1
admins = append(admins, *admin)

c.JSON(http.StatusCreated, admin)
}

// @Summary Update Admin
// @Description Update existing admin
// @Accept json
// @Produce json
// @Param id path int true "Admin ID"
// @Param admin body models.Admin true "Updated admin details"
// @Success 200 {object} models.Admin
// @Failure 404 {object} models.Error
// @Router /api/v1/manage_admin/:id [put]
func UpdateAdmin(c *gin.Context) {
id, err := strconv.Atoi(c.Param("id"))
if err != nil {
c.JSON(http.StatusBadRequest, models.Error{Error: "Invalid admin ID"})
return
}

updatedAdmin := new(models.Admin)
if err := c.BindJSON(updatedAdmin); err != nil {
c.JSON(http.StatusBadRequest, models.Error{Error: "Invalid request body"})
return
}

for i, admin := range admins {
if admin.ID == id {
updatedAdmin.ID = id
admins[i] = *updatedAdmin
c.JSON(http.StatusOK, updatedAdmin)
return
}
}

c.JSON(http.StatusNotFound, models.Error{Error: "Admin not found"})
}

// @Summary Delete Admin
// @Description Delete admin by ID
// @Param id path int true "Admin ID"
// @Success 204
// @Failure 404 {object} models.Error
// @Router /api/v1/manage_admin/:id [delete]
func DeleteAdmin(c *gin.Context) {
id, err := strconv.Atoi(c.Param("id"))
if err != nil {
c.JSON(http.StatusBadRequest, models.Error{Error: "Invalid admin ID"})
return
}

for i, admin := range admins {
if admin.ID == id {
admins = append(admins[:i], admins[i+1:]...)
c.JSON(http.StatusNoContent, nil)
return
}
}

c.JSON(http.StatusNotFound, models.Error{Error: "Admin not found"})
}

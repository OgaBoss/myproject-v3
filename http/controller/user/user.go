package user

import (
	"github.com/OgaBoss/myproject-v3/database/model"
	"github.com/System-Glitch/goyave/v3"
	"github.com/System-Glitch/goyave/v3/database"
	"net/http"
)

func Index(response *goyave.Response, request *goyave.Request) {
	users := []model.User{}
	result := database.Conn().Find(&users)
	if response.HandleDatabaseError(result) {
		response.JSON(http.StatusOK, users)
	}
}
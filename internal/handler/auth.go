package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"virtual-strike-backend-go/pkg/models"
	"virtual-strike-backend-go/pkg/modules"
)

type LoginInput struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (h *Handler) Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBind(&input); err != nil {
		var response modules.Response
		response.Status = false
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	token, err := models.LoginCheck(u.Username, u.Password)

	if err != nil {
		var response modules.Response
		response.Status = false
		response.Message = "username or password is incorrect"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var response modules.Response
	response.Status = true
	response.Message = token
	c.JSON(http.StatusOK, gin.H{"status": true, "message": token})
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		var response modules.Response
		response.Status = false
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	_, err := u.SaveUser()

	if err != nil {
		var response modules.Response
		response.Status = false
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var response modules.Response
	response.Status = true
	response.Message = "registration success"
	c.JSON(http.StatusOK, response)
}

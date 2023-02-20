package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"net/http"
	"virtual-strike-backend-go/pkg/modules"
	monitoring "virtual-strike-backend-go/pkg/moniroting"
)

func (h *Handler) Customer(c *gin.Context) {
	var jsonInput modules.CustomerRequest
	if err := c.ShouldBind(&jsonInput); err != nil {
		var response modules.Response
		response.Status = false
		response.Message = err.Error()
		logrus.Error(err)
		monitoring.ErrorHandler.With(prometheus.Labels{"error_message": err.Error()}).Inc()
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	code, any := h.services.CustomerLogic(jsonInput)

	if code != 200 {
		logrus.Error()
		monitoring.ErrorHandler.With(prometheus.Labels{"error_message": any.Message.ErrorMessage}).Inc()
	}

	c.JSON(code, any)
}

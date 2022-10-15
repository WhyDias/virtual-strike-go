package handler

import "github.com/gin-gonic/gin"

// @Summary      Usage
// @Description  Usage func
// @Tags         Post
// @Accept       json
// @Produce      json
// @Param        Token   body   modules.DataUsage  true  "Usage"
// @Success      200  {object}	modules.ResponseUsage
// @Failure      500  {object}  modules.ResponseUsage
// @Router       /api/v1/colvir/usage [post]
func (h *Handler) search(c *gin.Context) {
	//var jsonInput modules.DataUsage
	//if err := c.ShouldBindJSON(&jsonInput); err != nil {
	//	var response modules.ResponseUsage
	//	response.Code = 500
	//	response.Description = err.Error()
	//	logrus.Error(err)
	//	monitoring.ErrorHandler.With(prometheus.Labels{"error_message": err.Error()}).Inc()
	//	c.JSON(http.StatusInternalServerError, response)
	//	return
	//}
	//
	//code, any := h.services.CheckUsage(jsonInput)
	//
	//if code != 200 {
	//	logrus.Error()
	//	monitoring.ErrorHandler.With(prometheus.Labels{"error_message": any.Description}).Inc()
	//}
	//
	//c.JSON(code, any)

	c.JSON(200, "HELLO")
}

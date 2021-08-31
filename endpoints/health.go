package endpoints

import "github.com/gin-gonic/gin"

// @Summary Health check endpoint
// @produce text/plain
// @Success 200 {object} string "OK"
// @Router /health [get]
func Health(context *gin.Context) {
	context.String(200, "OK")
}

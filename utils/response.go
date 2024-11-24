package utils

import "github.com/gin-gonic/gin"

// RespondJSON sends a JSON response
func RespondJSON(c *gin.Context, status int, payload interface{}) {
	c.JSON(status, payload)
}

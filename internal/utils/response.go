package utils

import "github.com/gin-gonic/gin"

// RespondWithError sends a JSON response with an error message and HTTP status code.
func RespondWithError(ctx *gin.Context, code int, message string) {
	RespondWithJSON(ctx, code, gin.H{
		"error": message,
	})
}

// RespondWithJSON sends a JSON response with a given payload and HTTP status code.
func RespondWithJSON(ctx *gin.Context, code int, payload interface{}) {
	ctx.JSON(code, payload)
}

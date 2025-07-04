package middlewares

import (
	"fmt"
	"gin-app/errs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
				switch err := r.(type) {
				case *errs.InvalidInput:
					c.JSON(err.StatusCode, gin.H{
						"code":    err.Code,
						"error":   err.Error,
						"message": err.Message,
					})
				case *errs.NotFoundInput:
					c.JSON(err.StatusCode, gin.H{
						"statusCode": err.StatusCode,
						"code":       err.Code,
						"error":      err.Error,
						"message":    err.Message,
					})
				case *errs.BadRequest:
					c.JSON(err.StatusCode, gin.H{
						"statusCode": err.StatusCode,
						"code":       err.Code,
						"error":      err.Error,
						"message":    err.Message,
					})
				default:
					c.JSON(http.StatusInternalServerError, gin.H{
						"error":   "INTERNAL_ERROR",
						"message": "Something went wrong",
					})

				}
				c.Abort()
			}
		}()
		c.Next()
	}
}

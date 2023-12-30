package http

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Code  int    `json:"code"`
	Cause string `json:"cause"`
}

func RespondWithError(c *gin.Context, code int, cause string) {
	c.JSON(code, ErrorResponse{
		Code:  code,
		Cause: cause,
	})
}

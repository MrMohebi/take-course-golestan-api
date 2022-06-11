package common

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// ValidBindJson bind request body(json) and validate fields with Validator-v10. If err returns json response to gin
func ValidBindJson(c *gin.Context, reqBody interface{}) bool {
	var validate = validator.New()
	if err := c.BindJSON(reqBody); err != nil {
		c.JSON(http.StatusBadRequest, http.StatusBadRequest)
		return false
	}
	if validationErr := validate.Struct(reqBody); validationErr != nil {
		c.JSON(http.StatusBadRequest, validationErr.Error())
		return false
	}
	return true
}

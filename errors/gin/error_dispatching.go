package gin

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler func(c *gin.Context, err error)
type Selector func(c *gin.Context, err error) bool

type Dispatchers struct {
	handlers []Selector
}

func Handlers() Dispatchers {

	return Dispatchers{handlers: make([]Selector, 0)}
}

type MH map[string]Handler

func (dispatch *Dispatchers) AWS(awsHandlers MH, DefaultHandler Handler) {

	dispatch.handlers = append(dispatch.handlers, func(c *gin.Context, err error) bool {

		awsError, ok := err.(awserr.Error)
		if ok {
			for key, value := range awsHandlers {
				if key == awsError.Code() {
					value(c, err)
					return true
				}
			}
			DefaultHandler(c, err)
			return true
		} else {
			return false
		}
	})
}

func (dispatch *Dispatchers) Always(AlwaysRunHandler Handler) {

	dispatch.handlers = append(dispatch.handlers, func(c *gin.Context, err error) bool {

		AlwaysRunHandler(c, err)
		return true
	})
}

func (dispatch *Dispatchers) Validation(httpCode int) {

	dispatch.handlers = append(dispatch.handlers, func(c *gin.Context, err error) bool {

		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {

			return false
		}

		fieldErrors := make([]interface{}, 0)

		for _, element := range validationErrors {

			one := gin.H{
				"field": gin.H{
					"name":      element.Field(),
					"errorCode": element.Tag()},
			}
			fieldErrors = append(fieldErrors, one)
		}

		errorList := gin.H{"error": fieldErrors}
		c.JSON(httpCode, errorList)
		return true
	})
}

func (dispatch Dispatchers) HandleError(c *gin.Context, err error) bool {

	for _, handler := range dispatch.handlers {

		if handler(c, err) == true {
			return true
		}
	}

	return false
}

func HttpError(errorCode int) Handler {
	return func(c *gin.Context, err error) {

		c.Writer.WriteHeader(errorCode)
		return
	}
}

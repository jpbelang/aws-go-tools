package gin

import (
	"github.com/gin-gonic/gin"
	assertions "github.com/stretchr/testify/assert"
	"testing"
)

func TestDispatchers_AWS(t *testing.T) {

	assert := assertions.New(t)

	mockError := MockError{}
	mockError.On("Code").Return("someError")

	mockWriter := NewMockWriter()
	mockWriter.On("WriteHeader", 500).Once().Return()

	ginContext := gin.Context{Writer: mockWriter}
	handler := Handlers()
	handler.AWS(MH{"someError": func(c *gin.Context, err error) {

		c.Writer.WriteHeader(500)
	}}, func(c *gin.Context, err error) {
		assert.Fail("Shouldn't get to the default!")
	})

	handler.HandleError(&ginContext, mockError)
	mockWriter.AssertCalled(t, "WriteHeader", 500)
}

func TestDispatchers_FallsThrough_AWS(t *testing.T) {

	assert := assertions.New(t)

	mockError := new(MockError)
	mockError.On("Code").Return("notTheSameError")

	mockWriter := NewMockWriter()
	mockWriter.On("WriteHeader", 500).Once().Return()

	ginContext := gin.Context{Writer: mockWriter}
	handler := Handlers()
	handler.AWS(MH{"someError": func(c *gin.Context, err error) {

		assert.Fail("Shouldn't get to the a specific!")
	}}, func(c *gin.Context, err error) {

		c.Writer.WriteHeader(500)
	})

	handler.HandleError(&ginContext, mockError)
	mockWriter.AssertCalled(t, "WriteHeader", 500)
}

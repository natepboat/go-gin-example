package service

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {

	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	ListAlbum(c)

	assert.Equal(t, 200, response.Code)
	assert.NotNil(t, response.Body)
}

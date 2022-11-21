package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	exptedResp := Response{
		Message: "UP",
	}

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	request := httptest.NewRequest("GET", "/api/v1/health", nil)
	ctx.Request = request
	HealthCheck(ctx)
	var actualResponse Response

	json.NewDecoder(recorder.Body).Decode(&actualResponse)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, exptedResp, actualResponse)
}

package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Kubeitron/soda-api/cmd/api/handlers"
	"github.com/stretchr/testify/assert"
)

var healthJSON = `{"status":"Healthy!"}` + "\n"

func TestHealthCheck(t *testing.T) {
	h := MockAPI()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := h.API.NewContext(req, rec)
	c.SetPath("/health")

	hch := handlers.NewHealthcheckHandler()
	if assert.NoError(t, hch.Healthcheck(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, healthJSON, rec.Body.String())
	}

}

package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ryne777/Golang_first_project/internal/routers"
	"github.com/stretchr/testify/assert"
)

func TestMainRoute(t *testing.T) {
	router := routers.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

}

package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ryne777/Golang_first_project/internal/config"
	"github.com/Ryne777/Golang_first_project/internal/routers"
	"github.com/stretchr/testify/assert"
)

var cfg = config.GetConfig()
var router = routers.SetupRouter(cfg)

func TestMainRoute(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

}

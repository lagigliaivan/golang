package categoryht

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_that_service_exposes_a_ping_resource(t *testing.T) {

	router := SetupRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/ping", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())

}

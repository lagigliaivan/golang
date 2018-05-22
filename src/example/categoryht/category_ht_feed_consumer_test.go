package categoryht

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Ping_resource_returns_200Ok_when_receiving_a_GET(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func Test_handlingtime_resource_returns_400_when_receiving_a_POST_without_orderid(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/category/handlingtime", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func Test_handlingtime_resource_accepts_an_order_id(t *testing.T) {
	router := SetupRouter()
	var body = []byte(`{"order_id":123}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/category/handlingtime", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

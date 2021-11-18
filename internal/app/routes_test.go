package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealtCheckHandler(t *testing.T) {
	router := Start()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	var pingResponse struct {
		Message string `json:"message"`
	}

	jsonResponse, _ := ioutil.ReadAll(w.Body)

	if err := json.Unmarshal(jsonResponse, &pingResponse); err != nil {
		t.Fail()
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", pingResponse.Message)
}

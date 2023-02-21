package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	swagger "github.com/Nikolay200669/swag1/apiclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// run only after running application
func TestPingRouteSwagg(t *testing.T) {
	cnf := swagger.NewConfiguration()
	cl := swagger.NewAPIClient(cnf)
	cl.ChangeBasePath("http://:8080") // or "http://0.0.0.0:8080"
	resp, httpResponse, err := cl.ExampleApi.PingGet(context.Background())

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, httpResponse.StatusCode)
	assert.Equal(t, "pong", resp.Name)
}

// run any time
func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp Resp
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, "pong", resp.Name)
}

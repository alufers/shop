package shop

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
)

// MakeTestingCtx creates a simple context usable while testing
func MakeTestingCtx() *AppCtx {
	return InitAppCtx(&AppConfig{
		DBDialect: "sqlite3",
		DBArgs:    []interface{}{path.Join(os.TempDir(), "shop-test.db")},
	})
}

func PerformTestRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func PerformTestJSONRequest(r http.Handler, method, path string, body io.Reader, v interface{}) *httptest.ResponseRecorder {
	resp := PerformTestRequest(r, method, path, body)
	decoder := json.NewDecoder(resp.Body)
	err := decoder.Decode(v)
	if err != nil {
		panic(err)
	}
	return resp
}

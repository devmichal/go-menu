package cmd

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"testing"
)

var (
	foodJSON = `{"pricing":5,"menu":"Pizza"}`
)

func TestGetFood(t *testing.T) {
	app := Route()

	req := httptest.NewRequest("GET", "/menu/3", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
	defer resp.Body.Close()

	b, _ := io.ReadAll(resp.Body)

	assert.Equal(t, foodJSON, string(b))
}

func TestCreateFood(t *testing.T) {
	app := Route()

	req := httptest.NewRequest("POST", "/menu", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
	defer resp.Body.Close()

	b, _ := io.ReadAll(resp.Body)

	assert.Equal(t, foodJSON, string(b))
}

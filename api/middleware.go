package api

import (
	"os"

	"github.com/labstack/echo/v4"
)

func CheckApiKey(key string, c echo.Context) (bool, error) {
	API_KEY := os.Getenv("API_KEY")
	return key == API_KEY, nil
}

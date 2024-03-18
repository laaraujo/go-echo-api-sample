package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/labstack/echo/v4/middleware"
)

const defaultPort = 1323

var CustomLoggerConfig = middleware.LoggerConfig{
	Skipper:          middleware.DefaultSkipper,
	Format:           "${method} ${status} ${uri} ${remote_ip} ${user_agent}\n",
	CustomTimeFormat: "2006-01-02 15:04:05.00000",
}

func GetServerString() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = strconv.Itoa(defaultPort)
	}
	return fmt.Sprintf(":%v", port)
}

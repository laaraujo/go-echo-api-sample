package api

import (
	"golang-echo-api/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ExerciseRouter(g *echo.Group, app *handler.App) {
	g.GET("", app.GetExerciseList)
	g.POST("", app.CreateExercise)
	g.GET("/:id", app.GetExercise)
	g.PUT("/:id", app.UpdateExercise)
	g.DELETE("/:id", app.DeleteExercise)
}

func RoutineRouter(g *echo.Group, app *handler.App) {
	g.GET("", app.GetRoutineList)
	g.POST("", app.CreateRoutine)
	g.GET("/:id", app.GetRoutine)
	g.PUT("/:id", app.UpdateRoutine)
	g.DELETE("/:id", app.DeleteRoutine)
}

func Router(e *echo.Echo, app *handler.App) {
	e.GET("/ping", handler.PingHandler)
	RoutineRouter(e.Group("/routines", middleware.KeyAuth(CheckApiKey)), app)
	ExerciseRouter(e.Group("/exercises", middleware.KeyAuth(CheckApiKey)), app)
}

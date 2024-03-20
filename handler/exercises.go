package handler

import (
	db "golang-echo-api/db/sqlc"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (app *App) GetExerciseList(c echo.Context) error {
	cc := c.Request().Context()
	exercises, err := app.Queries.ListExercises(cc)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, exercises)
}

func (app *App) CreateExercise(c echo.Context) error {
	cc := c.Request().Context()
	e := new(db.CreateExerciseParams)
	err := c.Bind(e)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	exercise, err := app.Queries.CreateExercise(cc, *e)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, exercise)
}

func (app *App) GetExercise(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Exercise ID")
	}
	cc := c.Request().Context()
	e, err := app.Queries.GetExercise(cc, int32(id))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, "Exercise not found")
	}
	return c.JSON(http.StatusOK, e)
}

func (app *App) UpdateExercise(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Exercise ID")
	}

	cc := c.Request().Context()

	if _, err := app.Queries.GetExercise(cc, int32(id)); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, "Exercise not found")
	}

	exercise := new(db.UpdateExerciseParams)
	exercise.ID = int32(id)
	if err := c.Bind(exercise); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if err := app.Queries.UpdateExercise(cc, *exercise); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, exercise)
}

func (app *App) DeleteExercise(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Exercise ID")
	}

	cc := c.Request().Context()

	if err := app.Queries.DeleteExercise(cc, int32(id)); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, "")
}

package handler

import (
	db "golang-echo-api/db/sqlc"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (app *App) GetRoutineList(c echo.Context) error {
	cc := c.Request().Context()
	routines, err := app.Queries.ListRoutines(cc)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, routines)
}

func (app *App) CreateRoutine(c echo.Context) error {
	cc := c.Request().Context()
	r := new(struct {
		Name string `json:"name"`
	})
	err := c.Bind(r)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	routine, err := app.Queries.CreateRoutine(cc, r.Name)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, routine)
}

func (app *App) GetRoutine(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Routine ID")
	}
	cc := c.Request().Context()
	r, err := app.Queries.GetRoutine(cc, int32(id))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, "Routine not found")
	}
	return c.JSON(http.StatusOK, r)
}

func (app *App) UpdateRoutine(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Routine ID")
	}

	cc := c.Request().Context()

	if _, err := app.Queries.GetRoutine(cc, int32(id)); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, "Routine not found")
	}

	routine := new(db.UpdateRoutineParams)
	routine.ID = int32(id)
	if err := c.Bind(routine); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if err := app.Queries.UpdateRoutine(cc, *routine); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, routine)
}

func (app *App) DeleteRoutine(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Routine ID")
	}

	cc := c.Request().Context()

	if err := app.Queries.DeleteRoutine(cc, int32(id)); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, "")
}

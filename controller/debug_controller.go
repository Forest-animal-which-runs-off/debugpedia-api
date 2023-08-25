package controller

import (
	"debugpedia-api/model"
	"debugpedia-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IDebugController interface {
	GetAllDebugs(c echo.Context) error
	GetDebugById(c echo.Context) error
	CreateDebug(c echo.Context) error
	UpdateDebug(c echo.Context) error
	DeleteDebug(c echo.Context) error
}

type debugController struct {
	du usecase.IDebugUsecase
}

func NewDebugController(du usecase.IDebugUsecase) IDebugController {
	return &debugController{du}
}

func (dc *debugController) GetAllDebugs(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	debugRes, err := dc.du.GetAllDebugs(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, debugRes)
}

func (dc *debugController) GetDebugById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("taskId")
	debugId, _ := strconv.Atoi(id)
	debugRes, err := dc.du.GetDebugById(uint(userId.(float64)), uint(debugId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, debugRes)
}

func (dc *debugController) CreateDebug(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	debug := model.Debug{}
	if err := c.Bind(&debug); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	debug.UserId = uint(userId.(float64))
	debugRes, err := dc.du.CreateDebug(debug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, debugRes)
}

func (dc *debugController) UpdateDebug(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("taskId")
	debugId, _ := strconv.Atoi(id)

	debug := model.Debug{}
	if err := c.Bind(&debug); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	debugRes, err := dc.du.UpdateDebug(debug, uint(userId.(float64)), uint(debugId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, debugRes)
}

func (dc *debugController) DeleteDebug(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("taskId")
	debugId, _ := strconv.Atoi(id)

	err := dc.du.DeleteDebug(uint(userId.(float64)), uint(debugId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

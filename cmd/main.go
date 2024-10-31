package main

import (
	"devices"
	"devices/database"
	handler2 "devices/internal/handler"
	"devices/internal/model"
	repository2 "devices/internal/repository"
	"devices/internal/service"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)

func main() {
	cfg := database.NewMysqlConfig()
	db, err := database.NewDbConnection(cfg)
	if err != nil {
		panic(err)
	}

	err = database.ApplyMigrations(&devices.FS, "database/migration", cfg)
	if err != nil {
		panic(err)
	}

	repository := repository2.NewDeviceRepository(db)
	deviceService := service.NewDeviceService(repository)

	handler := handler2.NewApiHandler(deviceService)

	e := echo.New()

	// Root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Devices Management Service")
	})

	// list all devices
	e.GET("/v1/devices", func(c echo.Context) error {
		response := handler.HandleDeviceGetAll()

		return c.JSON(response.HttpCode, response.Data)
	})

	// get device by id
	e.GET("/v1/devices/:id", func(c echo.Context) error {
		deviceId, err := strconv.Atoi(c.Param("id"))
		if err != nil { // incorrect input parameter
			c.JSON(http.StatusBadRequest, handler2.MakeErrorObject(err))
		}

		response := handler.HandleDeviceGet(int64(deviceId))

		return c.JSON(response.HttpCode, response.Data)
	})

	// add device
	e.POST("/v1/devices", func(c echo.Context) error {
		device := &model.Device{}
		err := json.NewDecoder(c.Request().Body).Decode(&device)
		if err != nil {
			return c.JSON(http.StatusBadRequest, handler2.MakeErrorObject(err))
		}

		response := handler.HandleDeviceAdd(device)

		return c.JSON(response.HttpCode, response.Data)
	})

	// fully update device info
	e.PUT("/v1/devices/:id", func(c echo.Context) error {
		deviceId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, handler2.MakeErrorObject(err))
		}
		device := &model.Device{}
		err = json.NewDecoder(c.Request().Body).Decode(&device)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, handler2.MakeErrorObject(err))
		}
		device.Id = int64(deviceId)

		response := handler.HandleDeviceUpdate(device)

		return c.JSON(response.HttpCode, response.Data)
	})

	// partial update device info
	e.PATCH("/v1/devices/:id", func(c echo.Context) error {
		deviceId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, handler2.MakeErrorObject(err))
		}
		device := &model.Device{}
		err = json.NewDecoder(c.Request().Body).Decode(&device)
		if err != nil {
			return c.JSON(http.StatusBadRequest, handler2.MakeErrorObject(err))
		}
		device.Id = int64(deviceId)

		response := handler.HandleDeviceUpdatePartial(device)

		return c.JSON(response.HttpCode, response.Data)
	})

	e.DELETE("/v1/devices/:id", func(c echo.Context) error {
		deviceId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, handler2.MakeErrorObject(err))
		}

		response := handler.HandleDeviceDelete(int64(deviceId))

		return c.JSON(response.HttpCode, response.Data)
	})

	e.GET("/v1/devices/search", func(c echo.Context) error {
		brand := c.QueryParam("brand")
		if brand == "" {
			c.JSON(http.StatusBadRequest, handler2.MakeErrorObject(errors.New("brand parameter is required")))
		}
		response := handler.HandleDeviceSearchByBrand(brand)

		return c.JSON(response.HttpCode, response.Data)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

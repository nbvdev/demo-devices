package handler

import (
	"devices/interfaces"
	"devices/internal/model"
	"errors"
	"net/http"
)

type ApiHandler struct {
	service interfaces.DeviceService
}

type Response struct {
	HttpCode int
	Data     interface{}
}

func CreateResponse(httpCode int, data interface{}) Response {
	return Response{HttpCode: httpCode, Data: data}
}
func CreateErrorResponse(httpCode int, err error) Response {
	return Response{HttpCode: httpCode, Data: MakeErrorObject(err)}
}

func NewApiHandler(deviceService interfaces.DeviceService) ApiHandler {
	return ApiHandler{
		service: deviceService,
	}
}

func MakeErrorObject(err error) interface{} {
	data := make(map[string]string)
	data["error"] = err.Error()

	return data
}

func (h ApiHandler) HandleDeviceGetAll() Response {
	devices, err := h.service.GetAll()
	if err != nil {
		return CreateErrorResponse(http.StatusInternalServerError, err)
	}

	return CreateResponse(http.StatusOK, devices)
}

func (h ApiHandler) HandleDeviceGet(deviceId int64) Response {
	device, err := h.service.Get(deviceId)
	if err != nil {
		return CreateErrorResponse(http.StatusInternalServerError, err)
	}

	return CreateResponse(http.StatusOK, device)
}

func (h ApiHandler) HandleDeviceAdd(device *model.Device) Response {
	if device == nil {
		return CreateErrorResponse(http.StatusInternalServerError, errors.New("device required"))
	}
	device, err := h.service.AddNew(device)
	if err != nil {
		return CreateErrorResponse(http.StatusInternalServerError, err)
	}

	return CreateResponse(http.StatusOK, device)
}

func (h ApiHandler) HandleDeviceUpdate(device *model.Device) Response {
	device, err := h.service.Update(device)
	if err != nil {
		return CreateErrorResponse(http.StatusInternalServerError, err)
	}
	if device == nil {
		return CreateErrorResponse(http.StatusNotFound, errors.New("device not found"))
	}

	return CreateResponse(http.StatusOK, device)
}

func (h ApiHandler) HandleDeviceUpdatePartial(device *model.Device) Response {
	device, err := h.service.UpdatePartial(device)
	if err != nil {
		return CreateErrorResponse(http.StatusInternalServerError, err)
	}
	if device == nil {
		return CreateErrorResponse(http.StatusNotFound, errors.New("device not found"))
	}

	return CreateResponse(http.StatusOK, device)
}

func (h ApiHandler) HandleDeviceDelete(deviceId int64) Response {
	err := h.service.Delete(deviceId)
	if err != nil {
		return CreateErrorResponse(http.StatusInternalServerError, err)
	}

	return CreateResponse(http.StatusOK, nil)
}

func (h ApiHandler) HandleDeviceSearchByBrand(brand string) Response {
	devices, err := h.service.SearchByBrand(brand)
	if err != nil {
		return CreateErrorResponse(http.StatusInternalServerError, err)
	}

	return CreateResponse(http.StatusOK, devices)
}

package handler

import (
	"errors"
	"reflect"
	"testing"
)

func TestCreateResponse(t *testing.T) {
	testHttpCode := 123
	testMessage := "test message"
	resp := CreateResponse(testHttpCode, testMessage)
	if resp.HttpCode != testHttpCode {
		t.Error("expected ", testHttpCode, "got ", resp.HttpCode)
	}
	if resp.Data != testMessage {
		t.Error("expected ", testMessage, "got ", resp.Data)
	}
}

func TestCreateErrorResponse(t *testing.T) {
	testError := errors.New("test error")
	testHttpCode := 500
	resp := CreateErrorResponse(testHttpCode, testError)
	if resp.HttpCode != testHttpCode {
		t.Error("expected ", testHttpCode, "got ", resp.HttpCode)
	}
	expectedData := map[string]string{"error": testError.Error()}
	if reflect.DeepEqual(expectedData, resp.Data) == false {
		t.Error("expected ", expectedData, "got ", resp.Data)
	}
}

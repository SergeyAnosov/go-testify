package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", `/cafe?city=moscow&count=2`, nil)

	responseRecorder := httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(mainHandle)
	handlerFunc.ServeHTTP(responseRecorder, req)

	assert.Equal(t, responseRecorder.Code, http.StatusOK)

	body := responseRecorder.Body.String()
	require.NotEmpty(t, body)
}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=london&count=1", nil)

	responseRecorder := httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(mainHandle)
	handlerFunc.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusBadRequest)

	answer := "wrong city value"
	body := responseRecorder.Body.String()
	require.Equal(t, body, answer)

}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", `/cafe?city=moscow&count=200`, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	body := responseRecorder.Body.String()
	slice := strings.Split(body, ",")

	assert.Equal(t, responseRecorder.Code, http.StatusOK)
	require.Len(t, slice, totalCount)
}

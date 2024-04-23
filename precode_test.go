package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest(http.MethodGet, "/cafe?count="+strconv.Itoa(totalCount+1)+"&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	cafeTestList := strings.Split(responseRecorder.Body.String(), ",")

	// код 200
	require.Equal(t, responseRecorder.Code, http.StatusOK)
	// Возвращает все доступные кафе
	assert.Len(t, cafeTestList, totalCount)
}

func TestMainHandlerWhenCityNotSupport(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=4&city=qwerty", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки

	// код 400
	require.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	// тело содердит "wrong city value"
	assert.Equal(t, responseRecorder.Body.String(), "wrong city value")
}

func TestMainHandlerWhenCorrect(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки

	// код 200
	require.Equal(t, responseRecorder.Code, http.StatusOK)
	// тело не пустое
	assert.NotEqual(t, responseRecorder.Body.String(), "")
}

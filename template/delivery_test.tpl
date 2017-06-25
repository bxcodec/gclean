package {{ .ModelName }}_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/stretchr/testify/assert"

	"github.com/bxcodec/faker"
  {{ if   (gt (len .Imports) 0) }}

  {{- range $key, $val := .Imports}}
  		{{- if not (eq ($val.Alias) ($val.Path) ) }}
  	{{$val.Alias}}  "{{$val.Path}}"
  		{{- else }}
    "{{$val.Path}}"
  		{{- end }}
  {{- end}}
  {{ end }}
)


{{ $camelModel :=  .ModelName | camelcase }}
{{ $httpHandler:= index .Imports "handler" }}


func TestFetch(t *testing.T) {
	var mock{{ $camelModel }} models.{{ $camelModel }}
	err := faker.FakeData(&mock{{ $camelModel }})
	assert.NoError(t, err)
	mockUCase := new(mocks.{{ $camelModel }}Usecase)
	mockList{{ $camelModel }} := make([]*models.{{ $camelModel }}, 0)
	mockList{{ $camelModel }} = append(mockList{{ $camelModel }}, &mock{{ $camelModel }} )
	num := 1
	cursor := "2"
	mockUCase.On("Fetch", cursor, int64(num)).Return(mockList{{ $camelModel }}, "10", nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/article?num=1&cursor="+cursor, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := {{ $httpHandler.Alias }}.{{ $camelModel }}Handler{mockUCase}
	handler.Fetch{{ $camelModel }}(c)

	responseCursor := rec.Header().Get("X-Cursor")
	assert.Equal(t, "10", responseCursor)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertCalled(t, "Fetch", cursor, int64(num))
}

func TestFetchError(t *testing.T) {
	mockUCase := new(mocks.{{ $camelModel }}Usecase)
	num := 1
	cursor := "2"
	mockUCase.On("Fetch", cursor, int64(num)).Return(nil, "", errors.New("Internal Server Error "))

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/article?num=1&cursor="+cursor, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := {{ $httpHandler.Alias }}.{{ $camelModel }}Handler{mockUCase}
	handler.Fetch{{ $camelModel }}(c)

	responseCursor := rec.Header().Get("X-Cursor")
	assert.Equal(t, "", responseCursor)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUCase.AssertCalled(t, "Fetch", cursor, int64(num))
}

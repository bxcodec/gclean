package {{.ModelName | lower}}_test

import (
	"errors" 
	"testing"
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

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
{{ $ucase := index .Imports   "usecase" }}

func TestFetch(t *testing.T) {
	mock{{ $camelModel }}Repo := new(mocks.{{ $camelModel }}Repository)
	var mock{{ $camelModel }} models.{{ $camelModel }}
	err := faker.FakeData(&mock{{ $camelModel }})
	assert.NoError(t, err)

	mockList{{ $camelModel }} := make([]*models.{{ $camelModel }}, 0)
	mockList{{ $camelModel }} = append(mockList{{ $camelModel }}, &mock{{ $camelModel }})
	mock{{ $camelModel }}Repo.On("Fetch", mock.AnythingOfType("string"), mock.AnythingOfType("int64")).Return(mockList{{ $camelModel }}, nil)
	u := {{ $ucase.Alias }}.New{{ $camelModel }}Usecase(mock{{ $camelModel }}Repo)
	num := int64(1)
	cursor := "12"
	list, nextCursor, err := u.Fetch(cursor, num)
	cursorExpected :=""

	assert.Equal(t, cursorExpected, nextCursor)
	assert.NoError(t, err)
	assert.Len(t, list, len(mockList{{ $camelModel }}))

	mock{{ $camelModel }}Repo.AssertCalled(t, "Fetch", mock.AnythingOfType("string"), mock.AnythingOfType("int64"))

}

func TestFetchError(t *testing.T) {
	mock{{ $camelModel }}Repo := new(mocks.{{ $camelModel }}Repository)

	mock{{ $camelModel }}Repo.On("Fetch", mock.AnythingOfType("string"), mock.AnythingOfType("int64")).Return(nil, errors.New("Unexpexted Error"))
	u := {{ $ucase.Alias }}.New{{ $camelModel }}Usecase(mock{{ $camelModel }}Repo)
	num := int64(1)
	cursor := "12"
	list, nextCursor, err := u.Fetch(cursor, num)

	assert.Empty(t, nextCursor)
	assert.Error(t, err)
	assert.Len(t, list, 0)
	mock{{ $camelModel }}Repo.AssertCalled(t, "Fetch", mock.AnythingOfType("string"), mock.AnythingOfType("int64"))

}

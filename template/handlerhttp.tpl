package {{.ModelName | lower }}

 import (
  "net/http"
  "strconv"

{{- range $key, $val := .Imports}}
		{{- if not (eq ($val.Alias) ($val.Path) ) }}
	{{$val.Alias}}  "{{$val.Path}}"
		{{- else }}
  "{{$val.Path}}"
		{{- end }}
{{- end}}
)


{{ $modelCamelCase := .ModelName | camelcase -}}

{{- $framework := index .Imports "framework" -}}

{{- $modelLower := .ModelName | lower -}}
{{- $pkgModel := index .Imports  (lower ( nospace (cat $modelLower "usecase"))) -}}

type {{ $modelCamelCase }}Handler struct {
	{{ initials  $modelCamelCase  }}Usecase {{$pkgModel.Alias}}.{{$modelCamelCase}}Usecase
}

func (a *{{ $modelCamelCase }}Handler) Fetch{{ $modelCamelCase }}(c {{$framework.Alias}}.Context) error {

	numS := c.QueryParam("num")
	num, _ := strconv.Atoi(numS)

	cursor := c.QueryParam("cursor")

	listAr, nextCursor, err := a.{{ initials  $modelCamelCase  }}Usecase.Fetch(cursor, int64(num))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something Error On Our Services")
	}
	c.Response().Header().Set(`X-Cursor`, nextCursor)
	return c.JSON(http.StatusOK, listAr)
}

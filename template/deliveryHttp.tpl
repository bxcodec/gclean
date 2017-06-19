package http

{{ if   (gt (len .Imports) 0) }}
import (
{{- range $key, $val := .Imports}}
		{{- if not (eq ($val.Alias) ($val.Path) ) }}
	{{$val.Alias}}  "{{$val.Path}}"
		{{- else }}
  "{{$val.Path}}"
		{{- end }}
{{- end}}
)
{{ end }}


type HttpHandler struct {
	E *echo.Echo
}

{{$modelNameLower := .ModelName | lower }}

{{$pkgModel := index .Imports $modelNameLower }}

func (h *HttpHandler) Init{{.ModelName}}Handler(u {{$pkgModel.Alias}}.{{.ModelName}}Usecase) *HttpHandler {
	{{.ModelName | lower }}Handler := &artHandler.{{.ModelName}}Handler{ {{substr 0 1 .ModelName}}Usecase: u}
	h.E.GET(`/{{.ModelName | lower}}`, {{.ModelName | lower }}Handler.Fetch)
	return h
}

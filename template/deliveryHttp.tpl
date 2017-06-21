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

{{- $framework := index .Imports "framework" -}}

type HttpHandler struct {
	{{- if eq $framework.Alias "echo" }}
	 E *echo.Echo
	{{- end }}
}



{{- range $model := .Data }}

{{- $modelLower := $model.ModelName | lower -}}

{{- $pkgModel := index $model.Imports  "usecase" -}}
{{- $modelCamelCase := $model.ModelName | camelcase -}}



{{- $handlerPkg := index $model.Imports (lower (nospace ( cat $model.ModelName "handler"))) }}
// Init{{ $modelCamelCase }}HttpHandler Used for initializing the Handler for {{ $modelCamelCase }}
func (h *HttpHandler) Init{{ $modelCamelCase }}Handler(u {{$pkgModel.Alias}}.{{$modelCamelCase}}Usecase) *HttpHandler {
	{{$model.ModelName | lower }}H := &{{$handlerPkg.Alias}}.{{$modelCamelCase}}Handler{  {{ (initials $modelCamelCase) }}Usecase: u}
	h.E.GET(`/{{$model.ModelName | lower}}`, {{$model.ModelName | lower }}H.Fetch{{ $modelCamelCase }})
	return h
}

{{- end -}}

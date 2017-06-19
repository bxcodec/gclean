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

{{- $pkgModel := index $model.Imports $modelLower -}}
{{- $modelCamelCase := $model.ModelName | camelcase -}}



{{- $handlerPkg := index $model.Imports (lower (nospace ( cat $model.ModelName "handler"))) }}

func (h *HttpHandler) Init{{ $modelCamelCase }}Handler(u {{$pkgModel.Alias}}.{{$model.ModelName}}Usecase) *HttpHandler {
	{{$model.ModelName | lower }}_H := &{{$handlerPkg.Alias}}.{{$model.ModelName}}Handler{  {{ (initials $modelCamelCase) }}Usecase: u}
	h.E.GET(`/{{$model.ModelName | lower}}`, {{$model.ModelName | lower }}_H.Fetch)
	return h
}

{{- end -}}

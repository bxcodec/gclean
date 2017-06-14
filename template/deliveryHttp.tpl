package http

import (
 {{- range $key, $val := .Imports}}
    {{$key}}  "{{$val.Path}}"
 {{- end}}
 		"github.com/labstack/echo"
)




type HttpHandler struct {
	E *echo.Echo
}

func (h *HttpHandler) Init{{.ModelName}}Handler(u artUcase.{{.ModelName}}Usecase) *HttpHandler {
	{{.ModelName | lower }}Handler := &artHandler.{{.ModelName}}Handler{ {{substr 0 1 .ModelName}}Usecase: u}
	h.E.GET(`/{{.ModelName | lower}}`, {{.ModelName | lower }}Handler.Fetch)
	return h
}

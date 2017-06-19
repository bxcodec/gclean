// GENERATED BY glcean
// This file was generated by robots at
// {{ .TimeStamp }}

package models

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
type {{.ModelName | camelcase}} struct {
	{{ range $i,$att :=  .Attributes -}}
	 {{  $att.Name | camelcase    }}       {{$att.Type}}  `json:"{{$att.Name | snakecase}}"`
	{{ end -}}
}

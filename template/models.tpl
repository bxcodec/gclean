// go generate
// GENERATED BY THE COMMAND ABOVE
// This file was generated by robots at
// {{ .TimeStamp }}

package models


type {{.ModelName}} struct {
	{{- range $att :=  .Attributes }}
	 {{  $att.Name | camelcase    }} {{$att.Type}}  `json:"{{$att.Name | snakecase}}"`
	{{- end }}
}
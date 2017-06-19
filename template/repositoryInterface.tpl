package repository

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


{{- $pkgIm := index .Imports "models" -}}

{{- $modelName := .ModelName | camelcase -}}
{{- $Name := cat "*" $pkgIm.Alias "." $modelName  }}
{{- $model :=$Name| nospace -}}

type {{.ModelName | camelcase }}Repository interface{
  Fetch(cursor string , num int64) ([]{{$model}} ,error)
  Get(Id int)({{$model}},error)
  Update({{$model}}) (error)
  Delete(id int)(error)
}

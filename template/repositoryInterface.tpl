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

{{- $Name := cat "*" $pkgIm.Alias "." .ModelName  }}
{{ $modelName :=$Name| nospace }}
type {{.ModelName}}Repository interface{
  Fetch(cursor string , num int64) ([]{{$modelName}} ,error)
  Get(Id int)({{$modelName}},error)
  Update({{$modelName}}) (error)
  Delete(id int)(error)
}

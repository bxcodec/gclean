package {{.ModelName | lower}}

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



{{- $pkgIm := index .Imports "models" }}

{{ $modelName := .ModelName | camelcase }}
{{- $Name := cat "*" $pkgIm.Alias "." $modelName  }}
{{- $model :=$Name| nospace}}


type {{.ModelName | camelcase }}Usecase interface {
	Fetch(cursor string, num int64) ([]{{$model}}, string, error)
}

{{- $rS := cat .ModelName "Usecase" }}
{{- $repoStruct:= $rS|nospace | lower}}

type {{$repoStruct}} struct {
	{{.ModelName | lower}}Repos repository.{{.ModelName | camelcase }}Repository
}

func (a *{{$repoStruct}}) Fetch(cursor string, num int64) ([]{{$model}}, string, error) {
	if num == 0 {
		num = 10
	}

	listArticle, err := a.{{.ModelName | lower}}Repos.Fetch(cursor, num)
	if err != nil {
		return nil, "", err
	}
	nextCursor := ""


	return listArticle, nextCursor, nil
}

func New{{$repoStruct}}(a repository.{{.ModelName | camelcase }}Repository) {{ .ModelName | camelcase }}Usecase {
	return &{{$repoStruct}}{a}
}

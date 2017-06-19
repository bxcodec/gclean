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


{{- $pkgIm := index .Imports "models" -}}

{{- $modelName :=  .ModelName  | camelcase  }}

{{- $Name := cat "*" $pkgIm.Alias "." $modelName  }}
{{- $model :=$Name| nospace}}


{{- $rS := cat .Type  .ModelName "Repository" }}
{{- $repoStruct:= $rS|nospace | lower}}


{{ $sqlIm := index .Imports "sql" -}}
type {{$repoStruct  }} struct {
  Conn *{{$sqlIm.Alias}}.DB
}


func (h {{$repoStruct}}) fetch(query string, args ...interface{}) ([]{{$model}} ,error){

	rows, err := h.Conn.Query(query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result   []{{$model}}
	for rows.Next() {
		var  t {{$model}}
		err = rows.Scan(
    {{ range  $att :=  .Attributes -}}
  	      &t.{{  $att.Name | camelcase }},
  	{{ end -}}
		)

		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (h {{$repoStruct}} )  Fetch(cursor string , num int64) ([]{{$model}} ,error){
  query := `SELECT {{ range $att := .Attributes }}
                {{ $att.Name | snakecase}},
                {{- end }}
  						FROM {{.ModelName | lower}} WHERE ID > ? LIMIT ?`


	return h.fetch(query, cursor, num)
}

func (h {{$repoStruct}} )  Get(Id int)({{$model}},error){
  return nil,nil
}

func (h {{$repoStruct}} )  Update({{$model}}) (error){
  return nil
}

func (h {{$repoStruct}} )  Delete(id int)(error){
  return nil
}


{{$repoIm := index .Imports "repository" }}

func New{{.Type | title }}{{ .ModelName | camelcase }}Repository(c *{{$sqlIm.Alias}}.DB) {{$repoIm.Alias}}.{{.ModelName | camelcase }}Repository {
  repo:=&{{$repoStruct}}{
    Conn:c,
  }
  return repo
}

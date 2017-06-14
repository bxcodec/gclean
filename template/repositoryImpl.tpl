package {{.ModelName | lower}}

import (
 {{- range $key, $val := .Imports}}
    {{$key}}  "{{$val}}"
 {{- end}}
)
{{$Name := cat "*" .PackageShort "." .ModelName  }}
{{$modelName :=$Name| nospace}}


{{$rS := cat .Type  .ModelName "Repository" }}
{{$repoStruct:= $rS|nospace | lower}}


type {{$repoStruct}} struct {
  Conn *{{index .ImportsShort "sql" }}.DB
}


func (h {{$repoStruct}}) fetch(query string, args ...interface{}) ([]{{$modelName}} ,error){

	rows, err := h.Conn.Query(query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result   []{{$modelName}}
	for rows.Next() {
		var  t {{$modelName}}
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

func (h {{$repoStruct}} )  Fetch({{.FetchParams}}) ([]{{$modelName}} ,error){
  query := `SELECT {{ range $att := .Attributes }}
                {{ $att.Name | snakecase}},
                {{- end }}
  						FROM {{.ModelName | lower}} WHERE ID > ? LIMIT ?`


	return h.fetch(query, cursor, num)
}

func (h {{$repoStruct}} )  Get(Id int)({{$modelName}},error){
  return nil,nil
}

func (h {{$repoStruct}} )  Update({{$modelName}}) (error){
  return nil
}

func (h {{$repoStruct}} )  Delete(id int)(error){
  return nil
}



func New{{.Type | title }}{{ .ModelName }}Repository(c *{{index .ImportsShort "sql" }}.DB) {{index .ImportsShort "repository"}}.{{.ModelName}}Repository {
  repo:=&{{$repoStruct}}{
    Conn:c,
  }
  return repo
}

package repository

import (
 {{- range $key, $val := .Imports}}
    {{$key}}  "{{$val}}"
 {{- end}}
)

{{$pkgIm := index .ImportsShort "models" }}

{{$Name := cat "*" $pkgIm "." .ModelName  }}
{{$modelName :=$Name| nospace}}
type {{.ModelName}}Repository interface{
  Fetch({{.FetchParams}}) ([]{{$modelName}} ,error)
  Get(Id int)({{$modelName}},error)
  Update({{$modelName}}) (error)
  Delete(id int)(error)
}

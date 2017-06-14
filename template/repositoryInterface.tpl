package repository

import (
 {{- range $key, $val := .Imports}}
    {{$key}}  "{{$val.Path}}"
 {{- end}}
)

{{$pkgIm := index .Imports "models" }}

{{$Name := cat "*" $pkgIm.Alias "." .ModelName  }}
{{$modelName :=$Name| nospace}}
type {{.ModelName}}Repository interface{
  Fetch(cursor string , num int64) ([]{{$modelName}} ,error)
  Get(Id int)({{$modelName}},error)
  Update({{$modelName}}) (error)
  Delete(id int)(error)
}

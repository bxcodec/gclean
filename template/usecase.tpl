package {{.ModelName | lower}}

import (
 {{- range $key, $val := .Imports}}
    {{$key}}  "{{$val}}"
 {{- end}}
)

{{$Name := cat "*" .PackageShort "." .ModelName  }}
{{$modelName :=$Name| nospace}}


type {{.ModelName}}Usecase interface {
	Fetch(cursor string, num int64) ([]{{$modelName}}, string, error)
}

{{$rS := cat .ModelName "Usecase" }}
{{$repoStruct:= $rS|nospace | lower}}

type {{$repoStruct}} struct {
	articleRepos repository.{{.ModelName}}Repository
}

func (a *{{$repoStruct}}) Fetch(cursor string, num int64) ([]{{$modelName}}, string, error) {
	if num == 0 {
		num = 10
	}

	listArticle, err := a.articleRepos.Fetch(cursor, num)
	if err != nil {
		return nil, "", err
	}
	nextCursor := ""


	return listArticle, nextCursor, nil
}

func New{{$repoStruct}}(a repository.{{.ModelName}}Repository) ArticleUsecase {
	return &{{$repoStruct}}{a}
}

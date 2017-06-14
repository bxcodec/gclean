package subcommands

import (
	"fmt"
	"os"
	"strconv"
	"text/template"

	"github.com/Masterminds/sprig"
)

type Usecase struct {
	PackageShort string
	ModelName    string
	FetchParams  string
	Imports      map[string]string
	ImportsShort map[string]string
	Attributes   []*Attribute
}

func (s *Subs) generateUsecase(modelName string) {
	temp, err := template.New("").Funcs(sprig.TxtFuncMap()).ParseFiles("template/usecase.tpl")

	if err != nil {
		fmt.Println("GALGAL", err)
		os.Exit(0)
	}

	pathP := "usecase/" + modelName + "/"
	if _, er := os.Stat(pathP); os.IsNotExist(er) {
		os.MkdirAll(pathP, os.ModePerm)
	}

	k := 1

	mapImport := make(map[string]string)
	mapImport["models"] = "github.com/bxcodec/gclean/models"
	mapImport["time"] = "time"
	mapImport["sql"] = "database/sql"
	mapImport["repository"] = "github.com/bxcodec/gclean/repository"

	mapImportShort := make(map[string]string)
	mapImportShort["models"] = "github.com/bxcodec/gclean/models"
	mapImportShort["sql"] = "sql"
	mapImportShort["repository"] = "repository"

	id := &Attribute{
		Name: "ID",
		Type: "int64",
	}
	title := &Attribute{
		Name: "Title",
		Type: "string",
	}
	content := &Attribute{
		Name: "Content",
		Type: "string",
	}

	dataSend := &Usecase{
		PackageShort: "models",
		ModelName:    "Article",
		FetchParams:  "cursor string , num int64",
		Imports:      mapImport,
		ImportsShort: mapImportShort,
		Attributes:   []*Attribute{id, title, content},
	}
	f, err := os.Create(pathP + "sample" + strconv.Itoa(k) + ".go")
	if err != nil {
		fmt.Println("Erorr")
	}

	defer f.Close()
	err = temp.ExecuteTemplate(f, "usecase.tpl", dataSend)

	if err != nil {
		fmt.Println("ERROR ", err)
		os.Exit(0)
	}
}

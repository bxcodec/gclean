package subcommands

import (
	"fmt"
	"os"
	"strconv"
	"text/template"

	"github.com/Masterminds/sprig"
)

type Usecase struct {
	ModelName  string
	Imports    map[string]*Import
	Attributes []*Attribute
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

	mapImport := make(map[string]*Import)

	m := &Import{Alias: "models", Path: "github.com/bxcodec/gclean/models"}
	t := &Import{Alias: "time", Path: "time"}
	ss := &Import{Alias: "sql", Path: "database/sql"}
	r := &Import{Alias: "repository", Path: "github.com/bxcodec/gclean/repository"}
	mapImport["models"] = m
	mapImport["time"] = t
	mapImport["sql"] = ss
	mapImport["repository"] = r

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
		ModelName:  "Article",
		Imports:    mapImport,
		Attributes: []*Attribute{id, title, content},
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
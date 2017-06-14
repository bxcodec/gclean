package subcommands

import (
	"fmt"
	"os"
	"strconv"
	"text/template"

	"github.com/Masterminds/sprig"
)

type RepoGenerator struct {
	ModelName string
	Imports   map[string]*Import
}

type Import struct {
	Alias string
	Path  string
}

func (s *Subs) generateRepository() {
	temp, err := template.New("").Funcs(sprig.TxtFuncMap()).ParseFiles("template/repositoryInterface.tpl")

	if err != nil {
		fmt.Println("GALGAL", err)
		os.Exit(0)
	}

	pathP := "repository/"
	if _, er := os.Stat(pathP); os.IsNotExist(er) {
		os.MkdirAll(pathP, os.ModePerm)
	}

	k := 1
	t := &Import{Alias: "time", Path: "time"}
	m := &Import{Alias: "models", Path: "github.com/bxcodec/gclean/models"}
	mapImport := make(map[string]*Import)
	mapImport["models"] = m
	mapImport["time"] = t

	dataSend := &RepoGenerator{
		ModelName: "Article",
		Imports:   mapImport,
	}
	f, err := os.Create(pathP + "sample" + strconv.Itoa(k) + ".go")
	if err != nil {
		fmt.Println("Erorr")
	}

	defer f.Close()
	err = temp.ExecuteTemplate(f, "repositoryInterface.tpl", dataSend)

	if err != nil {
		fmt.Println("ERROR ", err)
		os.Exit(0)
	}
}

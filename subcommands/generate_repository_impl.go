package subcommands

import (
	"fmt"
	"os"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/bxcodec/gclean/subcommands/models"
)

func (s *Subs) generateRepositoryImpl(dataSend *models.DataGenerator) {
	temp, err := template.New("").Funcs(sprig.TxtFuncMap()).ParseFiles("template/repositoryImpl.tpl")

	if err != nil {
		fmt.Println("GALGAL", err)
		os.Exit(0)
	}

	pathP := "repository/" + dataSend.Type + "/" + dataSend.ModelName + "/"
	if _, er := os.Stat(pathP); os.IsNotExist(er) {
		os.MkdirAll(pathP, os.ModePerm)
	}
	//
	// k := 1

	// mapImport := make(map[string]models.Import)
	//
	// m := models.Import{Alias: "models", Path: "github.com/bxcodec/gclean/models"}
	// t := models.Import{Alias: "time", Path: "time"}
	// ss := models.Import{Alias: "sql", Path: "database/sql"}
	// r := models.Import{Alias: "repository", Path: "github.com/bxcodec/gclean/repository"}
	// mapImport["models"] = m
	// mapImport["time"] = t
	// mapImport["sql"] = ss
	// mapImport["repository"] = r

	f, err := os.Create(pathP + dataSend.ModelName + "_" + dataSend.Type + ".go")
	if err != nil {
		fmt.Println("Erorr")
	}

	defer f.Close()
	err = temp.ExecuteTemplate(f, "repositoryImpl.tpl", dataSend)

	if err != nil {
		fmt.Println("ERROR ", err)
		os.Exit(0)
	}
}

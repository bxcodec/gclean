package subcommands

import (
	"fmt"
	"os"
	"text/template"

	"github.com/Masterminds/sprig"
)

func (s *Subs) generateDelivery() {
	temp, err := template.New("").Funcs(sprig.TxtFuncMap()).ParseFiles("template/deliveryHttp.tpl")

	if err != nil {
		fmt.Println("GALGAL", err)
		os.Exit(0)
	}

	pathP := "delivery/http/"
	if _, er := os.Stat(pathP); os.IsNotExist(er) {
		os.MkdirAll(pathP, os.ModePerm)
	}

	mapImport := make(map[string]*Import)

	m := &Import{Alias: "models", Path: "github.com/bxcodec/gclean/models"}
	t := &Import{Alias: "time", Path: "time"}
	ss := &Import{Alias: "sql", Path: "database/sql"}
	r := &Import{Alias: "repository", Path: "github.com/bxcodec/gclean/repository"}
	a := &Import{Alias: "articleUcase", Path: "github.com/bxcodec/gclean/delivery/http/article"}
	mapImport["models"] = m
	mapImport["time"] = t
	mapImport["sql"] = ss
	mapImport["repository"] = r
	mapImport["article"] = a

	dataSend := &DataGenerator{
		ModelName: "Article",
		Imports:   mapImport,
	}
	f, err := os.Create(pathP + "http_deliver.go")
	if err != nil {
		fmt.Println("Erorr")
	}

	defer f.Close()
	err = temp.ExecuteTemplate(f, "deliveryHttp.tpl", dataSend)

	if err != nil {
		fmt.Println("ERROR ", err)
		os.Exit(0)
	}
}

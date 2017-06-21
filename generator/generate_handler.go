package generator

import (
	"fmt"
	"os"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/bxcodec/gclean/generator/models"
)

func (s *Subs) generateHandler(data *models.DataGenerator) {
	temp, err := template.New("").Funcs(sprig.TxtFuncMap()).ParseFiles("template/handlerhttp.tpl")

	if err != nil {
		fmt.Println("GALGAL", err)
		os.Exit(0)
	}

	pathP := "delivery/http/" + data.ModelName + "/"
	if _, er := os.Stat(pathP); os.IsNotExist(er) {
		os.MkdirAll(pathP, os.ModePerm)
	}

	f, err := os.Create(pathP + data.ModelName + "_handler.go")
	if err != nil {
		fmt.Println("Erorr")
	}

	defer f.Close()
	err = temp.ExecuteTemplate(f, "handlerhttp.tpl", data)

	if err != nil {
		fmt.Println("ERROR ", err)
		os.Exit(0)
	}
}

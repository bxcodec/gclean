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
